package kbaseserver

import (
	"time"

	"kbase.com/raft"
	"kbase.com/server/config"
)

func bootstrap(cfg config.ServerConfig) (b *bootstrapedServer, err error) {
	raft := bootstrapRaft()
	return &bootstrapedServer{
		raft: raft,
	}, nil
}

type bootstrapedServer struct {
	raft *bootstrapedRaft
}

type bootstrapedRaft struct {
	heartbeat time.Duration
	config    *raft.Config
}

func bootstrapRaft() *bootstrapedRaft {
	return &bootstrapedRaft{}
}

func (b bootstrapedRaft) newRaftNode() *raftNode {
	var n raft.Node
	n = raft.StartNode(b.config)

	return newRaftNode(
		raftNodeConfig{
			heartbeat: b.heartbeat,
			Node:      n,
		},
	)
}
