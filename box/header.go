package box

import (
	"bytes"
	"fmt"
	"reflect"
)

func headerFor(obj interface{}) []byte {
	var header bytes.Buffer

	rtype := reflect.TypeOf(obj)
	rtype = rtype.Elem()

	//messageHead points into the message data.
	//used for detecting padding.
	var messageHead uintptr

	switch rtype.Kind() {
	case reflect.Struct:

		for i := 0; i < rtype.NumField(); i++ {
			field := rtype.Field(i)

			if messageHead != field.Offset {
				panic("padding unimplemented")
			}

			messageHead += field.Type.Size()

			fmt.Println("offset ", field.Offset)

			switch field.Type.Kind() {
			case reflect.Float64:
				header.WriteByte(bits64 + uint8(i+1))
			}
		}

		header.WriteByte(0)

		return header.Bytes()
	default:
		panic("not supported yet")
	}
}

func readHeader(message []byte) (a, b []byte) {
	for i, v := range message {
		if v == 0 {
			return message[:i+1], message[i+1:]
		}
	}
	return message, nil
}
