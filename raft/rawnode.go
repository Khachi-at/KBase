package raft

type RawNode struct {
	raft *raft
}

// Tick advances the internal logical clock by a single tick.
func (r *RawNode) Tick() {
	r.raft.tick()
}

func (r *RawNode) HasReady() bool {
	return false
}

func (r *RawNode) acceptReady(rd Ready) {}
