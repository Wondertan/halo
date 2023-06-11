package consensus

import (
	"context"
	"sync"

	"github.com/rs/zerolog"
)

type newHeightFn func(uint64) (*height, error)

// height
// TODO(@Wondertan): Better name?
type height struct {
	verifier *Verifier
	executor *Executor
	store    *Store
	logger   zerolog.Logger

	activated chan struct{}
	proposed  chan struct{}
}

// heights
// TODO(@Wonderta): Possible GC strategy is to clean up a height
//
//	right after we committed, but we may wanna catch
//	late votes, so instead we may clean up the height once it N heights old.
type heights struct {
	newHeight newHeightFn
	heightsLk sync.RWMutex
	heights   map[uint64]*height
}

func newHeights(new newHeightFn) *heights {
	return &heights{
		newHeight: new,
		heights:   make(map[uint64]*height),
	}
}

// Get
// TODO(@Wondertan): Find a better name
//
//	Setup??
//	GetOrCreate??
func (rs *heights) Get(height uint64) (*height, error) {
	rs.heightsLk.RLock()
	h, ok := rs.heights[height]
	if ok {
		rs.heightsLk.Unlock()
		return h, nil
	}
	rs.heightsLk.Unlock()

	rs.heightsLk.Lock()
	defer rs.heightsLk.Unlock()
	if h, ok = rs.heights[height]; ok {
		return h, nil
	}

	h, err := rs.newHeight(height)
	if err != nil {
		return nil, err
	}

	rs.heights[height] = h
	return h, nil
}

func (r *height) Activate() {
	close(r.activated)
}

func (r *height) ReceiveProposal(ctx context.Context, proposal *Proposal) error {
	if err := proposal.ValidateForm(); err != nil {
		return err
	}

	select {
	case <-r.activated:
	case <-ctx.Done():
		return ctx.Err()
	}

	// Check that we don't already have the proposal
	if r.store.HasProposal(proposal.Round) {
		r.logger.Info().
			Uint64("height", proposal.Height).
			Uint32("height", proposal.Round).
			Msg("proposal already in executor")
		return nil
	}

	// Verify the proposal. Including that it came from the correct proposer and
	// that the signature for the data is valid.
	if err := r.verifier.VerifyProposal(ctx, proposal); err != nil {
		return err
	}

	// Add the proposal to the executor
	r.store.AddProposal(proposal)

	// pass the proposal on to the state machine to process
	r.executor.ProcessProposal(proposal.Round)

	close(r.proposed)
	return nil
}

func (r *height) ReceiveVote(ctx context.Context, vote *Vote) error {
	if err := vote.ValidateForm(); err != nil {
		return err
	}

	select {
	case <-r.activated:
	case <-ctx.Done():
		return ctx.Err()
	}

	// Check that we don't already have the vote
	if r.store.HasVote(vote) {
		r.logger.Info().
			Uint64("height", vote.Height).
			Uint32("height", vote.Round).
			Msg("vote is already in executor")
		return nil
	}

	// Get the proposal to verify the vote against
	proposalID, _ := r.store.GetProposal(vote.ProposalRound)
	if proposalID == nil {
		if vote.Height == r.verifier.Height() {
			r.store.AddPendingVote(vote)
		}
		// TODO: We may want to consider handling votes at the height directly above
		// the nodes current state.
		return nil
	}

	// verify the vote's signature and other components
	if err := r.verifier.VerifyVote(vote, proposalID); err != nil {
		return err
	}

	added := r.store.AddVote(vote)
	if !added {
		// there will have been no state transition so we can exit early
		return nil
	}

	votersWeight := r.verifier.GetMember(vote.MemberIndex).Weight()
	r.executor.ProcessVote(vote.Round, vote.ProposalRound, votersWeight, VoteType(vote.Commit))
	return nil
}
