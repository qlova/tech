// Package abi provides C ABI types for interoperability with shared C libraries.
package abi

import "unsafe"
import "C"

// Fixed width types.
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

// Func pointer in C with Go function type equivalent.
type Func[GoFunc any] Pointer[GoFunc]

// Opaque pointer type in C memory that cannot be
// dereferenced from Go.
type Opaque[T any] struct {
	_ [0]*T
	opaque
}

// String is a null-terminated string.
type String struct {
	ptr *Char
}

// String implements fmt.Stringer.
func (s String) String() string {
	return C.GoString((*C.char)(unsafe.Pointer(s.ptr)))
}

// NewString returns the given Go string in
// Go memory as a null-terminated C string.
func NewString(s string) String {
	if len(s) == 0 || s[len(s)-1] != 0 {
		s += "\x00"
	}
	return String{ptr: (*Char)(unsafe.Pointer(unsafe.StringData(s)))}
}

// Buffer is represented as a pointer to the first byte in
// the buffer and the length of the buffer. When passed to
// a C function, will be split into two subsequent arguments.
type Buffer struct {
	ptr *Uint8
	len Int
}

// Bytes returns the buffer as a Go byte slice.
func (s Buffer) Bytes() []byte {
	return C.GoBytes(unsafe.Pointer(s.ptr), C.int(s.len))
}

// Len returns the length of the buffer.
func (s Buffer) Len() Int {
	return s.len
}

// NewBuffer returns the given Go byte slice in Go memory
// as a C buffer.
func NewBuffer(s []byte) Buffer {
	return Buffer{&s[0], Int(len(s))}
}

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

type UnsafePointer unsafe.Pointer

type Pointer[T any] struct {
	_ [0]*T
	pointer
}

type opaque Uintptr

func (o opaque) Pointer() uintptr {
	return uintptr(o)
}

type pointer struct {
	val unsafe.Pointer
}

type IsPointer interface {
	Pointer() uintptr
}

func (p pointer) Pointer() uintptr {
	return uintptr(p.val)
}

func (p pointer) UnsafePointer() unsafe.Pointer {
	return p.val
}

func (s String) Pointer() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}
