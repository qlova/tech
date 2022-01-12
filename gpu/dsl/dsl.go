// Package dsl provides a domain-specific shading language.
package dsl

import (
	"bytes"

	"qlova.tech/gpu/internal/core"
	"qlova.tech/gpu/vertex"
	"qlova.tech/mat/mat2"
	"qlova.tech/mat/mat3"
	"qlova.tech/mat/mat4"
	"qlova.tech/rgb"
	"qlova.tech/rgb/rgba"
	"qlova.tech/vec/vec2"
	"qlova.tech/vec/vec3"
	"qlova.tech/vec/vec4"
)

type Shader func(Core)

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
	Float func(*float32) Float

	RGBA func(*rgba.Color) RGBA
	RGB  func(*rgb.Color) RGB

	Vec2 func(*vec2.Float32) Vec2
	Vec3 func(*vec3.Float32) Vec3
	Vec4 func(*vec4.Float32) Vec4

	Mat2 func(*mat2.Float32) Mat2
	Mat3 func(*mat3.Float32) Mat3
	Mat4 func(*mat4.Float32) Mat4

	Texture1D   func(*core.Texture) Texture1D
	Texture2D   func(*core.Texture) Texture2D
	Texture3D   func(*core.Texture) Texture3D
	TextureCube func(*core.Texture) TextureCube
}

type Definer struct {
	// Define an array of the given type.
	//Array func(size int) Definer

	Bool  func(Bool) Bool
	Int   func(Int) Int
	Uint  func(Uint) Uint
	Float func(Float) Float

	Vec2 func(Vec2) Vec2
	Vec3 func(Vec3) Vec3
	Vec4 func(Vec4) Vec4

	Mat2 func(Mat2) Mat2
	Mat3 func(Mat3) Mat3
	Mat4 func(Mat4) Mat4

	RGB  func(RGB) RGB
	RGBA func(RGBA) RGBA
}

type Constructor struct {
	Float func(f float32) Float
	Int   func(i int) Int
	Uint  func(u uint) Uint
	Bool  func(b bool) Bool

	Vec2 func(x, y Float) Vec2
	Vec3 func(x, y, z Float) Vec3
	Vec4 func(x, y, z, w Float) Vec4

	RGB  func(r, g, b Float) RGB
	RGBA func(r, g, b, a Float) RGBA
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

	Mat2 func(a, b Mat2)
	Mat3 func(a, b Mat3)
	Mat4 func(a, b Mat4)
}

// Core is the core handle for DSL
type Core struct {
	bytes.Buffer

	// Core has an embedded Definer used
	// for defining variables. These
	// variables can be mutated.
	Definer

	// Argument and Output Attributes for the function.
	// Flat means that fragment arguments will not be
	// interpolated based upon the vertex outputs.
	//
	// These functions must be called before Main.
	Arg, Out struct {
		Flat Attributes
		Attributes
	}

	// Uniforms are static for the entire frame.
	// Get allows you get a value per draw-call.
	//
	// These functions must be called before Main.
	Uniform, Get struct {

		// Get an array of the given type.
		//Array func(size int) Uniforms

		Uniforms
	}

	// New enables the construction of various GPU types.
	New Constructor

	// Set enables declerations and attributes to be mutated.
	Set Setter

	//builtin outputs.
	Position Vec4
	Fragment RGBA

	// Main is the entry-point for the function.
	Main func(fn func())

	// If directs the shader core to run fn if the
	// condition is true. It returns an IfElseChain
	// for ElseIf and Else control flow.
	If func(condition Bool, run func()) IfElseChain

	// Range is a for loop starting from min and running
	// until max, incrementing by one each time.
	Range func(min, max Int, do func(Int))

	// Discard the current fragment. Don't draw it at all.
	// Not valid in vertex functions.
	Discard func()

	// Break from the current loop.
	Break func()

	// Continue to the next iteration of the current loop.
	Continue func()

	// While condition is true run the function.
	While func(condition Bool, do func())
}

// IfElseChain is used for ElseIf and Else
// statements.
type IfElseChain struct {
	ElseIf func(c Bool, fn func()) IfElseChain
	Else   func(fn func())
}
