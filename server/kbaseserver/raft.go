package kbaseserver

import (
	"sync"
	"time"

	"go.etcd.io/etcd/pkg/contention"
	"go.uber.org/zap"
	"kbase.com/diplomat"
	"kbase.com/raft"
)

type raftNode struct {
	lg *zap.Logger

	tickMu *sync.Mutex
	raftNodeConfig

	// utility
	ticker *time.Ticker

	td *contention.TimeoutDetector

	stopped chan struct{}
}

type raftNodeConfig struct {
	lg *zap.Logger

	raft.Node

	heartbeat time.Duration
	// diplomat specifies the diplomat to send and receive messages to members.
	diplomat diplomat.Diplomat
}

func (r *raftNode) tick() {
	r.tickMu.Lock()
	r.Tick()
	r.tickMu.Unlock()
}

func newRaftNode(cfg raftNodeConfig) *raftNode {
	// TODO: set log
	r := &raftNode{
		lg:             cfg.lg,
		tickMu:         new(sync.Mutex),
		raftNodeConfig: cfg,
	}
	if r.heartbeat == 0 {
		r.ticker = &time.Ticker{}
	} else {
		r.ticker = time.NewTicker(r.heartbeat)
	}
	return r
}

func (r *raftNode) start(rh *raftReadyHandler) {
	go func() {
		defer r.stop()

		for {
			select {
			case <-r.ticker.C:
				r.tick()
			case rd := <-r.Ready():
				r.diplomat.Send(rd.Messages)

			case <-r.stopped:
				return
			}
		}
	}()
}

func (r *raftNode) stop() {}
