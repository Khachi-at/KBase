package kbaseserver

import (
	"context"

	"kbase.com/raft/raftpb"
	"kbase.com/server/config"
)

type KBaseServer struct {
	r raftNode

	// stop signals the run goroutine should shutdown.
	stop chan struct{}
}

func NewServer(cfg config.ServerConfig) (srv *KBaseServer, err error) {
	b, err := bootstrap(cfg)
	if err != nil {
		return nil, err
	}
	srv = &KBaseServer{
		r: *b.raft.newRaftNode(),
	}
	return
}

func (s *KBaseServer) Start() {
	// TODO:

	s.run()
}

func (s *KBaseServer) run() {
	raftReadyHandler := &raftReadyHandler{}
	s.r.start(raftReadyHandler)

	// TODO: exec when stop server
	// defer func() {}

	for {
		select {
		case <-s.stop:
			return
		}
	}
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
