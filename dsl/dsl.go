/*
	Package dsl provides the DSL Shading Language.

	DSL is a Domain Specific Shading Language, this implementation
	exists within Go, but DSL could be implemented in other languages too.

	DSL functions are called with a Core, which is a processing core
	on the GPU and enables shaders to be written that run on the GPU.

	Vertex Shader

	The vertex shader is called once for each vertex in the mesh being
	rendered. It can accept various attributes for each vertex, for example
	the position, normal, texture coordinates, etc.

	The simplist example of a vertex function takes a vertex position for
	each vertex and sets the core's position to that position.

		func(core dsl.Core) {
			core.Set.Vec4(core.Position, core.In.Vec4(vertex.Position))
		}

	Fragment Shader

	The fragment shader is called once for each pixel/fragment in the
	filled shape specified by the vertex shader. The purpose of this
	shader is to determine the color of the pixel/fragment and to
	configure the material properties that will be passed to the
	lighting shader.

	The simplist example of a fragment function sets the core's
	fragment color to red.

		func(core dsl.Core) {
			f := core.New.Float
			red := core.RGBA(f(1), f(0), f(0), f(1))

			core.Set.RGBA(core.Fragment, red)
		}

	Limitations

	No support for arrays, so use a Vec type instead (if you can).

*/
package dsl

import (
	"qlova.tech/rgb"
	"qlova.tech/xy"
	"qlova.tech/xyz"
)

// Hint is a hint that can be used to configure the
// behaviour of the GPU when rendering.
type Hint uint64

// Hints.
const (
	Cull Hint = 1 << iota
	Front
	Back
	Blend
	Wireframe
	Shaded
)

//TODO
type Reader interface{}

// Shader is a function that instructs a core on how to
// process a vertex or fragment.
type Shader func(Core)

type Attribute string

// Attributes that can be passed in or out of a Shader.
type Attributes struct {
	Bool  func(Attribute) Bool
	Int   func(Attribute) Int
	Uint  func(Attribute) Uint
	Float func(Attribute) Float

	RGB func(Attribute) RGB

	Vec2 func(Attribute) Vec2
	Vec3 func(Attribute) Vec3
	Vec4 func(Attribute) Vec4
}

// Texture is any struct that embeds this type.
// It has no implementations and simply acts
// as a flagging mechanisim.
type Texture interface {
	texture()
}

// Uniforms are Go values that can be passed
// to the shader, either once-per frame, or
// once-per draw call.
type Uniforms struct {
	Bool  func(*bool) Bool
	Int   func(*int32) Int
	Uint  func(*uint32) Uint
	Float func(*float32) Float

	RGB func(*rgb.Color) RGB

	Vec2 func(*xy.Vector) Vec2
	Vec3 func(*xyz.Vector) Vec3
	//Vec4 func(*vec4.Float32) Vec4

	//Mat2 func(*mat2.Float32) Mat2
	Mat3 func(*xy.Transform) Mat3
	Mat4 func(*xyz.Transform) Mat4

	Texture1D   func(Texture) Texture1D
	Texture2D   func(Texture) Texture2D
	Texture3D   func(Texture) Texture3D
	TextureCube func(Texture) TextureCube
}

// Definer can define variables with a
// specific type.
type Definer struct {
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

	RGB func(RGB) RGB
}

// enforcable literals.
type (
	float    float32
	integer  int32
	unsigned uint32
	boolean  bool
)

//Literals is required to be able to create constructor functions.
func Literals(
	a func(bool) Bool,
	b func(int32) Int,
	c func(uint32) Uint,
	d func(float32) Float,

) (
	func(boolean) Bool,
	func(integer) Int,
	func(unsigned) Uint,
	func(float) Float,

) {
	return func(b boolean) Bool {
			return a(bool(b))
		},
		func(i integer) Int {
			return b(int32(i))
		},
		func(u unsigned) Uint {
			return c(uint32(u))
		},
		func(f float) Float {
			return d(float32(f))
		}

}

// Constructor can be used to construct
// shader values, either using literals
// or constructed from multiple values.
type Constructor struct {
	Float func(float) Float
	Int   func(integer) Int
	Uint  func(unsigned) Uint
	Bool  func(boolean) Bool

	Vec2 func(x, y Float) Vec2
	Vec3 func(x, y, z Float) Vec3
	Vec4 func(x, y, z, w Float) Vec4

	RGB func(r, g, b, a Float) RGB
}

// Setter can be used to set shader values.
// Only variables returned from Var can be
// mutated this way.
type Setter struct {
	Bool  func(a, b Bool)
	Int   func(a, b Int)
	Uint  func(a, b Uint)
	Float func(a, b Float)

	RGB func(a, b RGB)

	Vec2 func(a, b Vec2)
	Vec3 func(a, b Vec3)
	Vec4 func(a, b Vec4)

	Mat2 func(a, b Mat2)
	Mat3 func(a, b Mat3)
	Mat4 func(a, b Mat4)
}

// Core is a handle to a gpu shader core, can be
// instructed on how to draw a mesh.
type Core struct {

	// Core has an embedded Constructor used
	// for constructing values.
	Constructor

	// Input and Output attributes for the function.
	// Flat means that fragment arguments will not be
	// interpolated based upon the vertex outputs.
	In, Out Attributes

	// Uniforms are static for the entire frame.
	// Get allows you get a value per draw-call.
	Uniform, Get Uniforms

	// Var enables the declaration of mutable
	// variables.
	Var Definer

	// Set sets the value of a variable or attribute.
	Set Setter

	// Position of the vertex that the GPU will draw.
	Position Vec3

	// Fragment color of the 'pixel' that will drawn.
	// Can be affected by lighting and other effects.
	Fragment RGB

	// Normal of the fragment that will be passed
	// to the lighting shader.
	Normal Vec3

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
