package writer

import (
	"github.com/truestore/cbs/cbs_proto"

	"code.google.com/p/goprotobuf/proto"
)

type Uint32Writer struct {
	data []uint32
}

func (u *Uint32Writer) Clear() {
	u.data = u.data[0:0]
}

func (u *Uint32Writer) IsFull() bool {
	return len(u.data) == cap(u.data)
}

func (u *Uint32Writer) Flush() (header *cbs_proto.Header, result []byte, err error) {
	if result, err = LzmaCompress(int64(len(u.data))*4, u.data); err != nil {
		return
	}

	header = &cbs_proto.Header{
		NumRows:             proto.Uint64(uint64(u.Len())),
		BlockSize:           proto.Uint64(uint64(len(u.data) * 4)),
		CompressedBlockSize: proto.Uint64(uint64(len(result))),
	}

	u.data = u.data[0:0]
	return
}

func NewUint32Writer(size int) *Uint32Writer {
	return &Uint32Writer{
		data: make([]uint32, 0, size),
	}
}

func (u *Uint32Writer) Len() int {
	return len(u.data)
}

func (u *Uint32Writer) Append(data interface{}) {
	if x, ok := data.(uint32); ok {
		u.data = append(u.data, x)
	}
}
