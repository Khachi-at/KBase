package raft

import (
	"context"

	pb "kbase.com/raft/raftpb"
)

type Node interface {
	// Tick increments the internal logic clock for the Node by a single tick.
	Tick()

	// Ready returns a channel that returns the current point-in-time state.
	Ready() <-chan Ready

	Step(ctx context.Context, msg pb.Message) error
}

type node struct {
	readyc chan Ready
	tickc  chan struct{}

	rn *RawNode
}

func newNode(rn *RawNode) node {
	return node{
		readyc: make(chan Ready),
		tickc:  make(chan struct{}),

		rn: rn,
	}
}

func (n *node) run() {
	var readyc chan Ready
	var rd Ready

	//r := n.rn.raft

	for {
		if n.rn.HasReady() {
			readyc = n.readyc
		}

		select {
		case <-n.tickc:
			n.rn.Tick()
		case readyc <- rd:
			n.rn.acceptReady(rd)
		}
	}

}

func (n *node) Tick() {
	select {
	case n.tickc <- struct{}{}:
	}
}

func (n *node) Ready() <-chan Ready { return n.readyc }

type Ready struct {
	*SoftState
	Messages []pb.Message
}

type SoftState struct {
	Lead      uint64
	RaftState StateType
}
