package writer

import (
	"cbs/cbs_proto"

	"bytes"
	"encoding/binary"

	"code.google.com/p/goprotobuf/proto"
    "code.google.com/p/lzma"
)

type Uint32Writer struct {
	data []uint32
}

func (u *Uint32Writer) Flush() (header *cbs_proto.Header, result []byte, err error) {
	var buf bytes.Buffer

	w := lzma.NewWriterSizeLevel(&buf, int64(len(u.data)) * 4, 9)
	if err = binary.Write(w, binary.BigEndian, u.data); err != nil {
		w.Close()
		return
	}
	w.Close()


	header = &cbs_proto.Header{
		NumRows: proto.Uint64(uint64(len(u.data))),
		BlockSize: proto.Uint64(uint64(len(u.data) * 4)),
		CompressedBlockSize: proto.Uint64(uint64(buf.Len())),
	}

	result = buf.Bytes()

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

func (u *Uint32Writer) Append(data interface {}) {
	if x, ok := data.(uint32); ok {
		u.data = append(u.data, x)
	}
}
