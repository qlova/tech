package glsl

import (
	"bytes"
	"fmt"

	"qlova.tech/vec/vec2"
	"qlova.tech/vec/vec3"
	"qlova.tech/vec/vec4"

	"qlova.tech/mat/mat3"
	"qlova.tech/mat/mat4"

	"qlova.tech/rgb"
	"qlova.tech/rgb/rgba"

	"qlova.tech/gpu"
	"qlova.tech/gpu/shader"
	"qlova.tech/gpu/vertex"
)

func NewBool(value string) shader.Bool {
	return shader.Bool{
		Value: value,
	}
}

func NewFloat(value string) shader.Float {
	return shader.Float{
		Value: value,

		LessThan: func(other shader.Float) shader.Bool {
			return NewBool(fmt.Sprintf("%s < %s", value, other.Value))
		},
	}
}

func NewMat4(value string) shader.Mat4 {
	return shader.Mat4{
		Value: value,

		Transform: func(other gpu.Vec4) gpu.Vec4 {
			return NewVec4(fmt.Sprintf("%s*%s", value, other.Value))
		},
		Times: func(other gpu.Mat4) gpu.Mat4 {
			return NewMat4(fmt.Sprintf("%s*%s", value, other.Value))
		},
	}
}

func NewMat3(value string) shader.Mat3 {
	return shader.Mat3{
		Value: value,
	}
}

func NewVec4(value string) shader.Vec4 {
	return shader.Vec4{
		Value: value,
	}
}

func NewVec3(value string) shader.Vec3 {
	return shader.Vec3{
		Value: value,
	}
}

func NewVec2(value string) shader.Vec2 {
	return shader.Vec2{
		Value: value,
	}
}

func NewRGB(value string) shader.RGB {
	return shader.RGB{
		Value: value,
	}
}

func NewRGBA(value string) shader.RGBA {
	return shader.RGBA{
		Value: value,

		R: NewFloat(fmt.Sprintf("%s.r", value)),
		G: NewFloat(fmt.Sprintf("%s.g", value)),
		B: NewFloat(fmt.Sprintf("%s.b", value)),
		A: NewFloat(fmt.Sprintf("%s.a", value)),
	}
}

func NewSampler(name string) shader.Sampler {
	return shader.Sampler{
		Value: name,
	}
}

func Core(buf *bytes.Buffer, vtx *shader.Core) *shader.Core {
	fragment := vtx != nil

	var core = vtx
	if core == nil {
		core = new(shader.Core)
	}

	line := func(format string, args ...interface{}) {
		if indent := core.Indentation(); indent > 0 {
			for i := 0; i < indent; i++ {
				buf.WriteString("\t")
			}
		}
		fmt.Fprintf(buf, format, args...)
	}

	core.Position = NewVec4("gl_Position")

	makeQualFor := func(attr vertex.Attribute, kind string, rtype string) string {
		var prefix = kind + "_"
		if kind == "in" {
			if fragment {
				prefix = "frag_"
			} else {
				prefix = ""
			}
		}
		if kind == "out" {
			if !fragment {
				prefix = "frag_"
			}
		}
		var name string
		if attr == "" {
			name = fmt.Sprintf(prefix+"%v", core.Unique(kind))
		} else {
			name = fmt.Sprintf(prefix+"%v", attr)
		}

		line("%s %s %s;\n", kind, rtype, name)

		return name
	}
	makeDeclFor := func(rtype string, value string) string {
		name := fmt.Sprintf("var_%v", core.Unique("var"))

		line("%s %s = %s;\n", rtype, name, value)

		return name
	}

	core.In.Vec2 = func(attr vertex.Attribute) shader.Vec2 { return NewVec2(makeQualFor(attr, "in", "vec2")) }
	core.In.Vec3 = func(attr vertex.Attribute) shader.Vec3 { return NewVec3(makeQualFor(attr, "in", "vec3")) }
	core.In.Vec4 = func(attr vertex.Attribute) shader.Vec4 { return NewVec4(makeQualFor(attr, "in", "vec4")) }
	core.In.RGB = func(attr vertex.Attribute) shader.RGB { return NewRGB(makeQualFor(attr, "in", "vec3")) }
	core.In.RGBA = func(attr vertex.Attribute) shader.RGBA { return NewRGBA(makeQualFor(attr, "in", "vec4")) }

	core.Out.Vec2 = func(attr vertex.Attribute) shader.Vec2 { return NewVec2(makeQualFor(attr, "out", "vec2")) }
	core.Out.Vec3 = func(attr vertex.Attribute) shader.Vec3 { return NewVec3(makeQualFor(attr, "out", "vec3")) }
	core.Out.Vec4 = func(attr vertex.Attribute) shader.Vec4 { return NewVec4(makeQualFor(attr, "out", "vec4")) }
	core.Out.RGB = func(attr vertex.Attribute) shader.RGB { return NewRGB(makeQualFor(attr, "out", "vec3")) }
	core.Out.RGBA = func(attr vertex.Attribute) shader.RGBA { return NewRGBA(makeQualFor(attr, "out", "vec4")) }

	core.Uniform.Vec2 = func(*vec2.Type) shader.Vec2 { return NewVec2(makeQualFor("", "uniform", "vec2")) }
	core.Uniform.Vec3 = func(*vec3.Type) shader.Vec3 { return NewVec3(makeQualFor("", "uniform", "vec3")) }
	core.Uniform.Vec4 = func(*vec4.Type) shader.Vec4 { return NewVec4(makeQualFor("", "uniform", "vec4")) }
	core.Uniform.RGB = func(*rgb.Color) shader.RGB { return NewRGB(makeQualFor("", "uniform", "vec3")) }
	core.Uniform.RGBA = func(*rgba.Color) shader.RGBA { return NewRGBA(makeQualFor("", "uniform", "vec4")) }
	core.Uniform.Sampler = func(*gpu.Texture) shader.Sampler { return NewSampler(makeQualFor("", "uniform", "sampler2D")) }
	core.Uniform.Mat4 = func(*mat4.Type) shader.Mat4 { return NewMat4(makeQualFor("", "uniform", "mat4")) }
	core.Uniform.Mat3 = func(*mat3.Type) shader.Mat3 { return NewMat3(makeQualFor("", "uniform", "mat3")) }

	core.Set.RGBA = func(a, b shader.RGBA) { line("%s = %s;\n", a.Value, b.Value) }
	core.Set.Vec2 = func(a, b shader.Vec2) { line("%s = %s;\n", a.Value, b.Value) }

	core.Definer.RGBA = func(v gpu.RGBA) shader.RGBA { return NewRGBA(makeDeclFor("vec4", v.Value)) }

	core.New.Float = func(v float64) shader.Float { return NewFloat(fmt.Sprintf("%vf", v)) }

	core.Main = func(fn func()) {
		line("\nvoid main() {\n")
		core.Indent(fn)
		line("}\n")
	}

	core.If = func(condition gpu.Bool, fn func()) shader.IfElseChain {
		line("if (%s) {\n", condition.Value)
		core.Indent(fn)
		line("}\n")
		return shader.IfElseChain{}
	}

	core.Sample = func(sampler shader.Sampler, coord gpu.Vec2) gpu.RGBA {
		return NewRGBA(fmt.Sprintf("texture(%v, %v)", sampler.Value, coord.Value))
	}

	core.Discard = func() {
		line("discard;\n")
	}

	core.Set.Vec4 = func(a, b gpu.Vec4) {
		line("%s = %s;\n", a.Value, b.Value)
	}

	return core
}
