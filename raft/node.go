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

func StartNode(cfg *Config) Node {
	rn, err := NewRawNode(cfg)
	if err != nil {
		panic(err)
	}
	n := newNode(rn)

	go n.run()
	return &n
}

type node struct {
	recvc  chan pb.Message
	readyc chan Ready
	tickc  chan struct{}

	rn *RawNode
}

func newNode(rn *RawNode) node {
	return node{
		recvc:  make(chan pb.Message),
		readyc: make(chan Ready),
		tickc:  make(chan struct{}),

		rn: rn,
	}
}

func (n *node) run() {
	var readyc chan Ready
	var rd Ready

	r := n.rn.raft

	for {
		if n.rn.HasReady() {
			// Populate a ready.
			rd = n.rn.readyWithoutAccept()
			readyc = n.readyc
		}

		select {
		case m := <-n.recvc:
			// TODO: filter out response message from unknown.
			r.Step(m)
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

func (n *node) Step(ctx context.Context, msg pb.Message) error {
	return n.step(ctx, msg)
}

func (n *node) step(ctx context.Context, msg pb.Message) error {
	if msg.Type != pb.MsgProp {
		select {
		case n.recvc <- msg:
			return nil
		}
	}
	return nil
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
