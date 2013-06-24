package writer

import (
	"bytes"
	"code.google.com/p/lzma"
)

func LzmaCompress(size int64, data []byte) ([]byte, error) {
	var buf bytes.Buffer
	wr := lzma.NewWriterSizeLevel(&buf, size, 9)
	_, err := wr.Write(data)
	wr.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
