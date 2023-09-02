package ffi

/*
#include <dlfcn.h>
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

func dlopen(filename string) (handle unsafe.Pointer) {
	s := C.CString(filename + "\x00")
	defer C.free(unsafe.Pointer(s))
	return C.dlopen(s, C.RTLD_NOW)
}

func dlerror() string {
	return C.GoString(C.dlerror())
}

func dlsym(handle unsafe.Pointer, symbol string) unsafe.Pointer {
	s := C.CString(symbol + "\x00")
	defer C.free(unsafe.Pointer(s))
	return C.dlsym(handle, s)
}
