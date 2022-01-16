//go:build js

package webgl

import (
	"fmt"
	"image"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall/js"
	"unsafe"

	"qlova.tech/app"
	"qlova.tech/dsl"
	"qlova.tech/dsl/dslutil"
	"qlova.tech/gpu"
	"qlova.tech/rgb"
	"qlova.tech/xy"
	"qlova.tech/xyz"

	"qlova.tech/dsl/glsl/glsl100"
)

var queue = make(chan func(), 256)

var window = js.Global()
var gl js.Value

func init() {
	gpu.Register("WebGL 1.0", func() (gpu.Driver, error) {
		if err := open(); err != nil {
			return gpu.Driver{}, err
		}

		return gpu.Driver{
			NewFrame:   newFrame,
			NewMesh:    newMesh,
			NewTexture: newTexture,
			NewProgram: newProgram,

			//no implementations yet.
			SetLighting: func(lights []gpu.Light) {},
			SetRaycasts: func(raycasts []gpu.Raycast) {},

			Draw: draw,
			Sync: func() {
				gl.Call("finish")
				for {
					select {
					case fn := <-queue:
						fn()
					default:
						return
					}
				}
			},
		}, nil
	})
}

func open() error {
	//look for a canvas with the id "gpu"

	canvas := window.Get("document").Call("getElementById", "gpu")
	if canvas.IsNull() {
		return fmt.Errorf("webgl: no 'gpu 'canvas found")
	}

	gl = canvas.Call("getContext", "webgl")
	if gl.IsNull() {
		return fmt.Errorf("webgl: a webgl context could not be instantiated")
	}

	gl.Call("enable", gl.Get("DEPTH_TEST"))

	return nil
}

func newFrame(color rgb.Color) {
	gl.Call("viewport", 0, 0, app.Width, app.Height)
	gl.Call("clearColor", float32(color.Red())/255, float32(color.Green())/255, float32(color.Blue())/255, 1)
	gl.Call("clear", gl.Get("COLOR_BUFFER_BIT").Int()|gl.Get("DEPTH_BUFFER_BIT").Int())
}

type vertexAttributePointer struct {
	attribute dsl.Attribute

	size    int32
	kind    uint32
	norm    bool
	stride  int32
	pointer js.Value
}

type vertexAttributeObject struct {
	// mutable flag for reuse. if non-zero
	// no other goroutine has a reference
	// and this vao can be reused.
	deleted uint32

	pointers []vertexAttributePointer

	//pointers to the buffers being used.
	buffers []js.Value

	count int

	indexed  bool
	indicies js.Value
}

var vaos hotswapVAOs

type hotswapVAOs struct {
	mutex   sync.Mutex
	writing uint32
	a, b    []vertexAttributeObject
}

func (vao *hotswapVAOs) Read() []vertexAttributeObject {
	switch atomic.LoadUint32(&vao.writing) {
	case 0:
		return vao.b
	default:
		return vao.a
	}
}

func (vaos *hotswapVAOs) Write(vao vertexAttributeObject) int {

	//fast lock-free method.
	slice := vaos.Read()
	for i := range slice {
		if atomic.CompareAndSwapUint32(&slice[i].deleted, 1, 0) {
			slice[i] = vao
			return i
		}
	}

	//slow append.
	vaos.mutex.Lock()
	defer vaos.mutex.Unlock()

	var index int

	switch atomic.LoadUint32(&vaos.writing) {
	case 0:
		index = len(vaos.b)
		vaos.a = append(vaos.b, vao)
		atomic.StoreUint32(&vaos.writing, 1)

		//copy in deleted vaos.
		for i := range vaos.b {
			if atomic.LoadUint32(&vaos.b[i].deleted) == 1 {
				vaos.a[i].deleted = vaos.b[i].deleted
			}
		}
	default:
		index = len(vaos.a)
		vaos.b = append(vaos.a, vao)
		atomic.StoreUint32(&vaos.writing, 0)

		//copy in deleted vaos.
		for i := range vaos.a {
			if atomic.LoadUint32(&vaos.a[i].deleted) == 1 {
				vaos.b[i].deleted = vaos.a[i].deleted
			}
		}
	}

	return index
}

