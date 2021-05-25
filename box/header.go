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

			size := uint8(i + 1)
			if size > 31 {
				size = 0
			}

			var kind byte

			switch field.Type.Size() {
			case 1:
				kind = bits8
			case 2:
				kind = bits16
			case 4:
				kind = bits32
			case 8:
				kind = bits64
			default:
				panic("unsupported size: " + fmt.Sprint(field.Type.Size()))
			}

			//Write kind nibble and size nibble
			header.WriteByte(kind + size)
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
