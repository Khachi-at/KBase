package kbaseserver

import (
	"context"

	"kbase.com/raft/raftpb"
)

type KBaseServer struct {
	r raftNode
}

func (s *KBaseServer) Process(ctx context.Context, m raftpb.Message) error {
	return s.r.Step(ctx, m)
}

type raftReadyHandler struct {
	getLead              func() (lead uint64)
	updateLead           func(lead uint64)
	updateLeadership     func(newLeader bool)
	updateCommittedIndex func(uint64)
}
