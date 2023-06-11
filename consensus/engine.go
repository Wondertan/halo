package consensus

import (
	"context"
	"crypto"
	"errors"
	"os"
	"sync/atomic"

	"github.com/cmwaters/halo/pkg/app"
	"github.com/cmwaters/halo/pkg/group"
	"github.com/cmwaters/halo/pkg/sign"
	"github.com/rs/zerolog"
)

// Used in the handshake for ensuring compatibility. All consensus engines in a network
// should be on the same version
const Version = 1

var _ Service = &Engine{}
var _ Consensus = &Engine{}

// Engine is the core struct that performs byzantine fault tolerant state
// machine replication using the Lock and Commit protocol.
//
// In order to function it depends on a networking implementation that completes
// the Gossip interface, a state machine for building, verifying, executing and
// persisting data and an optional signer which is necessary as a writer
// in the network to sign votes and proposals using a secured private key
//
// The engine runs only in memory and is thus not responsible for persistence and crash
// recovery. Each time start is called, a height and state machine is provided
type Engine struct {
	// namespace represents the unique id of the application. It is paired with
	// height and round to signal uniqueness of proposals and votes
	namespace []byte

	heights *heights

	// gossip represents a simple networking abstraction for broadcasting messages
	// that should eventually propagate to all non-faulty nodes in the network as
	// well as eventually receiving all messages generated from other nodes.
	gossip Gossip

	// signer is only used if the node is a validator or writer in the network
	// as opposed to a reader or full node, in which case this can be nil.
	// The signer is responsible for signing votes and proposals.
	signer sign.Signer

	// parameters entails the set of consensus specfic parameters that are used
	// to reach consensus
	parameters Parameters

	// hasher defines how proposal data is hashed for signing
	hasher crypto.Hash

	// status tracks if the engine is running or not.
	status atomic.Bool

	// The following are used for managing the lifecycle of the engine
	cancel context.CancelFunc
	done   chan struct{}

	logger zerolog.Logger
}

// New creates a new consensus engine
// TODO(@Wondertan): Combine Options and parameters
func New(namespace []byte, net Network, signer sign.Signer, parameters Parameters, opts ...Option) (*Engine, error) {
	gossip, err := net.Gossip(namespace)
	if err != nil {
		return nil, err
	}

	e := &Engine{
		namespace:  namespace,
		gossip:     gossip,
		signer:     signer,
		parameters: parameters,
		hasher:     DefaultHashFunc,
		logger:     zerolog.New(os.Stdout),
	}

	for _, opt := range opts {
		opt(e)
	}

	// TODO(@Wondertan): Cleanup as method
	e.heights = newHeights(func(heightN uint64) (*height, error) {
		// TODO(@Wondertan): more sanity checks whether the height is adequate

		store := NewStore()
		executor := newExecutor(
			e.parameters.ProposalTimeout,
			e.parameters.LockDelay,
			false,
		)

		return &height{
			verifier:  NewVerifier(e.namespace, heightN, e.hasher),
			executor:  executor,
			store:     store,
			logger:    e.logger,
			activated: make(chan struct{}),
			proposed:  make(chan struct{}),
		}, nil
	})

	// register engine for message notifications
	gossip.Notify(e)
	return e, nil
}

// Operational phases
const (
	Off = false
	On  = true
)

// Start implements the Service interface and starts the consensus engine
func (e *Engine) Start(ctx context.Context, height uint64, app app.StateMachine) error {
	if !e.status.CompareAndSwap(Off, On) {
		return errors.New("engine already running")
	}
	defer e.status.CompareAndSwap(On, Off)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	e.cancel = cancel
	e.done = make(chan struct{})
	defer close(e.done)
	for {
		group, err := app.Initialize(ctx, height)
		if err != nil {
			return err
		}
		data, commitment, err := e.Commit(ctx, height, group, app.Propose, app.VerifyProposal)
		if err != nil {
			return err
		}
		if err := app.Finalize(ctx, height, data, commitment); err != nil {
			return err
		}
		height++
	}
}

// Commit implements the Consensus interface. It executes a single instance of consensus between
// a group of participants. It takes in two hooks, one for proposing data and another for verifying
// the data. It returns the data that was agreed upon by the group under the lock and commit protocol
func (e *Engine) Commit(
	ctx context.Context,
	height uint64,
	group group.Group,
	appProposeFn app.Propose,
	verifyProposalFn app.VerifyProposal,
) ([]byte, group.Commitment, error) {
	// for each height, the process splits up the task through three components:
	// - Verifier: responsible for verifying the validity of proposals and votes
	// - Store: responsible for storing proposals and votes
	// - Executor: responsible for tallying votes and proposals and for voting and
	//          proposing according to the rules of the consensus protocol

	h, err := e.heights.Get(height)
	if err != nil {
		return nil, nil, err
	}
	voteFn, selfWeight := e.voteFn(height, group, h.store)
	proposeFn := e.proposeFn(height, group, h.store, appProposeFn)

	// TODO(@Wondertan): I don't like this "updating" but for now that's the quickiest
	//  workaround I found
	h.executor.Update(voteFn, proposeFn, selfWeight, group.TotalWeight())
	h.verifier.Update(group, verifyProposalFn)
	h.Activate()

	proposalRound, err := h.executor.Run(ctx)
	if err != nil {
		return nil, nil, err
	}

	_, proposal := h.store.GetProposal(proposalRound)
	commitment, err := h.store.CreateCommitment(proposalRound)
	if err != nil {
		return nil, nil, err
	}

	return proposal.Data, commitment, nil
}

func (e *Engine) Stop() error {
	if !e.status.CompareAndSwap(On, Off) {
		return errors.New("engine is not running")
	}
	e.cancel()
	<-e.Wait()
	return nil
}

func (e *Engine) StopAtHeight(height uint64) error {
	panic("not implemented")
}

func (e *Engine) Wait() <-chan struct{} {
	return e.done
}

var ErrApplicationShutdown = errors.New("application requested termination")
