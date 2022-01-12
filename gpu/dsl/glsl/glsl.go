//Package provides a GLSL 110 DSL compiler.
package glsl

import (
	"bytes"
	"fmt"

	"qlova.tech/gpu/dsl"
	"qlova.tech/gpu/dsl/dslutil"
)

type Program struct {
	dslutil.State

	vertexShader bytes.Buffer

	// fragment shader is split up
	// so that we can stitch together
	// the lighting shader.
	fragmentHead bytes.Buffer
	fragmentBody bytes.Buffer

	fragmentShader bytes.Buffer
}

func (p *Program) Shaders() ([]byte, []byte, error) {
	return p.vertexShader.Bytes(), append(append(p.fragmentHead.Bytes(), p.fragmentBody.Bytes()...), '}'), nil
}

// Cores returns GLSL cores based on GLSL version 110.
// Other version of GLSL can extend these cores.
func (p *Program) Cores() (vert, frag, shader dsl.Core) {
	p.State.TypeSystem = p

	vert = Shader{
		Buffer:    &p.vertexShader,
		Main:      &p.vertexShader,
		State:     &p.State,
		outPrefix: "frag_",
	}.Core()

	frag = Shader{
		Buffer:    &p.fragmentHead,
		Main:      &p.fragmentBody,
		State:     &p.State,
		argPrefix: "frag_",
	}.Core()

	shader = Shader{
		Buffer: &p.fragmentHead,
		Main:   &p.fragmentBody,
		State:  &p.State,
	}.Core()

	return
}

type Shader struct {
	*bytes.Buffer
	*dslutil.State

	Main *bytes.Buffer

	argPrefix string
	outPrefix string
}

func (s Shader) newIfElseChain() dsl.IfElseChain {
	return dsl.IfElseChain{
		ElseIf: func(condition dsl.Bool, fn func()) dsl.IfElseChain {
			s.WriteString("else if (")
			s.WriteString(string(condition.Value))
			s.WriteString(") {\n")
			fn()
			s.WriteString("}\n")
			return s.newIfElseChain()
		},
		Else: func(fn func()) {
			s.WriteString("else {\n")
			fn()
			s.WriteString("}\n")
		},
	}
}

func (s Shader) Core() dsl.Core {
	var core dsl.Core
	core.Definer = s.NewDefiner(s,
		"%v %v = %v;\n")

	core.Arg.Attributes = dslutil.Attributes(s, s,
		"varying %[2]v "+s.argPrefix+"%[1]v;\n", "%[1]v")
	core.Arg.Flat = dslutil.Attributes(s, s,
		"attribute %[2]v "+s.argPrefix+"%[1]v;\n", "%[1]v")
	core.Out.Attributes = dslutil.Attributes(s, s,
		"varying %[2]v "+s.outPrefix+"%[1]v;\n", "%[1]v")
	core.Out.Flat = dslutil.Attributes(s, s,
		"attribute %[2]v "+s.outPrefix+"%[1]v;\n", "%[1]v")

	core.Uniform.Uniforms = s.NewUniforms(s,
		"uniform %[2]v %[1]v;\n", "%[1]v")
	core.Get.Uniforms = s.NewUniforms(s,
		"uniform %[2]v %[1]v;\n", "%[1]v")

	core.New = s.NewConstructor(dslutil.Constructor{
		Bool:  "%v",
		Int:   "%v",
		Uint:  "%v",
		Float: "%#vf",
		Vec2:  "vec2(%v, %v)",
		Vec3:  "vec3(%v, %v, %v)",
		Vec4:  "vec4(%v, %v, %v, %v)",
		RGB:   "vec3(%v, %v, %v)",
		RGBA:  "vec4(%v, %v, %v, %v)",
	})

	core.Set = s.NewSetter(s, "%v = %v;\n")

	core.Position = s.NewVec4("gl_Position")
	core.Fragment = s.NewRGBA("gl_FragColor")

	core.Main = func(fn func()) {
		s.WriteString("\nvoid main() {\n")
		if s.Buffer == s.Main {
			fn()
			s.WriteString("}\n")
		} else {
			fn() //lighting will go below.
		}
	}

	core.If = func(condition dsl.Bool, fn func()) dsl.IfElseChain {
		s.WriteString("if (")
		s.WriteString(string(condition.Value))
		s.WriteString(") {\n")
		fn()
		s.WriteString("}\n")

		return s.newIfElseChain()
	}

	core.Range = func(start, end dsl.Int, fn func(dsl.Int)) {
		var name = s.GetVariableName()
		fmt.Fprintf(s, "for (int %v = %v; %v < %v; %v++) {\n",
			name, start.Value, name, end.Value, name)
		fn(s.NewInt(name))
		s.WriteString("}\n")
	}

	core.Discard = func() {
		s.WriteString("discard;\n")
	}
	core.Break = func() {
		s.WriteString("break;\n")
	}
	core.Continue = func() {
		s.WriteString("continue;\n")
	}
	core.While = func(condition dsl.Bool, fn func()) {
		s.WriteString("while (")
		s.WriteString(string(condition.Value))
		s.WriteString(") {\n")
		fn()
		s.WriteString("}\n")
	}

	return core
}
