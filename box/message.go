package box

import (
	"reflect"
	"unsafe"
)

func readMessage(message []byte, obj interface{}) error {
	rvalue := reflect.ValueOf(obj)

	size := rvalue.Type().Elem().Size()

	memcpy(
		unsafe.Pointer(rvalue.Pointer()),
		unsafe.Pointer(&message[0]),
		size)

	return nil
}

func messageFor(obj interface{}) []byte {
	rvalue := reflect.ValueOf(obj)

	size := rvalue.Type().Elem().Size()

	var buffer = make([]byte, size)

	memcpy(
		unsafe.Pointer(&buffer[0]),
		unsafe.Pointer(rvalue.Pointer()),
		size)

	return buffer
}
