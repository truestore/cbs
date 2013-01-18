package netflow

type V5Record struct {
	DOctets, DPkts uint32
	DstAs uint16
	DstMask uint8
	DstAddr uint32
	DstPort uint16
	EngineId, EngineType uint8
	ExAddr uint32
	// https://code.google.com/p/pyflowtools/source/browse/flowtools.c#440
	First uint64
	Input uint16
	Last uint64
	NextHop uint32
	Output uint16
	Prot uint8
	SrcAs uint16
	SrcMask uint8
	SrcAddr uint32
	SrcPort uint16
	TcpFlags uint8
	Tos uint8
}
