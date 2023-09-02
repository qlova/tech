package abi

type (
	Bool             = bool
	Char             = int8
	CharSigned       = int8
	CharUnsigned     = uint8
	Short            = int16
	ShortSigned      = int16
	ShortUnsigned    = uint16
	Int              = int32
	IntSigned        = int32
	Enum             = int32
	IntUnsigned      = uint32
	Long             = int64
	LongSigned       = int64
	LongUnsigned     = uint64
	LongLong         = int64
	LongLongSigned   = int64
	LongLongUnsigned = uint64

	Decimal32      = [4]byte
	Decimal64      = [8]byte
	Decimal128     = [16]byte
	Int128         = [16]byte
	Int128Signed   = [16]byte
	Int128Unsigned = [16]byte
	M64            = [8]byte
	M128           = [16]byte
	M256           = [32]byte

	Size = uintptr

	Pointer    = struct{ uintptr }
	Float      = float32
	Double     = float64
	DoubleLong = [16]byte

	ComplexDouble     = [2]float64
	ComplexFloat      = [2]float32
	ComplexDoubleLong = [2][16]byte

	ImaginaryDouble     = float64
	ImaginaryDoubleLong = [16]byte
	ImaginaryFloat      = float32
)
