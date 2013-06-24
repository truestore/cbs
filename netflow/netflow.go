package netflow

import (
	"fmt"
	"strconv"
)

type V5Record struct {
	DOctets, DPkts       uint32
	DstAs                uint16
	DstMask              uint8
	DstAddr              uint32
	DstPort              uint16
	EngineId, EngineType uint8
	ExAddr               uint32
	// https://code.google.com/p/pyflowtools/source/browse/flowtools.c#440
	First    uint64
	Input    uint16
	Last     uint64
	NextHop  uint32
	Output   uint16
	Prot     uint8
	SrcAs    uint16
	SrcMask  uint8
	SrcAddr  uint32
	SrcPort  uint16
	TcpFlags uint8
	Tos      uint8
}

func u32(s string) (uint32, error) {
	r, e := strconv.ParseUint(s, 10, 32)
	return uint32(r), e
}

func u16(s string) (uint16, error) {
	r, e := strconv.ParseUint(s, 10, 16)
	return uint16(r), e
}

func u8(s string) (uint8, error) {
	r, e := strconv.ParseUint(s, 10, 8)
	return uint8(r), e
}

func ut(s string) (uint64, error) {
	r, e := strconv.ParseFloat(s, 64)
	return uint64(r * 1000), e
}

func ParseCsv(record []string) (result *V5Record, err error) {
	if len(record) != 21 {
		return nil, fmt.Errorf("Failrecord: %s", len(record))
	}

	var r V5Record

	if r.DOctets, err = u32(record[0]); err != nil {
		return
	}

	if r.DPkts, err = u32(record[1]); err != nil {
		return
	}
	if r.DstAs, err = u16(record[2]); err != nil {
		return
	}
	if r.DstMask, err = u8(record[3]); err != nil {
		return
	}
	if r.DstAddr, err = u32(record[4]); err != nil {
		return
	}
	if r.DstPort, err = u16(record[5]); err != nil {
		return
	}
	if r.EngineId, err = u8(record[6]); err != nil {
		return
	}
	if r.EngineType, err = u8(record[7]); err != nil {
		return
	}
	if r.ExAddr, err = u32(record[8]); err != nil {
		return
	}
	if r.First, err = ut(record[9]); err != nil {
		return
	}
	if r.Input, err = u16(record[10]); err != nil {
		return
	}
	if r.Last, err = ut(record[11]); err != nil {
		return
	}
	if r.NextHop, err = u32(record[12]); err != nil {
		return
	}
	if r.Output, err = u16(record[13]); err != nil {
		return
	}
	if r.Prot, err = u8(record[14]); err != nil {
		return
	}
	if r.SrcAs, err = u16(record[15]); err != nil {
		return
	}
	if r.SrcMask, err = u8(record[16]); err != nil {
		return
	}
	if r.SrcAddr, err = u32(record[17]); err != nil {
		return
	}
	if r.SrcPort, err = u16(record[18]); err != nil {
		return
	}
	if r.TcpFlags, err = u8(record[19]); err != nil {
		return
	}
	if r.Tos, err = u8(record[20]); err != nil {
		return
	}

	result = &r
	return
}
