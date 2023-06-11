package consensus

import (
	"context"
	"io"

	"github.com/cmwaters/halo/pkg/app"
	"github.com/cmwaters/halo/pkg/group"
)

type (
	// Service embodies a process that is perpetually running the underlying consensus
	// protocol. Each iteration is known as a height. The service starts by being specified
	// a height and a state machine that it will perform replication on and continues until
	// an error is encountered, or it is stopped.
	Service interface {
		Start(context.Context, uint64, app.StateMachine) error
		Stop() error
		StopAtHeight(uint64) error
		Wait() <-chan struct{}
	}

	// Consensus is an interface that allows the caller to `Commit` a value, either
	// proposed from it's own process or by a participant in the "Group". Underlying
	// a Consensus instance is a networking layer, responsible for broadcasting proposals
	// and votes to one another.
	Consensus interface {
		Commit(
			context.Context,
			uint64,
			group.Group,
			app.Propose,
			app.VerifyProposal,
		) ([]byte, group.Commitment, error)
	}
)

// TODO(@Wondertan): Move to independent network pkg once dependency cycle on types is resolved

// Namespace
// TODO(@Wondertan): Actually use it
type Namespace string

type Network interface {
	Gossip(namespace []byte) (Gossip, error)
}

// Gossip is an interface which allows the consensus engine to both broadcast
// and receive messages to and from other nodes in the network. It must eventually
// propagate messages to all non-faulty nodes within the network. The algorithm
// for how this is done i.e. simply flooding the network or using some form of
// content addressing protocol is left to the implementer.
type Gossip interface {
	io.Closer
	Broadcaster
	Notifier
}

type Broadcaster interface {
	BroadcastProposal(context.Context, *Proposal) error
	BroadcastVote(context.Context, *Vote) error
}

type Notifier interface {
	// Notify registers Notifiee wishing to receive notifications about new messages.
	// Any non-nil error returned from On... handlers rejects the message as invalid.
	Notify(Notifiee)
}

// Notifiee
// TODO(@Wondertan): The more I look into this, the more I want to abstract
// Proposal and Vote into a Message.
type Notifiee interface {
	OnProposal(context.Context, *Proposal) error
	OnVote(context.Context, *Vote) error
}
