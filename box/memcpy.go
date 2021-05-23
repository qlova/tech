package box

import "unsafe"

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
