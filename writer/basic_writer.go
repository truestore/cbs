package writer

import (
	"encoding/binary"
)

import (
	"github.com/truestore/cbs/cbs_proto"
	"github.com/truestore/cbs/types"

	"code.google.com/p/goprotobuf/proto"
)

type BasicWriter struct {
	columnType types.ColumnType
	buf        []byte
}

// Writer for basic types
func NewBasicWriter(columnType types.ColumnType, size int) *BasicWriter {
	bsize := size * columnType.Size()
	return &BasicWriter{columnType: columnType, buf: make([]byte, 0, bsize)}
}

func (self *BasicWriter) Append(data interface{}) {
	check := func(a, b types.ColumnType) {
		if a != b {
			panic(TypeMismatch)
		}
	}

	var b [8]byte
	var wb []byte
	order := binary.BigEndian

	switch v := data.(type) {
	case int8:
		check(self.columnType, types.Int8)
		wb = b[:1]
		wb[0] = byte(v)
	case uint8:
		check(self.columnType, types.Uint8)
		wb = b[:1]
		wb[0] = byte(v)
	case int16:
		check(self.columnType, types.Int16)
		wb = b[:2]
		order.PutUint16(wb, uint16(v))
	case uint16:
		check(self.columnType, types.Uint16)
		wb = b[:2]
		order.PutUint16(wb, uint16(v))
	case int32:
		check(self.columnType, types.Int32)
		wb = b[:4]
		order.PutUint32(wb, uint32(v))
	case uint32:
		check(self.columnType, types.Uint32)
		wb = b[:4]
		order.PutUint32(wb, uint32(v))
	case int64:
		check(self.columnType, types.Int64)
		wb = b[:8]
		order.PutUint64(wb, uint64(v))
	case uint64:
		check(self.columnType, types.Uint64)
		wb = b[:8]
		order.PutUint64(wb, uint64(v))
	default:
		panic(TypeUnknown)
	}

	self.buf = append(self.buf, wb...)
}

func (self *BasicWriter) Flush() (header *cbs_proto.Header, result []byte, err error) {
	if result, err = LzmaCompress(int64(len(self.buf)), self.buf); err != nil {
		return
	}

	header = &cbs_proto.Header{
		NumRows:             proto.Uint64(uint64(self.Len())),
		BlockSize:           proto.Uint64(uint64(len(self.buf))),
		CompressedBlockSize: proto.Uint64(uint64(len(result))),
	}

	self.buf = self.buf[0:0]
	return
}

func (self *BasicWriter) Len() int {
	return len(self.buf) / self.columnType.Size()
}

func (self *BasicWriter) IsFull() bool {
	return len(self.buf) == cap(self.buf)
}

func (self *BasicWriter) Clear() {
	self.buf = self.buf[0:0]
}
