package opengl46

/*
typedef unsigned int uint;

typedef  struct {
	uint count;
	uint instanceCount;
	uint firstIndex;
	uint baseVertex;
	uint baseInstance;
} DrawElementsIndirectCommand;

typedef  struct {
	uint  count;
	uint  instanceCount;
	uint  first;
	uint  baseInstance;
} DrawArraysIndirectCommand;
*/
import "C"
import (
	"math"
	"reflect"
	"unsafe"

	"qlova.tech/gpu"

	"github.com/go-gl/gl/v4.6-core/gl"
)

type drawElementsIndirectCommand = C.DrawElementsIndirectCommand
type drawArraysIndirectCommand = C.DrawArraysIndirectCommand

type state struct {
	shader uint32
	vao    uint32

	indexed bool

	options gpu.DrawOptions
}

type bucket struct {
	state state

	id uint16

	pointer uint32

	drawArrays   []drawArraysIndirectCommand
	drawElements []drawElementsIndirectCommand

	transform  uint32
	transforms []gpu.Transform

	variables []variable
}

var buckets []bucket

func newBucket(s state) int {
	var b bucket
	b.state = s
	b.id = uint16(len(buckets))
	buckets = append(buckets, b)

	return len(buckets) - 1
}

var bucketMap = make(map[state]int)

//variable specific to a single draw call.
type variable struct {
	Name string

	//uniform
	pointer uint32

	Type int

	Address unsafe.Pointer
	Length  int
	Buffer  []uint32
}

func (v *variable) push() {
	for i := 0; i < v.Length; i++ {
		offset := uintptr(i * 4)
		v.Buffer = append(v.Buffer, *(*uint32)(unsafe.Pointer(uintptr(v.Address) + offset)))
	}
	//Padding for vec3 due to glsl alignment.
	if v.Length == 3 {
		v.Buffer = append(v.Buffer, math.Float32bits(1))
	}
}

const (
	pointerTypeFloat = iota + 1
	pointerTypeUint64
)

func loadVariables(shader gpu.Program) []variable {
	var rvalue = reflect.ValueOf(shader)
	if rvalue.Type().Kind() != reflect.Ptr {
		panic("shader must have pointer reciever")
	}

	rvalue = rvalue.Elem()
	if rvalue.Type().Kind() != reflect.Struct {
		return nil
	}

	var variables []variable
	for _, v := range shader.Variables() {
		field := reflect.ValueOf(v).Elem()
		T := field.Type()

		switch T {
		case reflect.TypeOf(gpu.Vec2{}),
			reflect.TypeOf(gpu.Vec3{}),
			reflect.TypeOf(gpu.Vec4{}),
			reflect.TypeOf(gpu.Texture{}):

			var p variable
			p.Address = unsafe.Pointer(field.UnsafeAddr())
			p.Length = 1

			if T.Kind() == reflect.Array {
				p.Length = field.Len()
				T = T.Elem()
			}

			if T == reflect.TypeOf(gpu.Texture{}) {
				p.Type = pointerTypeUint64
				p.Length *= 2
			} else {

				switch T.Kind() {
				case reflect.Uint64:
					p.Type = pointerTypeUint64
					p.Length *= 2
				case reflect.Float32:
					p.Type = pointerTypeFloat
				default:
					panic("material variable type not implemented: " + T.String())
				}
			}

			variables = append(variables, p)
		}
	}

	return variables
}

func draw(mesh gpu.Mesh, t gpu.Transform, options gpu.DrawOptions) error {
	key := state{
		indexed: mesh.Indexed(),
		shader:  uint32(mesh.Material.Pointer().Value()),
		vao:     uint32(mesh.Pointer().Value()),
		options: options,
	}

	//Figure out the bucket that this drawcall is using.
	index, ok := bucketMap[key]
	if !ok {
		index = newBucket(key)
		bucketMap[key] = index
	}

	bucket := &buckets[index]

	if bucket.variables == nil {
		//update attribute locations and re-link program.
		for name, index := range attributelocations {
			gl.BindAttribLocation(key.shader, index, gl.Str(name+"\000"))
		}
		gl.LinkProgram(key.shader)

		bucket.variables = loadVariables(mesh.Shader())
	}

	//Push transform and material values.
	bucket.transforms = append(bucket.transforms, mesh.Transform.Mul(t))
	for i := range bucket.variables {
		bucket.variables[i].push()
	}

	//Push drawcall.
	if !mesh.Indexed() {
		bucket.drawArrays = append(bucket.drawArrays, drawArraysIndirectCommand{
			C.uint(mesh.Count()),
			1,
			C.uint(mesh.Offset()),
			0,
		})
	} else {
		bucket.drawElements = append(bucket.drawElements, drawElementsIndirectCommand{
			C.uint(mesh.Count()),
			1,
			C.uint(mesh.Offset()),
			C.uint(mesh.VOffset()),
			0,
		})
	}

	return nil
}
