package diplomat

import (
	"context"

	"kbase.com/raft/raftpb"
)

// When Message reach, call function Process to handle.
type Raft interface {
	Process(ctx context.Context, m raftpb.Message)
}

// Diplomat provide functions which used by server.
// The interface conuld be impl by all protos.
type Diplomat interface {
	// Send sends out the given messages to the remote peers.
	Send(m []raftpb.Message)
}
