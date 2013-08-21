package cbs

import (
	"bytes"
	"testing"
)

func BenchmarkUint16Append(b *testing.B) {
	var buf bytes.Buffer
	w := NewUint16Writer(&buf)

	for i := 0; i < b.N; i++ {
		w.Append(uint16(4663))
	}

	w.Flush()

	b.Logf("appended %d values, column raw size: %d bytes", b.N, buf.Len())
}
