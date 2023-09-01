package ffi

import "C"
import "unsafe"

type String struct {
	ptr *byte
}

func (s String) String() string {
	return C.GoString((*C.char)(unsafe.Pointer(s.ptr)))
}

func NewString(s string) String {
	if len(s) == 0 || s[len(s)-1] != 0 {
		s += "\x00"
	}
	return String{unsafe.StringData(s)}
}
