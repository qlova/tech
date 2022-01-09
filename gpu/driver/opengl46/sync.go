package opengl46

import (
	"reflect"
	"unsafe"

	"qlova.tech/gpu"
	"qlova.tech/gpu/driver/opengl/glsl"
	"qlova.tech/vec/vec4"

	"github.com/go-gl/gl/v4.6-core/gl"
)

//Textured material applies a texture.
type shadowShader struct{}

var shadowShaderPointer uint32

func (s *shadowShader) Variables() []interface{} {
	return nil
}

func (s *shadowShader) CompileTo(string, string) (interface{}, error) {
	return glsl.Shader{
		Vertex: `#version 460 core
		
		in vec4 position;
		out float height;
		out vec4 Position;

		uniform mat4 camera;

		layout(binding=0,std430) readonly buffer transformBuffer{
			mat4 transformArray[];
		};
		#define transform transformArray[gl_DrawID]
		
		void main() {
			vec4 p = position;
			p.y = 0;
			p = transform*p;
			height = float(position.y < 0);
			p.x -= max(position.y, 0)/2;
			p.z -= max(position.y, 0)/2;
			Position = p;
			gl_Position = camera*p;
		}
	`,
		Fragment: `#version 460 core

		out vec4 FragColor;
		in float height;
		in vec4 Position;
		
		void main() {
			FragColor = vec4(1-height, Position.x, Position.y, Position.z);
		}
	`,
	}, nil
}

var variables = make(map[string]interface{})

func (s state) updateUniforms() {
	if s.shader == 0 {
		return
	}

	for variable, value := range variables {

		var location = gl.GetUniformLocation(s.shader, gl.Str(variable+"\000"))
		if location < 0 {
			continue
		}

		switch v := value.(type) {
		case bool:
			if v {
				gl.Uniform1i(location, 1)
			} else {
				gl.Uniform1i(location, 0)
			}

		case int:
			gl.Uniform1i(location, int32(v))
		case int32:
			gl.Uniform1i(location, v)
		case int64:
			gl.Uniform1i(location, int32(v))
		case []int32:
			gl.Uniform1iv(location, int32(len(v)), &v[0])
		case [2]int32:
			gl.Uniform2i(location, v[0], v[1])
		case [3]int32:
			gl.Uniform3i(location, v[0], v[1], v[2])
		case [4]int32:
			gl.Uniform4i(location, v[0], v[1], v[2], v[3])

		case uint:
			gl.Uniform1ui(location, uint32(v))
		case uint32:
			gl.Uniform1ui(location, v)
		case uint64:
			gl.Uniform1ui(location, uint32(v))
		case []uint32:
			gl.Uniform1uiv(location, int32(len(v)), &v[0])
		case [2]uint32:
			gl.Uniform2ui(location, v[0], v[1])
		case [3]uint32:
			gl.Uniform3ui(location, v[0], v[1], v[2])
		case [4]uint32:
			gl.Uniform4ui(location, v[0], v[1], v[2], v[3])

		case float32:
			gl.Uniform1f(location, v)
		case float64:
			gl.Uniform1f(location, float32(v))
		case []float32:
			gl.Uniform1fv(location, int32(len(v)), &v[0])
		case [2]float32:
			gl.Uniform2f(location, v[0], v[1])
		case [][2]float32:
			gl.Uniform2fv(location, int32(len(v)), &v[0][0])
		case [3]float32:
			gl.Uniform3f(location, v[0], v[1], v[2])
		case [4]float32:
			gl.Uniform4f(location, v[0], v[1], v[2], v[3])
		case vec4.Type:
			gl.Uniform4f(location, v[0], v[1], v[2], v[3])

		case [16]float32:
			gl.UniformMatrix4fv(location, 1, false, &v[0])

		case gpu.Transform:
			gl.UniformMatrix4fv(location, 1, false, &v[0])

		case gpu.Texture:
			gl.UniformHandleui64ARB(location, v.Value())

		default:
			panic("unsopported uniform translation: " + reflect.TypeOf(value).String())
		}
	}

	checkErrors()
}

