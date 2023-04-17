package consensus

import (
	"context"
	"crypto"
	"crypto/ed25519"
	"errors"
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/creachadair/taskgroup"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Used in the handshake for ensuring compatibility. All consensus engines in a network
// should be on the same version
const Version = 1

// Engine is the core struct that performs byzantine fault tolerant state
// machine replication using the Tendermint protocol.
//
// In order to function it depends on a networking implementation that completes
// the Sender and Receiver interfaces, a state machine for building, verifying,
// executing and persisting data and an optional signer which is necessary as a writer
// in the network to sign votes and proposals using a secured private key
//
// Engine can either be bundled in the same process (in the case of a state machine,
// networking layer and signer written in golang) or can be in a separate process
// (in which gRPC is used to communicate).
//
// The engine runs only in memory and is thus not responsible for persistence and crash
// recovery. Each time the application starts it uses the handshake with the application
// to set or restore the height and other parameters the engine needs to continue
type Engine struct {
	// The application the consensus engine is communicating with to provide SMR
	app Application

	// gossip represents a simple networking abstraction for broadcasting messages
	// that should eventually propagate to all non-faulty nodes in the network as
	// well as eventually receiving all messages generated from other nodes.
	gossip Gossip

	// signer is only used if the node is a validator or writer in the network
	// as opposed to a reader or full node, in which case this can be nil.
	// The signer is responsible for signing votes and proposals.
	signer Signer
	// we save our public key so we can recognise when we need to propose.
	ourPubkey []byte

	// status tracks if the engine is running or not.
	status uint32

	// Executor tracks the main consensus state: height, round and the
	// logic for deciding when to vote, what to vote and handling the
	// finalization of a proposal
	executor *executor

	// Verifier verifies proposals and votes
	verifier *Verifier

	// Tally keeps track of all votes and proposals.
	tally *tally

	// The following are used for managing the lifecycle of the engine
	closer chan struct{}
	done   chan struct{}

	logger zerolog.Logger
}

// Option is a set of configurable parameters. If left empty, defaults
// will be used
type Option func(e *Engine)

// WithVerifyFunc sets the verify function for verifying signatures
func WithCustomVerifyFunc(f VerifyFunc) Option {
	return func(e *Engine) {
		e.verifier.verifyFunc = f
	}
}

// WithHashFunc sets the hash function for hashing proposal data
func WithCustomHashFunc(f crypto.Hash) Option {
	return func(e *Engine) {
		e.verifier.hasher = f
	}
}

// New creates a new consensus engine
func New(app Application, gossip Gossip, signer Signer, opts ...Option) *Engine {
	e := &Engine{
		app:    app,
		gossip: gossip,
		signer: signer,
		verifier: &Verifier{
			verifyFunc: DefaultVerifyFunc(),
			hasher:     crypto.SHA256,
		},
		logger: zerolog.New(os.Stdout),
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

// Operational phases
const (
	Off = iota
	StartingUp
	On
	ShuttingDown
)

// retry handling of grpc connections
const (
	maxRetryAttempts         = 7
	exponentialBackoffFactor = 40 // 40ms
)

func (e *Engine) Run(ctx context.Context) error {
	if err := e.startUp(ctx); err != nil {
		return err
	}
	defer e.shutDown()
	return e.run(ctx)
}

func (e *Engine) IsRunning() bool {
	return atomic.LoadUint32(&e.status) == On
}

func (e *Engine) Start(ctx context.Context) error {
	if err := e.startUp(ctx); err != nil {
		return err
	}
	go e.run(context.Background())
	return nil
}

func (e *Engine) Stop() error {
	return e.shutDown()
}

func (e *Engine) startUp(ctx context.Context) error {
	if !atomic.CompareAndSwapUint32(&e.status, Off, StartingUp) {
		return errors.New("engine already running")
	}

	var (
		resp *HandshakeResponse
		err  error
	)

	// Attempt to establish the handshake with the application, providing version numbering
	// and obtaining state persisted info such as height and round. Unlike other gRPC calls
	// this happens indefinitely until the handshake is successful, an unrecoverable error
	// is returned or the context is cancelled
	attempts := 0
RETRY_LOOP:
	for {
		resp, err = e.app.Handshake(ctx, &HandshakeRequest{
			Version: Version,
		})
		switch status.Code(err) {
		case codes.OK:
			break RETRY_LOOP
		case codes.Unavailable, codes.DeadlineExceeded:
			e.logger.Info().
				Int("attempts", attempts).
				Str("response", err.Error()).
				Msg("handshaking with application; retrying with exponential backoff")
			time.Sleep(exponentialBackoff(attempts))
		default:
			atomic.StoreUint32(&e.status, Off)
			return err
		}
		attempts++
	}

	if e.signer != nil {
		resp, err := e.signer.Handshake(ctx, &SignerHandshakeRequest{
			Height:    resp.Height,
			Round:     resp.Round,
			Version:   Version,
			Namespace: resp.Namespace,
		})
		if err != nil {
			atomic.StoreUint32(&e.status, Off)
			return err
		}
		e.ourPubkey = resp.PubKey
	}

	atomic.CompareAndSwapUint32(&e.status, StartingUp, On)

	return nil
}

func (e *Engine) run(ctx context.Context) error {
	if atomic.LoadUint32(&e.status) != On {
		return errors.New("engine is not running")
	}
	defer close(e.done)
	ctx, cancel := context.WithCancel(ctx)
	tg := taskgroup.New(func(err error) error {
		// close all other go routines upon the first failure
		cancel()
		return err
	})

	tg.Go(func() error {
		return e.receiveProposal(ctx)
	})
	tg.Go(func() error {
		return e.receiveVotes(ctx)
	})
	tg.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-e.closer:
			return ErrApplicationShutdown
		}
	})

	return tg.Wait()
}

func (e *Engine) shutDown() error {
	if !atomic.CompareAndSwapUint32(&e.status, On, ShuttingDown) {
		return errors.New("engine is not running")
	}

	close(e.closer)

	<-e.done

	atomic.StoreUint32(&e.status, Off)

	return nil
}

// Default to ed25519
func DefaultVerifyFunc() VerifyFunc {
	return func(publicKey, message, signature []byte) bool {
		return ed25519.Verify(publicKey, message, signature)
	}
}

var ErrApplicationShutdown = errors.New("application requested termination")

// unrecoverable errors indicate that the consensus engine
// is in a state that is not recoverable. It thus logs the
// error and shutsdown.
type errUnrecoverable struct {
	err error
}

func unrecoverable(err error) error {
	return errUnrecoverable{
		err: err,
	}
}

func isUnrecoverable(err error) bool {
	_, ok := err.(errUnrecoverable)
	return ok
}

func (e errUnrecoverable) Error() string {
	return fmt.Sprintf("unrecoverable error: %w", e.err)
}
