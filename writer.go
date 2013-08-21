package cbs

import (
	"io"

	"code.google.com/p/goprotobuf/proto"
	"github.com/truestore/cbs/cbs_proto"
)

type Writer interface {
	Append(value interface{}) error
	Flush() error
}

type writer struct {
	w     io.Writer
	conv  byteConverter
	data  []byte
	nrows uint64
}

// check that writer implements Writer interface
var _ Writer = &writer{}

type byteConverter func(v interface{}) []byte

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
		w.Flush()
	}

	return nil
}

func (w *writer) Flush() error {
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

func NewUint16Writer(w io.Writer) Writer {
	return newWriter(w, uint16ToBytes)
}

func NewUint32Writer(w io.Writer) Writer {
	return newWriter(w, uint32ToBytes)
}
