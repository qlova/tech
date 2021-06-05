package box

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
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
				header.WriteByte(padding + byte(field.Offset-messageHead))
				messageHead = field.Offset
			}

			messageHead += field.Type.Size()

			//box is the Nth box that this field
			//will be placed in.
			box := uint8(i + 1)
			if box > 31 {
				box = 0
			}

			switch field.Type.Kind() {
			case reflect.Chan, reflect.Ptr, reflect.Map,
				reflect.String, reflect.Interface, reflect.UnsafePointer:
				panic("pointers not supported yet")
			}

			var size byte

			switch field.Type.Size() {
			case 1:
				size = bits8
			case 2:
				size = bits16
			case 4:
				size = bits32
			case 8:
				size = bits64
			default:
				panic("unsupported size: " + fmt.Sprint(field.Type.Size()))
			}

			//Write kind nibble and size nibble
			header.WriteByte(size + box)
		}

		if messageHead != rtype.Size() {
			header.WriteByte(padding + byte(rtype.Size()-messageHead))
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

func sprintHeader(header []byte) string {
	var s strings.Builder
	for _, b := range header {
		size := b & 0xF0
		box := b & 0x0F

		if size != command {

			if size == padding {
				fmt.Fprintf(&s, "padding(%vbits), ", box*8)
				continue
			} else {
				fmt.Fprintf(&s, "%v: ", box)
			}
		} else {
			s.WriteString("cmd.")
			switch box {
			case message:
				s.WriteString("message ")

			default:
				s.WriteString("?")
			}
			continue
		}

		switch size {
		case bits8:
			s.WriteString("bits8, ")
		case bits16:
			s.WriteString("bits16, ")
		case bits32:
			s.WriteString("bits32, ")
		case bits64:
			s.WriteString("bits64, ")
		case opening:
			s.WriteString("{")
		case exactly:
			s.WriteString("[")
		default:
			s.WriteString("?")
		}

	}
	return s.String()
}
