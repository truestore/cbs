package cbs

import (
	"bytes"
	"encoding/binary"
)

// There is better implementation!
func uint16ToBytes(v interface{}) []byte {
	return anyToBytes(v)
}

func uint32ToBytes(v interface{}) []byte {
	return anyToBytes(v)
}

func anyToBytes(v interface{}) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, v)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}
