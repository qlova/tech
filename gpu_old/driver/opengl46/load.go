package opengl46

import (
	"encoding/binary"
	"errors"
	"fmt"
	"image"
	imgdraw "image/draw"
	"reflect"
	"strings"

	"qlova.tech/gpu"
	"qlova.tech/gpu/driver/opengl/glsl"

	"github.com/go-gl/gl/v4.6-core/gl"
)

var vaobuffers = make(map[uint32][]uint32)

var attributelocations = make(map[string]uint32)

func getAttribLocation(name string) uint32 {
	loc, ok := attributelocations[name]
	if !ok {
		loc = uint32(len(attributelocations))
		attributelocations[name] = loc
	}
	return loc
}

func load(data interface{}) (buf uint64, err error) {
	if update, ok := data.(gpu.Update); ok {
		buf = update.Pointer.Value()
		data = update.Data
	}

	switch v := data.(type) {

	case gpu.Variable:
		variables[v.Name] = v.Value
		return 0, nil

	//Load the next frame.
	case gpu.Frame:
		gl.Finish()
		gl.ClearColor(v.ClearColor[0], v.ClearColor[1], v.ClearColor[2], v.ClearColor[3])
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
		return 0, nil

	case image.Image:
		return loadTexture(buf, v)

	//Load mesh data.
	case gpu.Attributes:
		return loadAttributes(buf, v)

	//Compile shader.
	case gpu.Shader:
		return loadShader(buf, v)

	default:
		return 0, fmt.Errorf("invalid load of type: ", reflect.TypeOf(data))
	}
}

func loadTexture(p uint64, img image.Image) (uint64, error) {
	var id uint32
	gl.GenTextures(1, &id)
	gl.BindTexture(gl.TEXTURE_2D, id)

	bounds := img.Bounds()
	width, height := int32(bounds.Max.X), int32(bounds.Max.Y)

	var format int32
	var data interface{}

	switch v := img.(type) {
	case *image.NRGBA:
		format = gl.RGBA
		data = v.Pix
	case *image.RGBA:
		format = gl.RGBA
		data = v.Pix
	default:

		m := image.NewNRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
		imgdraw.Draw(m, m.Bounds(), img, bounds.Min, imgdraw.Src)

		format = gl.RGBA
		data = m.Pix
	}

	gl.TexImage2D(gl.TEXTURE_2D, 0, format, width, height, 0, uint32(format), gl.UNSIGNED_BYTE, gl.Ptr(data))
	//if options&gpu.NoMipmaps == 0 {
	gl.GenerateMipmap(gl.TEXTURE_2D)
	//}
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	//if options&gpu.NoMipmaps == 0 {
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	//} else {
	//	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	//}

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	pointer := gl.GetTextureHandleARB(id)
	gl.MakeTextureHandleResidentARB(pointer)

	return pointer, nil
}

func loadShader(p uint64, shader gpu.Shader) (uint64, error) {
	var pointer = uint32(p)
	if pointer == 0 {
		pointer = gl.CreateProgram()
	}

	function, err := shader.CompileTo("glsl", "460")
	if err != nil {
		return 0, err
	}

	compileSource := func(kind uint32, source string) (uint32, error) {
		shader := gl.CreateShader(kind)

		csources, free := gl.Strs(string(source) + "\000")
		gl.ShaderSource(shader, 1, csources, nil)
		free()
		gl.CompileShader(shader)

		var status int32
		gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
		if status == gl.FALSE {
			var logLength int32
			gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

			log := strings.Repeat("\x00", int(logLength+1))
			gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

			gl.DeleteShader(shader)

			return 0, fmt.Errorf("failed to compile shader: %v\n %v", log, source)
		}

		return shader, nil
	}

	vertex, err := compileSource(gl.VERTEX_SHADER, function.(glsl.Shader).Vertex)
	if err != nil {
		return 0, err
	}

	fragment, err := compileSource(gl.FRAGMENT_SHADER, function.(glsl.Shader).Fragment)
	if err != nil {
		return 0, err
	}

	gl.AttachShader(pointer, vertex)
	gl.AttachShader(pointer, fragment)

	gl.DeleteShader(vertex)
	gl.DeleteShader(fragment)

	gl.LinkProgram(pointer)

	var status int32
	gl.GetProgramiv(pointer, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(pointer, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(pointer, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	checkErrors()

	return uint64(pointer), nil
}

//loadAttributes uploads attributes to the GPU and returns a VAO pointer as the 'buffer'.
func loadAttributes(p uint64, a gpu.Attributes) (uint64, error) {
	var vao = uint32(p)
	if vao == 0 {
		gl.GenVertexArrays(1, &vao)
	}

	if len(a) == 0 {
		return uint64(vao), nil
	}

	//Get/Generate the buffers where the attributes are stored.
	buffers := vaobuffers[vao]
	if len(buffers) == 0 {
		buffers = make([]uint32, len(a))
		gl.GenBuffers(int32(len(a)), &buffers[0])
	}

	checkErrors()

	gl.BindVertexArray(vao)

	var LastLength int = -1

	for i, attr := range a {
		size := binary.Size(attr)
		length := reflect.ValueOf(attr).Len()

		//Upload indicies.
		_, index := attr.(gpu.Indicies)
		if index {
			gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buffers[i])
			gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, size, gl.Ptr(attr), gl.STATIC_DRAW)
			continue
		}

		if LastLength == -1 {
			LastLength = length
		}

		if length != LastLength {
			return 0, errors.New("attribute length mismatch")
		}

		if size == -1 {
			return 0, errors.New("invalid attribute type")
		}

		loc := getAttribLocation(gpu.AttributeName(attr))

		gl.BindBuffer(gl.ARRAY_BUFFER, buffers[i])
		gl.BufferData(gl.ARRAY_BUFFER, size, gl.Ptr(attr), gl.STATIC_DRAW)
		gl.EnableVertexAttribArray(loc)

		switch T := reflect.TypeOf(attr).Elem(); T {
		case reflect.TypeOf([3]float32{}):
			gl.VertexAttribPointer(uint32(loc), 3, gl.FLOAT, false, 0, nil)
		case reflect.TypeOf([2]float32{}):
			gl.VertexAttribPointer(uint32(loc), 2, gl.FLOAT, false, 0, nil)
		default:
			return 0, fmt.Errorf("invalid attribute type: %v", T)
		}
	}

	vaobuffers[vao] = buffers
	gl.BindVertexArray(0)

	checkErrors()

	return uint64(vao), nil
}