func (b *bucket) updateVariables(discard bool) {
	for i := range b.variables {
		variable := &b.variables[i]

		if variable.pointer == 0 {
			gl.GenBuffers(1, &variable.pointer)
		}

		pointer := unsafe.Pointer(&variable.Buffer[0])
		length := len(variable.Buffer)

		gl.BindBufferBase(gl.SHADER_STORAGE_BUFFER, uint32(i)+1, variable.pointer)
		gl.BufferData(gl.SHADER_STORAGE_BUFFER, length*int(unsafe.Sizeof(variable.Buffer[0])), pointer, gl.DYNAMIC_DRAW)

		if discard {
			variable.Buffer = variable.Buffer[:0]
		}
	}

	if b.transform == 0 {
		gl.GenBuffers(1, &b.transform)
	}

	pointer := unsafe.Pointer(&b.transforms[0])
	length := len(b.transforms)

	gl.BindBufferBase(gl.SHADER_STORAGE_BUFFER, 0, b.transform)
	gl.BufferData(gl.SHADER_STORAGE_BUFFER, length*int(unsafe.Sizeof(b.transforms[0])), pointer, gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.SHADER_STORAGE_BUFFER, 0)

	if discard {
		b.transforms = b.transforms[:0]
	}
}

func (s state) render(b *bucket, discard bool) error {
	if len(b.drawArrays) == 0 && len(b.drawElements) == 0 {
		return nil
	}

	s.updateUniforms()
	b.updateVariables(discard)

	checkErrors()

	gl.BindVertexArray(s.vao)

	if b.pointer == 0 {
		gl.GenBuffers(1, &b.pointer)
	}
	gl.BindBuffer(gl.DRAW_INDIRECT_BUFFER, b.pointer)

	if s.options&gpu.Wireframe != 0 {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	} else {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
	}

	if s.options&gpu.FrontFaceCulling != 0 {
		gl.CullFace(gl.FRONT)
	} else {
		gl.CullFace(gl.BACK)
	}

	checkErrors()

	//Upload command buffer.
	if !s.indexed {

		gl.BufferData(gl.DRAW_INDIRECT_BUFFER,
			len(b.drawArrays)*int(unsafe.Sizeof(drawArraysIndirectCommand{})),
			gl.Ptr(b.drawArrays),
			gl.DYNAMIC_DRAW)

		gl.MultiDrawArraysIndirect(gl.TRIANGLES, nil, int32(len(b.drawArrays)), 0)
		if discard {
			b.drawArrays = b.drawArrays[:0]
		}
		checkErrors()
	} else {
		gl.BufferData(gl.DRAW_INDIRECT_BUFFER,
			len(b.drawElements)*int(unsafe.Sizeof(drawElementsIndirectCommand{})),
			gl.Ptr(b.drawElements),
			gl.DYNAMIC_DRAW)

		checkErrors()

		gl.MultiDrawElementsIndirect(gl.TRIANGLES, gl.UNSIGNED_INT, nil, int32(len(b.drawElements)), 0)
		if discard {
			b.drawElements = b.drawElements[:0]
		}
		checkErrors()
	}

	gl.BindVertexArray(0)

	checkErrors()

	return nil
}

func sync() error {
	gl.Enable(gl.DEPTH_TEST)
	gl.Disable(gl.BLEND)

	gl.CullFace(gl.BACK)
	gl.Enable(gl.CULL_FACE)

	//shadow pass.
	var viewport [4]int32
	gl.GetIntegerv(gl.VIEWPORT, &viewport[0])
	gpu.Set("viewport", [2]float32{float32(viewport[2]), float32(viewport[3])})

	gl.Viewport(0, 0, shadowSize, shadowSize)
	gl.BindFramebuffer(gl.FRAMEBUFFER, shadowMapFBO)
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	//for each shader, for each buffer: render command lists.
	for i := range buckets {
		bucket := &buckets[i]
		state := bucket.state

		if state.options&gpu.NoShadows != 0 {
			continue
		}

		if shadowShaderPointer == 0 {
			/*p, err := loadShader(0, new(shadowShader))
			if err != nil {
				return err
			}
			shadowShaderPointer = uint32(p)*/
		}
		state.shader = shadowShaderPointer

		checkErrors()

		gl.UseProgram(state.shader)

		checkErrors()

		if err := state.render(bucket, false); err != nil {
			return err
		}

		checkErrors()
	}

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Viewport(0, 0, viewport[2], viewport[3])

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, shadowMapTex)

	//for each shader, for each buffer: render command lists.
	for i := range buckets {
		bucket := &buckets[i]
		state := bucket.state

		gl.UseProgram(state.shader)

		if err := state.render(bucket, true); err != nil {
			return err
		}
	}

	return nil
}
