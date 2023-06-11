package consensus

import "context"

func (e *Engine) OnProposal(ctx context.Context, prp *Proposal) error {
	h, err := e.heights.Get(prp.Height)
	if err != nil {
		return err
	}

	err = h.ReceiveProposal(ctx, prp)
	if err != nil {
		return err
	}

	return nil
}

func (e *Engine) OnVote(ctx context.Context, vote *Vote) error {
	h, err := e.heights.Get(vote.Height)
	if err != nil {
		return err
	}

	err = h.ReceiveVote(ctx, vote)
	if err != nil {
		return err
	}

	return nil
}
