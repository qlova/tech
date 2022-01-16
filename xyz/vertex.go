package xyz

import (
	"reflect"
	"unsafe"
)

// Pointer is a low level description of where to
// find a vertex attribute inside of a Vertices Buffer.
type Pointer struct {
	Attribute                     string
	Kind                          reflect.Kind
	Buffer, Count, Offset, Stride uint
}

type Buffer []byte

func NewBuffer[T any](slice []T) Buffer {
	return unsafe.Slice((*byte)(unsafe.Pointer(&slice[0])), uintptr(len(slice))*unsafe.Sizeof(slice[0]))
}

type Vertices interface {
	Length() int
	Layout() []Pointer
	Buffers() []Buffer
	Indexed() (int, bool)
}
