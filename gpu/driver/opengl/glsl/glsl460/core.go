package glsl460

import (
	"bytes"
	"fmt"

	"qlova.tech/gpu"
	"qlova.tech/vec/vec2"
	"qlova.tech/vec/vec3"
	"qlova.tech/vec/vec4"

	"qlova.tech/mat/mat3"
	"qlova.tech/mat/mat4"

	"qlova.tech/rgb"
	"qlova.tech/rgb/rgba"

	"qlova.tech/gpu/driver/opengl/glsl"
	"qlova.tech/gpu/shader"
)

//Compile a vertex and fragment shader to GLSL 430.
func Compile(program gpu.Program) (vert []byte, frag []byte, err error) {
	var core *shader.Core

	vert, core, err = compile(program.Vertex, nil)
	if err != nil {
		return nil, nil, err
	}

	frag, _, err = compile(program.Fragment, core)
	if err != nil {
		return nil, nil, err
	}

	return vert, frag, nil
}

func compile(fn func(*shader.Core), vtx *shader.Core) ([]byte, *shader.Core, error) {
	fragment := vtx != nil

	var buf bytes.Buffer

	var core = glsl.Core(&buf, vtx)
	line := func(format string, args ...interface{}) {
		if indent := core.Indentation(); indent > 0 {
			for i := 0; i < indent; i++ {
				buf.WriteString("\t")
			}
		}
		fmt.Fprintf(&buf, format+"\n", args...)
	}

	line("#version 460")
	line("#extension GL_ARB_bindless_texture : require\n")

	if !fragment {
		line("out int gpu_DrawID;")
	} else {
		line("flat in int gpu_DrawID;")
	}

	makeCurrent := func(rtype string) string {
		index := core.Unique("current")
		name := fmt.Sprintf("current_%v", index)

		line("layout(binding=%v,std430) readonly buffer %s_buffer{", index, name)
		line("\t%s %s[];", rtype, name)
		line("};")

		return fmt.Sprintf("%v[gpu_DrawID]", name)
	}

	core.Current.Vec2 = func(*vec2.Type) shader.Vec2 { return glsl.NewVec2(makeCurrent("vec2")) }
	core.Current.Vec3 = func(*vec3.Type) shader.Vec3 { return glsl.NewVec3(makeCurrent("vec3")) }
	core.Current.Vec4 = func(*vec4.Type) shader.Vec4 { return glsl.NewVec4(makeCurrent("vec4")) }
	core.Current.RGB = func(*rgb.Color) shader.RGB { return glsl.NewRGB(makeCurrent("vec3")) }
	core.Current.RGBA = func(*rgba.Color) shader.RGBA { return glsl.NewRGBA(makeCurrent("vec4")) }
	core.Current.Mat3 = func(*mat3.Type) shader.Mat3 { return glsl.NewMat3(makeCurrent("mat3")) }
	core.Current.Mat4 = func(*mat4.Type) shader.Mat4 { return glsl.NewMat4(makeCurrent("mat4")) }
	core.Current.Sampler = func(*gpu.Texture) shader.Sampler { return glsl.NewSampler(makeCurrent("sampler2D")) }

	core.Main = func(fn func()) {
		line("\nvoid main() {")
		if !fragment {
			line("\tgpu_DrawID = gl_DrawID;")
		}
		core.Indent(fn)
		line("}\n")
	}

	fn(core)

	return buf.Bytes(), core, nil
}
