package cgo

/*
#include <dlfcn.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func dlopen(filename string) (handle unsafe.Pointer) {
	s := C.CString(filename)
	defer C.free(unsafe.Pointer(s))
	return C.dlopen(s, C.RTLD_NOW)
}

func dlerror() string {
	return C.GoString(C.dlerror())
}

func dlsym(handle unsafe.Pointer, symbol string) unsafe.Pointer {
	s := C.CString(symbol)
	defer C.free(unsafe.Pointer(s))
	return C.dlsym(handle, s)
}
