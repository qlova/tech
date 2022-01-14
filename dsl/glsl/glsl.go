//Package provides a core GLSL compiler.
package glsl

import (
	"bytes"
	"fmt"

	"qlova.tech/dsl"
	"qlova.tech/dsl/dslutil"
)

type Source struct {
	dslutil.State

	vertexHead bytes.Buffer
	vertexBody bytes.Buffer

	fragmentHead bytes.Buffer
	fragmentBody bytes.Buffer
}

func (p *Source) Files() (vert []byte, frag []byte, err error) {
	vert = append(vert, p.vertexHead.Bytes()...)
	vert = append(vert, "\nvoid main() {\n"...)
	vert = append(vert, p.vertexBody.Bytes()...)
	vert = append(vert, "}\n"...)

	frag = append(frag, p.fragmentHead.Bytes()...)
	frag = append(frag, "\nvoid main() {\n"...)
	frag = append(frag, p.fragmentBody.Bytes()...)
	frag = append(frag, "}\n"...)

	return
}

type stage struct {
	*dslutil.State

	Head *bytes.Buffer
	Main *bytes.Buffer

	argPrefix string
	outPrefix string
}

// Cores returns GLSL cores based on GLSL version 110.
// Other version of GLSL can extend these cores.
func (s *Source) Cores() (vert, frag dsl.Core) {

	vert = core(stage{
		Head:      &s.vertexHead,
		Main:      &s.vertexBody,
		State:     &s.State,
		outPrefix: "frag_",
	})

	frag = core(stage{
		Head:      &s.fragmentHead,
		Main:      &s.fragmentBody,
		State:     &s.State,
		argPrefix: "frag_",
	})

	return
}

func (s stage) newIfElseChain() dsl.IfElseChain {
	return dsl.IfElseChain{
		ElseIf: func(condition dsl.Bool, fn func()) dsl.IfElseChain {
			s.Main.WriteString("else if (")
			s.Main.WriteString(string(condition.Value))
			s.Main.WriteString(") {\n")
			fn()
			s.Main.WriteString("}\n")
			return s.newIfElseChain()
		},
		Else: func(fn func()) {
			s.Main.WriteString("else {\n")
			fn()
			s.Main.WriteString("}\n")
		},
	}
}

func core(s stage) dsl.Core {
	var core dsl.Core

	core.Var = s.NewDefiner(s.Main, s,
		"%v %v = %v;\n")

	core.In = dslutil.Attributes(s.Head, s,
		"attribute %[2]v "+s.argPrefix+"%[1]v;\n", "%[1]v")
	core.Out = dslutil.Attributes(s.Head, s,
		"attribute %[2]v "+s.outPrefix+"%[1]v;\n", "%[1]v")

	core.Uniform = s.NewUniforms(s.Head, s,
		"uniform %[2]v %[1]v;\n", "%[1]v")
	core.Get = s.NewUniforms(s.Head, s,
		"uniform %[2]v %[1]v;\n", "%[1]v")

	core.Constructor = s.NewConstructor(s, dslutil.Constructor{
		Bool:  "%v",
		Int:   "%v",
		Uint:  "%v",
		Float: "%#v",
		Vec2:  "vec2(%v, %v)",
		Vec3:  "vec3(%v, %v, %v)",
		Vec4:  "vec4(%v, %v, %v, %v)",
		RGB:   "vec3(%v, %v, %v)",
		RGBA:  "vec4(%v, %v, %v, %v)",
	})

	core.Set = s.NewSetter(s.Main, "%v = %v;\n")

	core.Position = s.NewVec4("gl_Position")
	core.Fragment = s.NewRGBA("gl_FragColor")

	core.If = func(condition dsl.Bool, fn func()) dsl.IfElseChain {
		s.Main.WriteString("if (")
		s.Main.WriteString(string(condition.Value))
		s.Main.WriteString(") {\n")
		fn()
		s.Main.WriteString("}\n")

		return s.newIfElseChain()
	}

	core.Range = func(start, end dsl.Int, fn func(dsl.Int)) {
		var name = s.GetVariableName()
		fmt.Fprintf(s.Main, "for (int %v = %v; %v < %v; %v++) {\n",
			name, start.Value, name, end.Value, name)
		fn(s.NewInt(name))
		s.Main.WriteString("}\n")
	}

	core.Discard = func() {
		s.Main.WriteString("discard;\n")
	}
	core.Break = func() {
		s.Main.WriteString("break;\n")
	}
	core.Continue = func() {
		s.Main.WriteString("continue;\n")
	}
	core.While = func(condition dsl.Bool, fn func()) {
		s.Main.WriteString("while (")
		s.Main.WriteString(string(condition.Value))
		s.Main.WriteString(") {\n")
		fn()
		s.Main.WriteString("}\n")
	}

	return core
}
