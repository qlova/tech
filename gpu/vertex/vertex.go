/*
	Package vertex provides functions for working with vertex data.
*/
package vertex

import (
	"reflect"
	"sort"
	"unsafe"

	"qlova.tech/rgb"
	"qlova.tech/rgb/rgba"
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
	kind   reflect.Kind
	count  int
}

type data interface {
	//basic types.
	//We can accept single values or up to 4 values.
	~bool | [1]bool | [2]bool | [3]bool | [4]bool |
		~int | [1]int | [2]int | [3]int | [4]int |
		~int8 | [1]int8 | [2]int8 | [3]int8 | [4]int8 |
		~int16 | [1]int16 | [2]int16 | [3]int16 | [4]int16 |
		~int32 | [1]int32 | [2]int32 | [3]int32 | [4]int32 |
		~int64 | [1]int64 | [2]int64 | [3]int64 | [4]int64 |
		~uint | [1]uint | [2]uint | [3]uint | [4]uint |
		~uint8 | [1]uint8 | [2]uint8 | [3]uint8 | [4]uint8 |
		~uint16 | [1]uint16 | [2]uint16 | [3]uint16 | [4]uint16 |
		~uint32 | [1]uint32 | [2]uint32 | [3]uint32 | [4]uint32 |
		~uint64 | [1]uint64 | [2]uint64 | [3]uint64 | [4]uint64 |
		~float32 | [1]float32 | [2]float32 | [3]float32 | [4]float32 |
		~float64 | [1]float64 | [2]float64 | [3]float64 | [4]float64 |

		~complex64 | [1]complex64 | [2]complex64 |
		~complex128 | [1]complex128 | [2]complex128 |
		rgb.Color | rgba.Color
}

//Data returns the given data slice as a Slice.
func Data[T data](data []T) Slice {
	var elem = &data[0]
	var rtype = reflect.TypeOf(elem)
	var count = 1
	if rtype.Kind() == reflect.Array {
		count = rtype.Len()
	}

	return Slice{
		kind:   rtype.Kind(),
		count:  count,
		buffer: unsafe.Slice((*byte)(unsafe.Pointer(elem)), unsafe.Sizeof(elem)*uintptr(len(data))),
	}
}
