package dyncall

// #include <dyncall.h>
// #include <dyncall_callback.h>
// #include <stdlib.h>
//
// extern DCsigchar bridge_callback(DCCallback*, DCArgs*, DCValue*, uintptr_t);
//
// DCCallback *goNewCallback(const DCsigchar * signature, uintptr_t userdata) {
//	return dcbNewCallback(signature, (DCCallbackHandler*)bridge_callback, (void*)userdata);
// }
import "C"
import "unsafe"

type Callback C.DCCallback

type CallbackHandler func(*Callback, *Args, unsafe.Pointer) rune

func NewCallback(sig Signature, handler CallbackHandler) *Callback {
	functions = append(functions, handler)

	s := C.CString(string(sig.Args) + ")" + string(sig.Returns))
	defer C.free(unsafe.Pointer(s))
	return (*Callback)(C.goNewCallback((*C.DCsigchar)(s), C.uintptr_t(len(functions))))
}

func (callback *Callback) Free() {
	C.dcbFreeCallback((*C.DCCallback)(callback))
}

type VM C.DCCallVM

func NewVM(size int) *VM {
	return (*VM)(C.dcNewCallVM(C.size_t(size)))
}

func (vm *VM) Reset() {
	C.dcReset((*C.DCCallVM)(vm))
}

func (vm *VM) Free() {
	C.dcFree((*C.DCCallVM)(vm))
}

func (vm *VM) PushBool(value bool) {
	var v C.DCbool
	if value {
		v = 1
	}
	C.dcArgBool((*C.DCCallVM)(vm), v)
}

func (vm *VM) PushInt8(value int8) {
	C.dcArgChar((*C.DCCallVM)(vm), C.DCchar(value))
}

func (vm *VM) PushInt16(value int16) {
	C.dcArgShort((*C.DCCallVM)(vm), C.DCshort(value))
}

func (vm *VM) PushInt32(value int32) {
	C.dcArgInt((*C.DCCallVM)(vm), C.DCint(value))
}

func (vm *VM) PushInt(value int) {
	C.dcArgLong((*C.DCCallVM)(vm), C.DClong(value))
}

func (vm *VM) PushInt64(value int64) {
	C.dcArgLongLong((*C.DCCallVM)(vm), C.DClonglong(value))
}

func (vm *VM) PushFloat32(value float32) {
	C.dcArgFloat((*C.DCCallVM)(vm), C.float(value))
}

func (vm *VM) PushFloat64(value float64) {
	C.dcArgDouble((*C.DCCallVM)(vm), C.double(value))
}

func (vm *VM) PushPointer(value unsafe.Pointer) {
	C.dcArgPointer((*C.DCCallVM)(vm), (C.DCpointer)(C.DCpointer(value)))
}

func (vm *VM) Call(address unsafe.Pointer) {
	C.dcCallVoid((*C.DCCallVM)(vm), (C.DCpointer)(unsafe.Pointer(address)))
}

func (vm *VM) CallBool(address unsafe.Pointer) bool {
	return C.dcCallBool((*C.DCCallVM)(vm), (C.DCpointer)(unsafe.Pointer(address))) != 0
}

func (vm *VM) CallInt8(address unsafe.Pointer) int8 {
	return int8(C.dcCallChar((*C.DCCallVM)(vm), (C.DCpointer)(unsafe.Pointer(address))))
}

func (vm *VM) CallInt16(address unsafe.Pointer) int16 {
	return int16(C.dcCallShort((*C.DCCallVM)(vm), (C.DCpointer)(unsafe.Pointer(address))))
}

func (vm *VM) CallInt32(address unsafe.Pointer) int32 {
	return int32(C.dcCallInt((*C.DCCallVM)(vm), (C.DCpointer)(unsafe.Pointer(address))))
}

func (vm *VM) CallInt(address unsafe.Pointer) int {
	return int(C.dcCallLong((*C.DCCallVM)(vm), (C.DCpointer)(unsafe.Pointer(address))))
}

func (vm *VM) CallInt64(address unsafe.Pointer) int64 {
	return int64(C.dcCallLongLong((*C.DCCallVM)(vm), (C.DCpointer)(unsafe.Pointer(address))))
}

func (vm *VM) CallFloat32(address unsafe.Pointer) float32 {
	return float32(C.dcCallFloat((*C.DCCallVM)(vm), (C.DCpointer)(unsafe.Pointer(address))))
}

func (vm *VM) CallFloat64(address unsafe.Pointer) float64 {
	return float64(C.dcCallDouble((*C.DCCallVM)(vm), (C.DCpointer)(unsafe.Pointer(address))))
}

func (vm *VM) CallPointer(address unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.dcCallPointer((*C.DCCallVM)(vm), (C.DCpointer)(unsafe.Pointer(address))))
}
