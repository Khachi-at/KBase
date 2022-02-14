package raftpb

// Message is a instance of one message.
type Message struct {
	Type MessageType `json:"type"`
}

type MessageType int32

const (
	MsgHup  MessageType = iota + 1 // MsgHup is used for election.
	MsgBeat                        // MsgBeat is an internal type that signals the leader to send a heartbeat of MsgHeartBeat type.
	MsgProp                        // MsgProp proposes to append data to its log entries.
)
