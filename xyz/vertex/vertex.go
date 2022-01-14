/*
	Package vertex provides functions for working with vertex data.
*/
package vertex

import (
	"reflect"
	"sort"
	"unsafe"
)

// Hint for how the GPU should handle a vertex array.
type Hint uint64

// Hints.
const (
	// Read enables reading from the vertex array.
	Read Hint = 1 << iota

	// Write enables rewriting to the vertex array.
	Write

	// Stream should be set if the vertex array
	// should be updated every frame.
	Stream
)

//Attribute is a named attribute of a vertex.
type Attribute string

// Common attributes, use these in your shaders
// for maximum compatibility.
const (
	Indicies Attribute = "indicies"

	Position Attribute = "position"
	Normal   Attribute = "normal"
	UV       Attribute = "uv"
	Color    Attribute = "color"
	Weight   Attribute = "weight"
	Joint    Attribute = "joint"
)

//Reader TODO
type Reader interface{}

// Buffer containing vertex data.
type Buffer []byte

// Array is an array of vertices.
type Array interface {
	Length() int
	Layout() []AttributePointer
	Buffers() []Buffer
}

// AttributePointer is a low level description of where to
// find an attribute inside of an Array.
type AttributePointer struct {
	Attribute                     Attribute
	Kind                          reflect.Kind
	Buffer, Count, Offset, Stride uint
}

// Attributes can be used to specify vertex data by
// specifying the vertex data for each attribute.
// Each Attribute's Slice must be the same length.
type Attributes map[Attribute]Slice

// Length returns the number of vertices.
func (a Attributes) Length() int {
	var keys = make([]string, 0, len(a))
	for k := range a {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)

	var length = -1
	for _, key := range keys {
		attr := a[Attribute(key)]
		l := len(attr.buffer) / attr.size
		if length == -1 {
			length = l
		} else {
			if l < length {
				length = l
			}
		}
	}

	if length == -1 {
		return 0
	}

	return length
}

// Layout returns the vertex attribute layout for
// the Attributes, each attribute is tightly packed
// in a seperate buffer.
func (a Attributes) Layout() []AttributePointer {
	var pointers = make([]AttributePointer, 0, len(a))

	var keys = make([]string, 0, len(a))
	for k := range a {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)

	for i, key := range keys {
		attribute := Attribute(key)
		data := a[attribute]

		pointers = append(pointers, AttributePointer{
			Attribute: attribute,
			Kind:      data.kind,
			Buffer:    uint(i),
			Count:     uint(data.count),
		})
	}

	return pointers
}

// Buffers returns the buffers used to store the
// vertex data of the Attributes.
func (a Attributes) Buffers() []Buffer {
	var buffers = make([]Buffer, 0, len(a))

	var keys = make([]string, 0, len(a))
	for k := range a {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)

	for _, key := range keys {
		attribute := Attribute(key)
		data := a[attribute]

		buffers = append(buffers, data.buffer)
	}

	return buffers
}

//Slice is a typed slice of vertex data.
//Sliced on a single Attribute.
type Slice struct {
	buffer Buffer
	size   int //element size.
	kind   reflect.Kind
	count  int
}

type data interface {
	//basic types.
	//We can accept single values or up to 4 values.
	~bool | ~[1]bool | ~[2]bool | ~[3]bool | ~[4]bool |
		~int8 | ~[1]int8 | ~[2]int8 | ~[3]int8 | ~[4]int8 |
		~int16 | ~[1]int16 | ~[2]int16 | ~[3]int16 | ~[4]int16 |
		~int32 | ~[1]int32 | ~[2]int32 | ~[3]int32 | ~[4]int32 |
		~uint8 | ~[1]uint8 | ~[2]uint8 | ~[3]uint8 | ~[4]uint8 |
		~uint16 | ~[1]uint16 | ~[2]uint16 | ~[3]uint16 | ~[4]uint16 |
		~uint32 | ~[1]uint32 | ~[2]uint32 | ~[3]uint32 | ~[4]uint32 |
		~float32 | ~[1]float32 | ~[2]float32 | ~[3]float32 | ~[4]float32 |
		~float64 | ~[1]float64 | ~[2]float64 | ~[3]float64 | ~[4]float64 |

		~complex64 | ~[1]complex64 | ~[2]complex64
}

//Data returns the given data slice as a Slice.
func Data[T data](data []T) Slice {
	var elem = &data[0]
	var rtype = reflect.TypeOf(elem)
	var kind = rtype.Elem().Kind()
	var count = 1
	if kind == reflect.Array {
		count = rtype.Elem().Len()
		kind = rtype.Elem().Elem().Kind()
	}

	return Slice{
		kind:   kind,
		count:  count,
		size:   int(unsafe.Sizeof(elem)),
		buffer: unsafe.Slice((*byte)(unsafe.Pointer(elem)), unsafe.Sizeof(data[0])*uintptr(len(data))),
	}
}
