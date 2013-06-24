package writer

import (
	"github.com/truestore/cbs/types"
)

func NewUint32Writer(size int) Writer {
	return NewBasicWriter(types.Uint32, size)
}
