package raft

import (
	"errors"
	"math/rand"
	"sync"
	"time"

	pb "kbase.com/raft/raftpb"
)

// None is a placeholder node ID used when there is no leader.
const None uint64 = 0

const (
	StateFollower StateType = iota
	StateCandidate
	StateLeader
)

type StateType uint64

var stmp = [...]string{
	"StateFollower",
	"StateCandidate",
	"StateLeader",
	//	"StatePreCandidate",
}

type lockedRand struct {
	mu   sync.Mutex
	rand *rand.Rand
}

func (r *lockedRand) Intn(n int) int {
	r.mu.Lock()
	v := r.rand.Intn(n)
	r.mu.Unlock()
	return v
}

var globalRand = &lockedRand{
	rand: rand.New(rand.NewSource(time.Now().UnixNano())),
}

type Config struct {
	// ID is the identity of the local raft. ID can not be 0.
	ID uint64

	// If a follower dose not receive any message from leader of current
	// term before ElectionTick has elapsed, it will become candidate and
	// start an election. ElectionTick must be greater than HeartbeatTick.
	ElectionTick int

	HeartbeatTick int

	// TODO: add logger module
}

func (c *Config) validate() error {
	if c.ID == None {
		return errors.New("cannot use none ad id")
	}
	if c.HeartbeatTick <= 0 {
		return errors.New("heartbeat tick must be greater than 0")
	}
	if c.ElectionTick <= c.HeartbeatTick {
		return errors.New("election tick must be greater than heartbeat tick")
	}
	return nil
}

type raft struct {
	id uint64

	Term uint64
	Vote uint64

	state StateType

	msgs []pb.Message

	// the lead id
	lead uint64

	heartbeatTimeout int
	electionTimeout  int
	// randomizedElectionTimeout is a random number between
	// [electionTimeout, 2 * electionTimeout - 1]. It gets
	// reset when raft changes its state to follower or candidate.
	randomizedElectionTimeout int

	tick func()
}

func newRaft(c *Config) *raft {
	if err := c.validate(); err != nil {
		panic(err.Error())
	}

	r := &raft{
		id:               c.ID,
		lead:             None,
		electionTimeout:  c.ElectionTick,
		heartbeatTimeout: c.HeartbeatTick,
	}

	r.becomeFollower(r.Term, None)

	return r
}

// tickElection is run by followers and candidates after r.electionTimeout
func (r *raft) tickElection() {
	// TODO: be candidate and start election
}

func (r *raft) becomeFollower(term, lead uint64) {
	r.reset(term)
	r.tick = r.tickElection
	r.lead = lead
	r.state = StateFollower
}

func (r *raft) reset(term uint64) {
	if r.Term != term {
		r.Term = term
		r.Vote = None
	}
	r.lead = None

	r.resetRandomizedElectionTimeout()
}

func (r *raft) resetRandomizedElectionTimeout() {
	r.randomizedElectionTimeout = r.electionTimeout + globalRand.Intn(r.electionTimeout)
}

func (r *raft) send(m pb.Message) {
	r.msgs = append(r.msgs, m)
}

func (r *raft) Step(msg pb.Message) error {
	switch {
	case msg.Term == 0:
		// local message.
	case msg.Term > r.Term:
		if msg.Type == pb.MsgVote || msg.Type == pb.MsgPreVote {
			return nil
		}
		switch {
		case msg.Type == pb.MsgPreVote:
		case msg.Type == pb.MsgPreVoteResp:
		default:
			// TODO log
			if msg.Type == pb.MsgApp || msg.Type == pb.MsgHeartBeat || msg.Type == pb.MsgSnap {
				r.becomeFollower(msg.Term, msg.From)
			} else {
				r.becomeFollower(msg.Term, None)
			}

		}
	case msg.Term < r.Term:
	}

	switch msg.Type {
	case pb.MsgHup:
	default:
	}
	return nil
}
