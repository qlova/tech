// Package dsl provides a domain-specific shading language.
package dsl

import (
	"qlova.tech/gpu/internal/core"
	"qlova.tech/gpu/vertex"
	"qlova.tech/mat/mat3"
	"qlova.tech/mat/mat4"
	"qlova.tech/rgb"
	"qlova.tech/rgb/rgba"
	"qlova.tech/vec/vec2"
	"qlova.tech/vec/vec3"
	"qlova.tech/vec/vec4"
)

type Attributes struct {
	Bool  func(vertex.Attribute) Bool
	Int   func(vertex.Attribute) Int
	Uint  func(vertex.Attribute) Uint
	Float func(vertex.Attribute) Float

	RGBA func(vertex.Attribute) RGBA
	RGB  func(vertex.Attribute) RGB

	Vec2 func(vertex.Attribute) Vec2
	Vec3 func(vertex.Attribute) Vec3
	Vec4 func(vertex.Attribute) Vec4
}

type Uniforms struct {
	Bool  func(*bool) Bool
	Int   func(*int) Int
	Uint  func(*uint) Uint
	Float func(*float64) Float

	RGBA func(*rgba.Color) RGBA
	RGB  func(*rgb.Color) RGB

	Vec2 func(*vec2.Type) Vec2
	Vec3 func(*vec3.Type) Vec3
	Vec4 func(*vec4.Type) Vec4

	Mat3 func(*mat3.Type) Mat3
	Mat4 func(*mat4.Type) Mat4

	Sampler func(*core.Texture) Sampler
}

type Setter struct {
	Bool  func(a, b Bool)
	Int   func(a, b Int)
	Uint  func(a, b Uint)
	Float func(a, b Float)

	RGBA func(a, b RGBA)
	RGB  func(a, b RGB)

	Vec2 func(a, b Vec2)
	Vec3 func(a, b Vec3)
	Vec4 func(a, b Vec4)

	Mat3 func(a, b Mat3)
	Mat4 func(a, b Mat4)
}

type Constructor struct {
	Float func(f float64) Float
	Int   func(i int) Int
	Uint  func(u uint) Uint
	Bool  func(b bool) Bool

	Vec2 func(x, y Float) Vec2
	Vec3 func(x, y, z Float) Vec3
	Vec4 func(x, y, z, w Float) Vec4

	RGB  func(r, g, b Float) RGB
	RGBA func(r, g, b, a Float) RGBA
}

type Definer struct {
	Make func(size int) Definer

	Bool  func(Bool) Bool
	Int   func(Int) Int
	Uint  func(Uint) Uint
	Float func(Float) Float

	Vec2 func(Vec2) Vec2
	Vec3 func(Vec3) Vec3
	Vec4 func(Vec4) Vec4

	RGB  func(RGB) RGB
	RGBA func(RGBA) RGBA
}

// Core is the core handle for DSL
type Core struct {
	Definer

	Arg, Out struct {
		Flat Attributes
		Attributes
	}

	Uniform, Get struct {
		Array func(size int) Uniforms
		Uniforms
	}

	New Constructor

	Set Setter

	//builtin variables
	Position Vec4
	Fragment RGBA

	Main     func(fn func())
	If       func(c Bool, fn func()) IfElseChain
	Range    func(min, max Int, fn func())
	Discard  func()
	Return   func(v interface{})
	Break    func()
	Continue func()
	While    func(b Bool, do func())
	Sample   func(s Sampler, u Vec2) RGBA

	counts map[string]int
	indent int
}

// Unique returns a unique value everytime it is
// called on the same core for the same kind
// use this for variable names.
func (core *Core) Unique(kind string) int {
	if core.counts == nil {
		core.counts = make(map[string]int)
	}
	core.counts[kind] = core.counts[kind] + 1
	return core.counts[kind] - 1
}

func (core *Core) Indent(fn func()) {
	core.indent++
	fn()
	core.indent--
}

func (core *Core) Indentation() int {
	return core.indent
}

type IfElseChain struct {
	ElseIf func(c Bool, fn func()) IfElseChain
	Else   func(fn func())
}
