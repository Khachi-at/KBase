package codec

import "io"

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implement
)

type NewCodecFunc func(io.ReadWriteCloser) Codec

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
