package rpc

import (
	"time"

	"kbase.com/pkg/rpc/codec"
)

const MagicNumber = 0x3bef5c

type Option struct {
	CodecType      codec.Type
	MagicNumber    int
	ConnectTimeout time.Duration // 0 means no limit
	HandleTimeout  time.Duration
}

var DefaultOption = &Option{
	CodecType:      codec.GobType,
	MagicNumber:    MagicNumber,
	ConnectTimeout: time.Second * 10,
}
