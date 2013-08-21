package cbs

import (
	"code.google.com/p/goprotobuf/proto"
	"github.com/truestore/cbs/cbs_proto"
)

type block struct {
	Header *cbs_proto.Header
	Data   []byte
}

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
