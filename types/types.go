package types

type ColumnType int

const (
	Int8 = ColumnType(iota)
	Uint8
	Int16
	Uint16
	Int32
	Uint32
	Int64
	Uint64
)

func (self ColumnType) Size() int {
	switch self {
	case Int8:
		fallthrough
	case Uint8:
		return 1
	case Int16:
		fallthrough
	case Uint16:
		return 2
	case Int32:
		fallthrough
	case Uint32:
		return 4
	case Int64:
		fallthrough
	case Uint64:
		return 8
	default:
		panic("unknown type")
	}
}
