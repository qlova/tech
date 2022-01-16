//Package provides an opengl 2.1 gpu driver.
package opengl

import (
	"fmt"
	"image"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"

	"qlova.tech/app"
	"qlova.tech/dsl"
	"qlova.tech/dsl/dslutil"
	"qlova.tech/dsl/glsl/glsl110"
	"qlova.tech/gpu"
	"qlova.tech/rgb"
	"qlova.tech/xy"
	"qlova.tech/xyz"
)

func init() {
	gpu.Register("OpenGL 2.1", func() (gpu.Driver, error) {
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

	if err := gl.Init(); err != nil {
		return err
	}

	gl.Enable(gl.DEPTH_TEST)

	atomic.StoreInt32(&opened, 1)
	return nil
}

func newFrame(color rgb.Color) {
	gl.Viewport(0, 0, int32(app.Width), int32(app.Height))

	gl.ClearColor(float32(color.Red())/255, float32(color.Green())/255, float32(color.Blue())/255, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

type vertexAttributePointer struct {
	attribute dsl.Attribute

	size    int32
	kind    uint32
	norm    bool
	stride  int32
	pointer uint32
}

type vertexAttributeObject struct {
	// mutable flag for reuse. if non-zero
	// no other goroutine has a reference
	// and this vao can be reused.
	deleted uint32

	pointers []vertexAttributePointer

	//pointers to the buffers being used.
	buffers []uint32

	count int

	indexed  bool
	indicies uint32
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

	var pointers = make([]uint32, len(buffers))
	gl.GenBuffers(int32(len(pointers)), &pointers[0])

	for i, buffer := range buffers {
		if indexed && i == indicies {
			gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, pointers[i])
			gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(buffer), gl.Ptr(&buffer[0]), gl.STATIC_DRAW)
		} else {
			gl.BindBuffer(gl.ARRAY_BUFFER, pointers[i])
			gl.BufferData(gl.ARRAY_BUFFER, len(buffer), gl.Ptr(&buffer[0]), gl.STATIC_DRAW)
		}
	}

	var attributes = make([]vertexAttributePointer, len(layout))
	for i, attribute := range layout {
		var kind uint32
		var size = attribute.Count

		switch attribute.Kind {
		case reflect.Bool, reflect.Int8:
			kind = gl.BYTE
		case reflect.Uint8:
			kind = gl.UNSIGNED_BYTE
		case reflect.Int16:
			kind = gl.SHORT
		case reflect.Uint16:
			kind = gl.UNSIGNED_SHORT
		case reflect.Int32:
			kind = gl.INT
		case reflect.Uint32:
			kind = gl.UNSIGNED_INT
		case reflect.Float32:
			kind = gl.FLOAT
		case reflect.Complex64:
			kind = gl.FLOAT
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
			gl.DeleteBuffers(int32(len(buffers)), &buffers[0])
		}

		atomic.StoreUint32(&vaos[i].deleted, 1)
	})

	return unsafe.Pointer(&index), nil
}

func newTexture(tex image.Image, hints ...gpu.TextureHint) (unsafe.Pointer, error) {
	width, height := tex.Bounds().Dx(), tex.Bounds().Dy()

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
	})

	return unsafe.Pointer(&ptr), nil
}

func compileError(shader uint32) error {
	var logLength int32
	gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

	if logLength > 0 {
		var log = make([]byte, logLength)
		gl.GetShaderInfoLog(shader, logLength, nil, &log[0])
		return fmt.Errorf("%s", log)
	}

	return nil
}

var programUniforms = make(map[uint32][]dslutil.Uniform)

func newProgram(v, f dsl.Shader, hints ...gpu.ProgramHint) (unsafe.Pointer, error) {
	var program = gl.CreateProgram()

	var source glsl110.Source

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

	programUniforms[program] = source.Uniforms

	var vertex = gl.CreateShader(gl.VERTEX_SHADER)
	vsources, free := gl.Strs(string(vertSrc))
	gl.ShaderSource(vertex, 1, vsources, nil)
	gl.CompileShader(vertex)
	free()
	if err := compileError(vertex); err != nil {
		return nil, err
	}

	var fragment = gl.CreateShader(gl.FRAGMENT_SHADER)
	fsources, free := gl.Strs(string(fragSrc))
	gl.ShaderSource(fragment, 1, fsources, nil)
	gl.CompileShader(fragment)
	free()
	if err := compileError(fragment); err != nil {
		return nil, err
	}

	gl.AttachShader(program, vertex)
	gl.AttachShader(program, fragment)
	gl.DeleteShader(vertex)
	gl.DeleteShader(fragment)

	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)
		if logLength > 0 {
			var log = make([]byte, logLength)
			gl.GetProgramInfoLog(program, logLength, nil, &log[0])
			return nil, fmt.Errorf("%s", log)
		}
	}

	runtime.SetFinalizer(&program, func(program *uint32) {
		queue <- func() {
			gl.DeleteProgram(*program)
		}
	})

	return unsafe.Pointer(&program), nil
}

func draw(program, mesh unsafe.Pointer) {
	gl.UseProgram(*(*uint32)(program))
	var vao = vaos.Read()[*(*uint32)(mesh)]

	//load uniforms
	var uniforms = programUniforms[*(*uint32)(program)]
	for _, uniform := range uniforms {
		var location = gl.GetUniformLocation(*(*uint32)(program), gl.Str(uniform.Name+"\x00"))

		switch v := uniform.Pointer.(type) {
		case *bool:
			if *v {
				gl.Uniform1i(location, 1)
			} else {
				gl.Uniform1i(location, 0)
			}
		case int32:
			gl.Uniform1i(location, int32(v))
		case uint32:
			gl.Uniform1i(location, int32(v))
		case *float32:
			gl.Uniform1f(location, *v)

		case *xy.Vector:
			gl.Uniform2fv(location, 1, &v[0])
		case *xyz.Vector:
			gl.Uniform3fv(location, 1, &v[0])
		//case *vec4.Float32:
		//	gl.Uniform4fv(location, 1, &v[0])

		case *rgb.Color:
			a := [4]float32{
				float32(v.Red()) / 255,
				float32(v.Green()) / 255,
				float32(v.Blue()) / 255,
				float32(v.Alpha()) / 255,
			}
			gl.Uniform4fv(location, 1, &a[0])

		//case *mat2.Float32:
		//	gl.UniformMatrix2fv(location, 1, false, &v[0])
		case *xy.Transform:
			gl.UniformMatrix3fv(location, 1, false, &v[0])
		case *xyz.Transform:
			gl.UniformMatrix4fv(location, 1, false, &v[0])
		}
	}

	for _, pointer := range vao.pointers {
		if index := gl.GetAttribLocation(*(*uint32)(program), gl.Str(string(pointer.attribute+"\x00"))); index >= 0 {
			gl.EnableVertexAttribArray(uint32(index))
			gl.BindBuffer(gl.ARRAY_BUFFER, pointer.pointer)
			gl.VertexAttribPointerWithOffset(uint32(index), pointer.size, pointer.kind, pointer.norm, pointer.stride, 0)
		}
	}

	if vao.indexed {
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vao.indicies)
		gl.DrawElementsWithOffset(gl.TRIANGLES, int32(vao.count), gl.UNSIGNED_SHORT, 0)
	} else {
		gl.DrawArrays(gl.TRIANGLES, 0, int32(vao.count))
	}

	for _, pointer := range vao.pointers {
		if index := gl.GetAttribLocation(*(*uint32)(program), gl.Str(string(pointer.attribute+"\x00"))); index >= 0 {
			gl.DisableVertexAttribArray(uint32(index))
		}
	}
}
