package writer

import (
	"code.google.com/p/lzma"
	"bytes"
	"encoding/binary"
)

type WriterType interface {

}

func LzmaCompress(size int64, data interface {}) ([]byte, error) {
	var buf bytes.Buffer
	wr := lzma.NewWriterSizeLevel(&buf, size, 9)
	err := binary.Write(wr, binary.BigEndian, data)
	wr.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
