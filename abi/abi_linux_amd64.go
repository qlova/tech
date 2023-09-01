package abi

type (
	Bool             = bool
	Char             = int8
	SignedChar       = int8
	UnsignedChar     = uint8
	Short            = int16
	SignedShort      = int16
	UnsignedShort    = uint16
	Int              = int32
	SignedInt        = int32
	Enum             = int32
	UnsignedInt      = uint32
	Long             = int64
	SignedLong       = int64
	LongLong         = int64
	SignedLongLong   = int64
	UnsignedLong     = uint64
	UnsignedLongLong = uint64
	Int128           = [16]byte
	SignedInt128     = [16]byte
	UnsignedInt128   = [16]byte
	Pointer          = struct{ uintptr }
	Float            = float32
	Double           = float64
	LongDouble       = [16]byte
	Float128         = [16]byte
	Decimal32        = [4]byte
	Decimal64        = [8]byte
	Decimal128       = [16]byte
	M64              = [8]byte
	M128             = [16]byte
	M256             = [32]byte
)
