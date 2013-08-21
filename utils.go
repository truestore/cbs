package writer

import (
	"bytes"
	"code.google.com/p/lzma"
)

func lzmaCompress(size int64, data []byte) ([]byte, error) {
	var buf bytes.Buffer
	wr := lzma.NewWriterSizeLevel(&buf, size, 9)
	if _, err := wr.Write(data); err != nil {
		return nil, err
	}
	if err := wr.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