func newMesh(vertices xyz.Vertices, hints ...gpu.MeshHint) (unsafe.Pointer, error) {
	var buffers = vertices.Buffers()
	var layout = vertices.Layout()
	var indicies, indexed = vertices.Indexed()

	var pointers = make([]js.Value, len(buffers))

	for i, buffer := range buffers {
		pointers[i] = gl.Call("createBuffer")

		var jsBuffer = window.Get("Uint8Array").New(len(buffer))
		js.CopyBytesToJS(jsBuffer, buffer)

		if indexed && i == indicies {
			gl.Call("bindBuffer", gl.Get("ELEMENT_ARRAY_BUFFER"), pointers[i])
			gl.Call("bufferData", gl.Get("ELEMENT_ARRAY_BUFFER"), jsBuffer, gl.Get("STATIC_DRAW"))
		} else {
			gl.Call("bindBuffer", gl.Get("ARRAY_BUFFER"), pointers[i])
			gl.Call("bufferData", gl.Get("ARRAY_BUFFER"), jsBuffer, gl.Get("STATIC_DRAW"))
		}
	}

	var attributes = make([]vertexAttributePointer, len(layout))
	for i, attribute := range layout {
		var kind int
		var size = attribute.Count

		switch attribute.Kind {
		case reflect.Bool, reflect.Int8:
			kind = gl.Get("BYTE").Int()
		case reflect.Uint8:
			kind = gl.Get("UNSIGNED_BYTE").Int()
		case reflect.Int16:
			kind = gl.Get("SHORT").Int()
		case reflect.Uint16:
			kind = gl.Get("UNSIGNED_SHORT").Int()
		case reflect.Int32:
			kind = gl.Get("INT").Int()
		case reflect.Uint32:
			kind = gl.Get("UNSIGNED_INT").Int()
		case reflect.Float32:
			kind = gl.Get("FLOAT").Int()
		case reflect.Complex64:
			kind = gl.Get("FLOAT").Int()
			size /= 2
		default:
			return nil, fmt.Errorf("unsupported vertex attribute kind: %s", attribute.Kind)
		}

		attributes[i] = vertexAttributePointer{
			attribute: dsl.Attribute(attribute.Attribute),
			size:      int32(size),
			kind:      uint32(kind),
			norm:      true,
			stride:    int32(attribute.Stride),
			pointer:   pointers[i],
		}
	}

	vao := vertexAttributeObject{
		pointers: attributes,
		buffers:  pointers,
		count:    vertices.Length(),
		indexed:  indexed,
	}
	if indexed {
		vao.indicies = pointers[indicies]
	}

	index := vaos.Write(vao)

	runtime.SetFinalizer(&index, func(index *int) {
		vaos := vaos.Read()

		var i = *index
		var buffers = vaos[i].buffers

		//this needs to run on the main thread...
		queue <- func() {
			for _, buffer := range buffers {
				gl.Call("deleteBuffer", buffer)
			}
		}

		atomic.StoreUint32(&vaos[i].deleted, 1)
	})

	return unsafe.Pointer(&index), nil
}

func newTexture(image image.Image, hints ...gpu.TextureHint) (unsafe.Pointer, error) {
	return nil, nil
}

func compileError(shader js.Value) error {
	var log = gl.Call("getShaderInfoLog", shader).String()

	if len(log) > 0 {
		return fmt.Errorf("%s", log)
	}

	return nil
}

type shader struct {
	program  js.Value
	uniforms []dslutil.Uniform
}

