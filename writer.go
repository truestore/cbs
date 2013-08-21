package cbs

import (
	"io"

	"code.google.com/p/goprotobuf/proto"
	"github.com/truestore/cbs/cbs_proto"
)

type Writer interface {
	Append(value interface{}) error
}

func NewUint16Writer(w io.Writer) Writer {
	return newWriter(w, uint16ToBytes)
}

func NewUint32Writer(w io.Writer) Writer {
	return newWriter(w, uint32ToBytes)
}

type writer struct {
	w     io.Writer
	conv  byteConverter
	data  []byte
	nrows uint64
}

func newWriter(w io.Writer, conv byteConverter) *writer {
	return &writer{w: w, conv: conv}
}

var blockSize = uint64(4096)

func (w *writer) Append(v interface{}) error {
	if v != nil {
		bs := w.conv(v)
		w.data = append(w.data, bs...)
		w.nrows++
	} else {
		panic("TODO: bitmap")
	}

	if w.nrows >= blockSize {
		w.flush()
	}

	return nil
}

func (w *writer) flush() error {
	block, err := packBlock(w.nrows, w.data)
	if err != nil {
		return err
	}

	msg, err := proto.Marshal(block.Header)
	if err != nil {
		return err
	}

	if _, err = w.w.Write(msg); err != nil {
		return err
	}

	if _, err = w.w.Write(block.Data); err != nil {
		return err
	}

	w.nrows = 0
	w.data = w.data[0:0]

	return nil
}

type block struct {
	Header *cbs_proto.Header
	Data   []byte
}

type byteConverter func(v interface{}) []byte

func packBlock(nrows uint64, data []byte) (*block, error) {
	compressed, err := lzmaCompress(int64(len(data)), data)
	if err != nil {
		return nil, err
	}

	header := cbs_proto.Header{
		NumRows:             proto.Uint64(nrows),
		BlockSize:           proto.Uint64(uint64(len(data))),
		CompressedBlockSize: proto.Uint64(uint64(len(compressed))),
	}

	return &block{&header, compressed}, nil
}
