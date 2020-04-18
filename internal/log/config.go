package log

import (
	"github.com/hashicorp/raft"
)

type Config struct {
	Segment
	Raft
}

type Segment struct {
	MaxStoreBytes uint64
	MaxIndexBytes uint64
	InitialOffset uint64
}

type Raft struct {
	raft.Config
	StreamLayer *StreamLayer
	Bootstrap   bool
}
