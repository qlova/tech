package gpu

import (
	"math/bits"
	"reflect"
	"unsafe"
)

//Shader is a program that run on a GPU and can be compiled to a given gpu platform and version.
type Shader interface {
	//Draw(Mesh, Transform, Mode)
	CompileTo(platform, version string) (interface{}, error)

	//Variables returns the variable addresses of this shader.
	Variables() []interface{}
}

//materialWordSize determines how many words shader storage can be before it is placed on the heap.
const materialWordSize = 5

//Material is a material type.
type Material struct {
	gpuPointer Pointer

	//if possible, the variables of the material will be stored here.
	words [materialWordSize]uint

	pointer unsafe.Pointer
	shader  Shader
}

func (mat *Material) Pointer() Pointer {
	return mat.gpuPointer
}

type usp = unsafe.Pointer

//SetShader sets the material's shader.
func (mat *Material) SetShader(val Shader) {

	//just as C's memcpy, make sure the dest and src capcity >= len
	//Memcpy can not handle dest and src overlap condition
	memcpy := func(dest, src unsafe.Pointer, len uintptr) unsafe.Pointer {

		cnt := len >> 3
		var i uintptr = 0
		for i = 0; i < cnt; i++ {
			var pdest *uint64 = (*uint64)(usp(uintptr(dest) + uintptr(8*i)))
			var psrc *uint64 = (*uint64)(usp(uintptr(src) + uintptr(8*i)))
			*pdest = *psrc
		}
		left := len & 7
		for i = 0; i < left; i++ {
			var pdest *uint8 = (*uint8)(usp(uintptr(dest) + uintptr(8*cnt+i)))
			var psrc *uint8 = (*uint8)(usp(uintptr(src) + uintptr(8*cnt+i)))

			*pdest = *psrc
		}
		return dest
	}

	T := reflect.TypeOf(val)
	V := reflect.ValueOf(val)

	size := T.Elem().Size()

	if size < bits.UintSize*materialWordSize {
		memcpy(unsafe.Pointer(&mat.words), unsafe.Pointer(V.Pointer()), size)
		mat.shader = reflect.NewAt(T.Elem(), unsafe.Pointer(&mat.words)).Interface().(Shader)
		mat.pointer = unsafe.Pointer(&mat.words)
	} else {
		mat.shader = val
		mat.pointer = unsafe.Pointer(V.Pointer())
	}

	if context.Load != nil {
		p, err := context.Load(mat.shader)
		if err != nil {
			panic(err)
		}
		mat.gpuPointer = Pointer{p}
	}
}

//Shader returns the mutable shader of this material.
func (mat *Material) Shader() Shader {
	return mat.shader
}
