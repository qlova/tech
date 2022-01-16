//go:build android

//Package provides an OpenGL ES 2.0 gpu driver.
package opengl

import (
	"fmt"
	"image"
	"log"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"

	es "golang.org/x/mobile/gl"

	"qlova.tech/app"
	"qlova.tech/dsl"
	"qlova.tech/dsl/dslutil"
	"qlova.tech/dsl/glsl/glsl100"
	"qlova.tech/gpu"
	"qlova.tech/rgb"
	"qlova.tech/xy"
	"qlova.tech/xyz"
)

var gl es.Context

func init() {
	gpu.Register("OpenGL ES 2.0", func() (gpu.Driver, error) {
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
				gl.Finish()
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

var opened int32

func open() error {
	if atomic.LoadInt32(&opened) == 1 {
		return nil //if already opened, do nothing.
	}

	if app.GL != nil {
		gl = app.GL
	} else {
		return fmt.Errorf("no OpenGL ES 2.0 context available")
	}

	gl.Enable(es.DEPTH_TEST)

	atomic.StoreInt32(&opened, 1)
	return nil
}

func newFrame(color rgb.Color) {
	gl.Viewport(0, 0, app.Width, app.Height)

	gl.ClearColor(float32(color.Red())/255, float32(color.Green())/255, float32(color.Blue())/255, 1)
	gl.Clear(es.COLOR_BUFFER_BIT | es.DEPTH_BUFFER_BIT)
}

type vertexAttributePointer struct {
	attribute dsl.Attribute

	size    int32
	kind    uint32
	norm    bool
	stride  int32
	pointer es.Buffer
}

type vertexAttributeObject struct {
	// mutable flag for reuse. if non-zero
	// no other goroutine has a reference
	// and this vao can be reused.
	deleted uint32

	pointers []vertexAttributePointer

	//pointers to the buffers being used.
	buffers []es.Buffer

	count int

	indexed  bool
	indicies es.Buffer
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

var queue = make(chan func(), 256)

// naive mesh loader, we make no transformation to the
// vertex array and use it as is.
func newMesh(vertices xyz.Vertices, hints ...gpu.MeshHint) (unsafe.Pointer, error) {
	var buffers = vertices.Buffers()
	var layout = vertices.Layout()
	var indicies, indexed = vertices.Indexed()

	var pointers = make([]es.Buffer, len(buffers))

	for i, buffer := range buffers {
		pointers[i] = gl.CreateBuffer()

		if indexed && i == indicies {
			gl.BindBuffer(es.ELEMENT_ARRAY_BUFFER, pointers[i])
			gl.BufferData(es.ELEMENT_ARRAY_BUFFER, buffer, es.STATIC_DRAW)
		} else {
			gl.BindBuffer(es.ARRAY_BUFFER, pointers[i])
			gl.BufferData(es.ARRAY_BUFFER, buffer, es.STATIC_DRAW)
		}
	}

	var attributes = make([]vertexAttributePointer, len(layout))
	for i, attribute := range layout {
		var kind uint32
		var size = attribute.Count

		switch attribute.Kind {
		case reflect.Bool, reflect.Int8:
			kind = es.BYTE
		case reflect.Uint8:
			kind = es.UNSIGNED_BYTE
		case reflect.Int16:
			kind = es.SHORT
		case reflect.Uint16:
			kind = es.UNSIGNED_SHORT
		case reflect.Int32:
			kind = es.INT
		case reflect.Uint32:
			kind = es.UNSIGNED_INT
		case reflect.Float32:
			kind = es.FLOAT
		case reflect.Complex64:
			kind = es.FLOAT
			size /= 2
		default:
			return nil, fmt.Errorf("unsupported vertex attribute kind: %s", attribute.Kind)
		}

		attributes[i] = vertexAttributePointer{
			attribute: dsl.Attribute(attribute.Attribute),
			size:      int32(size),
			kind:      kind,
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
				gl.DeleteBuffer(buffer)
			}
		}

		atomic.StoreUint32(&vaos[i].deleted, 1)
	})

	return unsafe.Pointer(&index), nil
}

func newTexture(tex image.Image, hints ...gpu.TextureHint) (unsafe.Pointer, error) {
	/*width, height := tex.Bounds().Dx(), tex.Bounds().Dy()

	var ptr uint32
	gl.GenTextures(1, &ptr)
	gl.BindTexture(gl.TEXTURE_2D, ptr)

	switch v := tex.(type) {
	case *image.NRGBA:
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(width), int32(height), 0, gl.RGB, gl.UNSIGNED_BYTE, gl.Ptr(v.Pix))
	default:
		return nil, fmt.Errorf("unsupported image type: %T", tex)
	}

	runtime.SetFinalizer(&ptr, func(ptr *uint32) {
		queue <- func() {
			gl.DeleteTextures(1, ptr)
		}
	})*/

	return nil, nil
}

func compileError(shader es.Shader) error {
	var log = gl.GetShaderInfoLog(shader)
	if len(log) > 0 {
		return fmt.Errorf("%s", log)
	}

	return nil
}

type shader struct {
	program  es.Program
	uniforms []dslutil.Uniform
}

func newProgram(v, f dsl.Shader, hints ...gpu.ProgramHint) (unsafe.Pointer, error) {
	var s shader

	var program = gl.CreateProgram()

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

	var vertex = gl.CreateShader(es.VERTEX_SHADER)
	gl.ShaderSource(vertex, string(vertSrc))
	gl.CompileShader(vertex)
	if err := compileError(vertex); err != nil {
		return nil, err
	}

	var fragment = gl.CreateShader(es.FRAGMENT_SHADER)
	gl.ShaderSource(fragment, string(fragSrc))
	gl.CompileShader(fragment)
	if err := compileError(fragment); err != nil {
		return nil, err
	}

	gl.AttachShader(program, vertex)
	gl.AttachShader(program, fragment)
	gl.DeleteShader(vertex)
	gl.DeleteShader(fragment)

	gl.LinkProgram(program)

	var status = gl.GetProgrami(program, es.LINK_STATUS)
	if status == es.FALSE {
		return nil, fmt.Errorf("%s", gl.GetProgramInfoLog(program))
	}

	s.program = program

	runtime.SetFinalizer(&s, func(s *shader) {
		program := s.program
		queue <- func() {
			gl.DeleteProgram(program)
		}
	})

	log.Println(s)

	return unsafe.Pointer(&s), nil
}

func draw(program, mesh unsafe.Pointer) {
	log.Println(program)

	var s = *(*shader)(program)

	gl.UseProgram(s.program)
	var vao = vaos.Read()[*(*uint32)(mesh)]

	//load uniforms
	var uniforms = s.uniforms
	for _, uniform := range uniforms {
		var location = gl.GetUniformLocation(s.program, uniform.Name)

		switch v := uniform.Pointer.(type) {
		case *bool:
			if *v {
				gl.Uniform1i(location, 1)
			} else {
				gl.Uniform1i(location, 0)
			}
		case int32:
			gl.Uniform1i(location, int(v))
		case uint32:
			gl.Uniform1i(location, int(v))
		case *float32:
			gl.Uniform1f(location, *v)

		case *xy.Vector:
			gl.Uniform2fv(location, v[:])
		case *xyz.Vector:
			gl.Uniform3fv(location, v[:])
		//case *vec4.Float32:
		//	gl.Uniform4fv(location, 1, &v[0])

		case *rgb.Color:
			a := [4]float32{
				float32(v.Red()) / 255,
				float32(v.Green()) / 255,
				float32(v.Blue()) / 255,
				float32(v.Alpha()) / 255,
			}
			gl.Uniform4fv(location, a[:])

		//case *mat2.Float32:
		//	gl.UniformMatrix2fv(location, 1, false, &v[0])
		case *xy.Transform:
			gl.UniformMatrix3fv(location, v[:])
		case *xyz.Transform:
			gl.UniformMatrix4fv(location, v[:])
		}
	}

	for _, pointer := range vao.pointers {
		index := gl.GetAttribLocation(s.program, string(pointer.attribute))
		gl.EnableVertexAttribArray(index)
		gl.BindBuffer(es.ARRAY_BUFFER, pointer.pointer)
		gl.VertexAttribPointer(index, int(pointer.size), es.Enum(pointer.kind), pointer.norm, int(pointer.stride), 0)
	}

	if vao.indexed {
		gl.BindBuffer(es.ELEMENT_ARRAY_BUFFER, vao.indicies)
		gl.DrawElements(es.TRIANGLES, int(vao.count), es.UNSIGNED_SHORT, 0)
	} else {
		gl.DrawArrays(es.TRIANGLES, 0, int(vao.count))
	}

	for _, pointer := range vao.pointers {
		index := gl.GetAttribLocation(s.program, string(pointer.attribute))
		gl.DisableVertexAttribArray(index)
	}
}
