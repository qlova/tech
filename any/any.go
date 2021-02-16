//Package any provides non-heap allocated unions of a given size with a fallback to the heap.
package any

import (
	"math/bits"
	"reflect"
	"unsafe"
)

type usp = unsafe.Pointer

//just as C's memcpy, make sure the dest and src capcity >= len
//Memcpy can not handle dest and src overlap condition
func memcpy(dest, src unsafe.Pointer, len uintptr) unsafe.Pointer {

	cnt := len >> 3
	var i uintptr = 0
	for i = 0; i < cnt; i++ {
		var pdest *uint64 = (*uint64)(usp(uintptr(dest) + uintptr(8*i)))
		var psrc *uint64 = (*uint64)(usp(uintptr(src) + uintptr(8*i)))
		*pdest = *psrc
	}
	left := len & 7
	for i = 0; i < left; i++ {
		var pdest *uint8 = (*uint8)(usp(uintptr(dest) + uintptr(8*cnt+i)))
		var psrc *uint8 = (*uint8)(usp(uintptr(src) + uintptr(8*cnt+i)))

		*pdest = *psrc
	}
	return dest
}

//Size5 can store any type (without pointers) with a size of less than 6 words without heap allocation.
type Size5 struct {
	words [5]uint

	//Value can be type switched.
	Value interface{}
}

//Set the value of the union.
func (s *Size5) Set(val interface{}) {
	T := reflect.TypeOf(val)
	V := reflect.ValueOf(val)

	size := T.Elem().Size()

	if size < bits.UintSize*5 {
		memcpy(unsafe.Pointer(&s.words), unsafe.Pointer(V.Pointer()), size)
		s.Value = reflect.NewAt(T.Elem(), unsafe.Pointer(&s.words)).Interface()
	} else {
		s.Value = val
	}
}
