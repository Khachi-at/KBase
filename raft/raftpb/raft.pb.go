package raftpb

// Message is a instance of one message.
type Message struct {
	Type MessageType `json:"type"`
	To   uint64
	From uint64
	Term uint64
}

type MessageType int32

const (
	MsgHup  MessageType = iota + 1 // MsgHup is used for election.
	MsgBeat                        // MsgBeat is an internal type that signals the leader to send a heartbeat of MsgHeartBeat type.
	MsgProp                        // MsgProp proposes to append data to its log entries.
	MsgApp
	MsgVote
	MsgVoteResp
	MsgSnap
	MsgHeartBeat
	MsgHeartBeatResp
	MsgPreVote
	MsgPreVoteResp
)
