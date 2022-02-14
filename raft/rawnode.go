package raft

type RawNode struct {
	raft *raft
}

func NewRawNode(cfg *Config) (*RawNode, error) {
	return &RawNode{
		raft: newRaft(cfg),
	}, nil
}

// Tick advances the internal logical clock by a single tick.
func (r *RawNode) Tick() {
	r.raft.tick()
}

func (r *RawNode) HasReady() bool {
	if len(r.raft.msgs) > 0 {
		return true
	}
	return false
}

func (r *RawNode) readyWithoutAccept() Ready {
	return newReady(r.raft)
}

func (r *RawNode) acceptReady(rd Ready) {
	r.raft.msgs = nil
}

func newReady(r *raft) Ready {
	rd := Ready{
		Messages: r.msgs,
	}
	return rd
}
