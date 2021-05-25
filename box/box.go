//Package box provides a marshalling and unmarshaling of binary object explanations.
package box

import (
	"bytes"
	"fmt"
	"reflect"
	"time"
)

const (
	hasSchema = 1 << iota
)

//5bit command
const (
	//message denotes the end of the
	//header and the start of the
	//message.
	message = 0x00

	//closing closes the previous opening field.
	closing = 0x01

	//typedef defines a type that can be
	//referenced.
	typedef = 0x02

	//
	decrypt = 0x03
	deflate = 0x04

	version = 0x06
	timeout = 0x07
	passkey = 0x08

	creator = 0x09
	created = 0x0A
	modtime = 0x0B

	failure = 0x0C

	timefix
)

//3bit structure
const (
	//command
	command = 0x00

	//fixed sizes.
	bits8  = 0x10
	bits16 = 0x20
	bits32 = 0x30
	bits64 = 0x40

	//opening denotes the beginning of
	//a sub-structure with a different
	//field-number mapping.
	opening = 0x50

	//exactly overrides the field number
	//to indicate how many repetitions
	//of the next header byte there will
	//be.
	exactly = 0x60

	padding = 0x70
)

//schema entry
const (
	//command = 0x00

	boolean = 1
	natural = 2
	integer = 3
	ieee754 = 4
	complex = 5

	//exactly = 6
	pointer = 7
	mapping = 8
	listing = 9
	//message = 10

	unicode = 11
	bintime = 12

	//catchall types
	bindata = 13
	dynamic = 14
	unknown = 15
)

type Header struct {
	Creator string
	Version string

	PassKey string

	Timeout time.Duration
	Created time.Time
	ModTime time.Time
}

type header []byte

var headers = make(map[reflect.Type]header)

func Marshal(obj interface{}) ([]byte, error) {
	var buffer bytes.Buffer

	buffer.Write(headerFor(obj))
	buffer.Write(messageFor(obj))

	return buffer.Bytes(), nil
}

func Unmarshal(data []byte, obj interface{}) error {

	var header, message = readHeader(data)

	fmt.Println(sprintHeader(header))

	if bytes.Compare(header, headerFor(obj)) != 0 {
		return fmt.Errorf("cannot decode other language yet")
	}

	return readMessage(message, obj)
}
