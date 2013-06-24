package writer

import (
	"github.com/truestore/cbs/types"
)

func NewUint16Writer(size int) Writer {
	return NewBasicWriter(types.Uint16, size)
}
