package abi

import "unsafe"
import "C"

type (
	Int8    = int8
	Int16   = int16
	Int32   = int32
	Int64   = int64
	Uint8   = uint8
	Uint16  = uint16
	Uint32  = uint32
	Uint64  = uint64
	Uintptr = uintptr
)

type (
	Error             Int
	FloatException    Int
	FloatRoundingMode Int
	LocaleCategory    Int
	FloatClass        Int
	Signal            Int
	BufferMode        Int
	SeekMode          Int
	TimeType          Int
)

func (s String) String() string {
	return C.GoString((*C.char)(unsafe.Pointer(uintptr(s.uint64))))
}
