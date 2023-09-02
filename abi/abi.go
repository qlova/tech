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
	Enum              Int
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

type Func[T any] Pointer

type Pointer struct {
	pointer
}

type pointer Uintptr

type IsPointer interface {
	Pointer() uintptr
}

func (p pointer) Pointer() uintptr {
	return uintptr(p)
}

type String struct {
	ptr *Char
}

func (s String) String() string {
	return C.GoString((*C.char)(unsafe.Pointer(s.ptr)))
}

func (s String) Pointer() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func NewString(s string) String {
	if len(s) == 0 || s[len(s)-1] != 0 {
		s += "\x00"
	}
	return String{ptr: (*Char)(unsafe.Pointer(unsafe.StringData(s)))}
}

type Buffer struct {
	ptr *Uint8
	len Int
}

func (s Buffer) Bytes() []byte {
	return C.GoBytes(unsafe.Pointer(s.ptr), C.int(s.len))
}

func NewBuffer(s []byte) Buffer {
	return Buffer{&s[0], Int(len(s))}
}
