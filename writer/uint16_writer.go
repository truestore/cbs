package writer

import (
	"github.com/truestore/cbs/cbs_proto"

	"code.google.com/p/goprotobuf/proto"
)

type Uint16Writer struct {
	data []uint16
}

func (u *Uint16Writer) Clear() {
	u.data = u.data[0:0]
}

func (u *Uint16Writer) IsFull() bool {
	return len(u.data) == cap(u.data)
}

func (u *Uint16Writer) Flush() (header *cbs_proto.Header, result []byte, err error) {
	if result, err = LzmaCompress(int64(len(u.data))*2, u.data); err != nil {
		return
	}

	header = &cbs_proto.Header{
		NumRows:             proto.Uint64(uint64(u.Len())),
		BlockSize:           proto.Uint64(uint64(len(u.data) * 2)),
		CompressedBlockSize: proto.Uint64(uint64(len(result))),
	}

	u.data = u.data[0:0]
	return
}

func NewUint16Writer(size int) *Uint16Writer {
	return &Uint16Writer{
		data: make([]uint16, 0, size),
	}
}

func (u *Uint16Writer) Len() int {
	return len(u.data)
}

func (u *Uint16Writer) Append(data interface{}) {
	if x, ok := data.(uint16); ok {
		u.data = append(u.data, x)
	}
}
