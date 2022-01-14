//Package provides an opengl 2.1 gpu driver.
package opengl

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/go-gl/gl/v2.1/gl"

	"qlova.tech/dsl"
	"qlova.tech/dsl/glsl/glsl110"
	"qlova.tech/gpu"
	"qlova.tech/rgb/led"
	"qlova.tech/rgb"
	"qlova.tech/rgb/tex"
	"qlova.tech/xyz/vtx"
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

			//no implementation yet.
			SetLighting: func(lights ...led.Light) {},

			Draw: draw,
			Sync: gl.Finish,
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

	atomic.StoreInt32(&opened, 1)
	return nil
}

func newFrame(color rgb.Color) {
	//gl.Viewport(0, 0, int32(win.Width), int32(win.Height))
	gl.ClearColor(float32(color.Red())/255, float32(color.Green())/255, float32(color.Blue())/255, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.Enable(gl.DEPTH)
}

type vertexAttributePointer struct {
	attribute vtx.Attribute

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
	default:
		index = len(vaos.a)
		vaos.b = append(vaos.a, vao)
		atomic.StoreUint32(&vaos.writing, 0)
	}

	return index
}

var queue = make(chan func(), 256)

// naive mesh loader, we make no transformation to the
// vertex array and use it as is.
func newMesh(vertices vtx.Array, hints ...vtx.Hint) (reader vtx.Reader, ptr gpu.Pointer, err error) {
	var buffers = vertices.Buffers()
	var layout = vertices.Layout()

	var pointers = make([]uint32, len(buffers))
	gl.GenBuffers(int32(len(pointers)), &pointers[0])

	for i, buffer := range buffers {
		gl.BindBuffer(gl.ARRAY_BUFFER, pointers[i])
		gl.BufferData(gl.ARRAY_BUFFER, len(buffer), gl.Ptr(&buffer[0]), gl.STATIC_DRAW)
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
			return nil, gpu.Pointer{}, fmt.Errorf("unsupported vertex attribute kind: %s", attribute.Kind)
		}

		attributes[i] = vertexAttributePointer{
			attribute: attribute.Attribute,
			size:      int32(size),
			kind:      kind,
			norm:      true,
			stride:    int32(attribute.Stride),
			pointer:   pointers[i],
		}
	}

	index := vaos.Write(vertexAttributeObject{
		pointers: attributes,
		buffers:  pointers,
	})

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

	return &index, gpu.Pointer{uint64(index), uint64(vertices.Length())}, nil
}

var supportedTextureFormats = []tex.Format{
	tex.RGB + 8,
	tex.RGB*tex.Alpha + 8,
}

func newTexture(texture tex.Data, hints ...tex.Hint) (tex.Reader, gpu.Pointer, error) {
	width, height := texture.TextureSize()
	format, data, err := texture.TextureData(supportedTextureFormats...)
	if err != nil {
		return nil, gpu.Pointer{}, err
	}

	var ptr uint32
	gl.GenTextures(1, &ptr)
	gl.BindTexture(gl.TEXTURE_2D, ptr)

	switch format {
	case tex.RGB + 8:
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, int32(width), int32(height), 0, gl.RGB, gl.UNSIGNED_BYTE, gl.Ptr(data))
	case tex.RGB*tex.Alpha + 8:
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(width), int32(height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(data))
	default:
		return nil, gpu.Pointer{}, fmt.Errorf("unsupported texture format: %v", format)
	}

	runtime.SetFinalizer(&ptr, func(ptr *uint32) {
		queue <- func() {
			gl.DeleteTextures(1, ptr)
		}
	})

	return &ptr, gpu.Pointer{uint64(ptr)}, nil
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

func newProgram(f, v dsl.Shader, hints ...dsl.Hint) (dsl.Reader, gpu.Pointer, error) {
	var program = gl.CreateProgram()

	vertSrc, fragSrc, err := glsl110.Compile(f, v)
	if err != nil {
		return nil, gpu.Pointer{}, err
	}

	var vertex = gl.CreateShader(gl.VERTEX_SHADER)
	vsources, free := gl.Strs(string(vertSrc))
	gl.ShaderSource(vertex, 1, vsources, nil)
	free()
	gl.CompileShader(vertex)
	if err := compileError(vertex); err != nil {
		return nil, gpu.Pointer{}, err
	}

	var fragment = gl.CreateShader(gl.FRAGMENT_SHADER)
	fsources, free := gl.Strs(string(fragSrc))
	gl.ShaderSource(fragment, 1, fsources, nil)
	free()
	gl.CompileShader(fragment)
	if err := compileError(fragment); err != nil {
		return nil, gpu.Pointer{}, err
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
			return nil, gpu.Pointer{}, fmt.Errorf("%s", log)
		}
	}

	runtime.SetFinalizer(&program, func(program *uint32) {
		queue <- func() {
			gl.DeleteProgram(*program)
		}
	})

	return &program, gpu.Pointer{uint64(program)}, nil
}

func draw(program, mesh gpu.Pointer) {
	gl.UseProgram(uint32(program[0]))

	var vao = vaos.Read()[uint32(mesh[0])]

	for _, pointer := range vao.pointers {
		if index := gl.GetAttribLocation(uint32(program[0]), gl.Str(string(pointer.attribute+"\x00"))); index >= 0 {
			gl.EnableVertexAttribArray(uint32(index))
			gl.BindBuffer(gl.ARRAY_BUFFER, pointer.pointer)
			gl.VertexAttribPointerWithOffset(uint32(index), pointer.size, pointer.kind, pointer.norm, pointer.stride, 0)
		}
	}

	gl.DrawArrays(gl.TRIANGLES, 0, int32(mesh[1]))

	for _, pointer := range vao.pointers {
		if index := gl.GetAttribLocation(uint32(program[0]), gl.Str(string(pointer.attribute+"\x00"))); index >= 0 {
			gl.DisableVertexAttribArray(uint32(index))
		}
	}
}
