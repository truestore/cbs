package writer

import (
	"errors"

	"github.com/truestore/cbs/cbs_proto"
)

type Writer interface {
	Append(data interface{})
	Flush() (header *cbs_proto.Header, result []byte, err error)
	IsFull() bool
	Len() int
	Clear()
}

var (
	TypeMismatch = errors.New("Type mismatch")
	TypeUnknown  = errors.New("Type unknown")
)
