package glsl

import (
	"bytes"
	"fmt"

	"qlova.tech/gpu"
)

type State struct {
	Uniforms  []Uniform
	Constants int
	Variables int

	ForLoopNesting int
}

type File struct {
	*State
	Head, Body bytes.Buffer
}

type IfElseChain struct {
	*File
}

func (chain IfElseChain) ElseIf(c gpu.Bool, fn func()) gpu.IfElseChain {
	fmt.Fprintf(&chain.Body, "else if (%s) {\n", c)
	fn()
	chain.Body.WriteString("}\n")
	return chain
}

func (chain IfElseChain) Else(fn func()) {
	fmt.Fprintf(&chain.Body, "else {\n")
	fn()
	chain.Body.WriteString("}\n")
}

type Source struct {
	State

	Vertex, Fragment File
}

func (s *Source) Cores() (gpu.Vertex, gpu.Material, gpu.Fragment) {
	var v, f = s.Vertex.Core(), s.Fragment.Core()

	var m = gpu.Material{
		Core: f,
	}

	return gpu.Vertex{
			Core:     v,
			Position: NewValue[Vec3]("gl_Position"),
		}, m, gpu.Fragment{
			Material: m,
			Color:    NewValue[RGBA]("gl_FragColor"),
		}
}

func (f *File) Core() gpu.Core {
	var shader = Shader{f}
	return gpu.Core{
		Bool: shader.Bool,
		Size: shader.Size,
		Vec2: shader.Vec2,
		Vec3: shader.Vec3,
		RGBA: shader.RGBA,

		Uniform: Uniforms{f},
		Get:     Uniforms{f},
		Const:   Constants{Uniforms{f}},

		Var: Variables{f},

		If: func(condition gpu.Bool, run func()) gpu.IfElseChain {
			fmt.Fprintf(&f.Head, "if (%s) {\n", condition)
			run()
			f.Body.WriteString("}\n")
			return IfElseChain{f}
		},

		For: func(min, max int, do func(int)) {
			f.ForLoopNesting++
			var i = fmt.Sprintf("i_%d", f.ForLoopNesting)
			fmt.Fprintf(&f.Body, "for (int %s = %d; %s < %d; %s++) {\n", i, min, i, max, i)
			do(-f.ForLoopNesting)
			f.Body.WriteString("}\n")
		},
	}
}

type Shader struct {
	*File
}

func (s Shader) Bool(name string) gpu.Bool {
	fmt.Fprintf(&s.Head, "attribute bool %s;\n", name)
	return NewValue[Bool](name)
}

func (s Shader) Size(name string) gpu.Size {
	fmt.Fprintf(&s.Head, "attribute float %s;\n", name)
	return NewValue[Size](name)
}

func (s Shader) Vec2(name string) gpu.Vec2 {
	fmt.Fprintf(&s.Head, "attribute vec2 %s;\n", name)
	return NewValue[Vec2](name)
}

func (s Shader) Vec3(name string) gpu.Vec3 {
	fmt.Fprintf(&s.Head, "attribute vec3 %s;\n", name)
	return NewValue[Vec3](name)
}

func (s Shader) RGBA(name string) gpu.RGBA {
	fmt.Fprintf(&s.Head, "attribute vec4 %s;\n", name)
	return NewValue[RGBA](name)
}