func newProgram(v, f dsl.Shader, hints ...gpu.ProgramHint) (unsafe.Pointer, error) {
	var s shader

	var program = gl.Call("createProgram")

	var source glsl100.Source

	vert, frag := source.Cores()
	v(vert)
	f(frag)

	vertSrc, fragSrc, err := source.Files()
	if err != nil {
		return nil, err
	}
	vertSrc = append(vertSrc, 0)
	fragSrc = append(fragSrc, 0)

	//fmt.Println(string(vertSrc))
	//fmt.Println(string(fragSrc))

	s.uniforms = source.Uniforms

	var vertex = gl.Call("createShader", gl.Get("VERTEX_SHADER"))
	gl.Call("shaderSource", vertex, string(vertSrc))
	gl.Call("compileShader", vertex)
	if err := compileError(vertex); err != nil {
		return nil, err
	}

	var fragment = gl.Call("createShader", gl.Get("FRAGMENT_SHADER"))
	gl.Call("shaderSource", fragment, string(fragSrc))
	gl.Call("compileShader", fragment)
	if err := compileError(fragment); err != nil {
		return nil, err
	}

	gl.Call("attachShader", program, vertex)
	gl.Call("attachShader", program, fragment)
	gl.Call("deleteShader", vertex)
	gl.Call("deleteShader", fragment)

	gl.Call("linkProgram", program)

	var status = gl.Call("getProgramParameter",
		program, gl.Get("LINK_STATUS")).Bool()
	if !status {
		return nil, fmt.Errorf("program link failed: %s",
			gl.Call("getProgramInfoLog", program))
	}

	runtime.SetFinalizer(&s, func(s *shader) {
		program := s.program
		queue <- func() {
			gl.Call("deleteProgram", program)
		}
	})

	s.program = program

	return unsafe.Pointer(&s), nil
}

func draw(program, mesh unsafe.Pointer) {
	var s = *(*shader)(program)
	var idx = *(*uint32)(mesh)

	gl.Call("useProgram", s.program)
	var vao = vaos.Read()[idx]

	//load uniforms
	var uniforms = s.uniforms
	for _, uniform := range uniforms {
		var location = gl.Call("getUniformLocation", s.program, uniform.Name)

		switch v := uniform.Pointer.(type) {
		case *bool:
			if *v {
				gl.Call("uniform1i", location, 1)
			} else {
				gl.Call("uniform1i", location, 0)
			}
		case int32:
			gl.Call("uniform1i", location, int32(v))
		case uint32:
			gl.Call("uniform1i", location, int32(v))
		case *float32:
			gl.Call("uniform1f", location, *v)

		case *xy.Vector:
			gl.Call("uniform2f", location, v[0], v[1])
		case *xyz.Vector:
			gl.Call("uniform3f", location, v[0], v[1], v[2])
		//case *vec4.Float32:
		//	gl.Uniform4fv(location, 1, &v[0])

		case *rgb.Color:
			a := [4]float32{
				float32(v.Red()) / 255,
				float32(v.Green()) / 255,
				float32(v.Blue()) / 255,
				float32(v.Alpha()) / 255,
			}
			gl.Call("uniform4f", location, a[0], a[1], a[2], a[3])

		//case *mat2.Float32:
		//	gl.UniformMatrix2fv(location, 1, false, &v[0])
		case *xy.Transform:
			gl.Call("uniformMatrix3fv", location, false, v)
		case *xyz.Transform:
			var array = js.Global().Get("Float32Array").New(16)
			for i := 0; i < 16; i++ {
				array.SetIndex(i, v[i])
			}
			gl.Call("uniformMatrix4fv", location, false, array)
		}
	}

	for _, pointer := range vao.pointers {
		if index := gl.Call("getAttribLocation", s.program, string(pointer.attribute)).Int(); index >= 0 {
			gl.Call("enableVertexAttribArray", uint32(index))
			gl.Call("bindBuffer", gl.Get("ARRAY_BUFFER"), pointer.pointer)
			gl.Call("vertexAttribPointer", uint32(index), pointer.size, pointer.kind, pointer.norm, pointer.stride, 0)
		}
	}

	if vao.indexed {
		gl.Call("bindBuffer", gl.Get("ELEMENT_ARRAY_BUFFER"), vao.indicies)
		gl.Call("drawElements", gl.Get("TRIANGLES"), int32(vao.count), gl.Get("UNSIGNED_SHORT"), 0)
	} else {
		gl.Call("drawArrays", gl.Get("TRIANGLES"), 0, int32(vao.count))
	}

	for _, pointer := range vao.pointers {
		if index := gl.Call("getAttribLocation", s.program, string(pointer.attribute)).Int(); index >= 0 {
			gl.Call("disableVertexAttribArray", uint32(index))
		}
	}
}
