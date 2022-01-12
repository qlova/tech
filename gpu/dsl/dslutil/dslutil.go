//Package dslutil has helpful utitilies for creating DSL compilers.
package dslutil

import (
	"fmt"
	"io"

	"qlova.tech/gpu/dsl"
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

// TypeSystem is a DSL type system that
// can construct DSL types from a string.
type TypeSystem interface {
	TypeOf(dsl.Type) string

	NewBool(string) dsl.Bool
	NewInt(string) dsl.Int
	NewUint(string) dsl.Uint
	NewFloat(string) dsl.Float

	NewVec2(string) dsl.Vec2
	NewVec3(string) dsl.Vec3
	NewVec4(string) dsl.Vec4

	NewMat2(string) dsl.Mat2
	NewMat3(string) dsl.Mat3
	NewMat4(string) dsl.Mat4

	NewRGB(string) dsl.RGB
	NewRGBA(string) dsl.RGBA

	NewTexture1D(string) dsl.Texture1D
	NewTexture2D(string) dsl.Texture2D
	NewTexture3D(string) dsl.Texture3D
	NewTextureCube(string) dsl.TextureCube
}

//Attributes format
// fmt.Sprintf(format, vertex.Attribute, dsl.Type)
func Attributes(w io.Writer, sl TypeSystem, aformat, vformat string) dsl.Attributes {
	return dsl.Attributes{
		Bool: func(a vertex.Attribute) (t dsl.Bool) {
			fmt.Fprintf(w, fmt.Sprintf(aformat, a, sl.TypeOf(t)))
			return sl.NewBool(fmt.Sprintf(vformat, a, sl.TypeOf(t)))
		},
		Int: func(a vertex.Attribute) (t dsl.Int) {
			fmt.Fprintf(w, fmt.Sprintf(aformat, a, sl.TypeOf(t)))
			return sl.NewInt(fmt.Sprintf(vformat, a, sl.TypeOf(t)))
		},
		Uint: func(a vertex.Attribute) (t dsl.Uint) {
			fmt.Fprintf(w, fmt.Sprintf(aformat, a, sl.TypeOf(t)))
			return sl.NewUint(fmt.Sprintf(vformat, a, sl.TypeOf(t)))
		},
		Float: func(a vertex.Attribute) (t dsl.Float) {
			fmt.Fprintf(w, fmt.Sprintf(aformat, a, sl.TypeOf(t)))
			return sl.NewFloat(fmt.Sprintf(vformat, a, sl.TypeOf(t)))
		},
		Vec2: func(a vertex.Attribute) (t dsl.Vec2) {
			fmt.Fprintf(w, fmt.Sprintf(aformat, a, sl.TypeOf(t)))
			return sl.NewVec2(fmt.Sprintf(vformat, a, sl.TypeOf(t)))
		},
		Vec3: func(a vertex.Attribute) (t dsl.Vec3) {
			fmt.Fprintf(w, fmt.Sprintf(aformat, a, sl.TypeOf(t)))
			return sl.NewVec3(fmt.Sprintf(vformat, a, sl.TypeOf(t)))
		},
		Vec4: func(a vertex.Attribute) (t dsl.Vec4) {
			fmt.Fprintf(w, fmt.Sprintf(aformat, a, sl.TypeOf(t)))
			return sl.NewVec4(fmt.Sprintf(vformat, a, sl.TypeOf(t)))
		},
		RGB: func(a vertex.Attribute) (t dsl.RGB) {
			fmt.Fprintf(w, fmt.Sprintf(aformat, a, sl.TypeOf(t)))
			return sl.NewRGB(fmt.Sprintf(vformat, a, sl.TypeOf(t)))
		},
		RGBA: func(a vertex.Attribute) (t dsl.RGBA) {
			fmt.Fprintf(w, fmt.Sprintf(aformat, a, sl.TypeOf(t)))
			return sl.NewRGBA(fmt.Sprintf(vformat, a, sl.TypeOf(t)))
		},
	}
}

func (s State) NewSetter(w io.Writer, format string) dsl.Setter {
	return dsl.Setter{
		Bool: func(a, b dsl.Bool) {
			fmt.Fprintf(w, fmt.Sprintf(format, a.Value, b.Value))
		},
		Int: func(a, b dsl.Int) {
			fmt.Fprintf(w, fmt.Sprintf(format, a.Value, b.Value))
		},
		Uint: func(a, b dsl.Uint) {
			fmt.Fprintf(w, fmt.Sprintf(format, a.Value, b.Value))
		},
		Float: func(a, b dsl.Float) {
			fmt.Fprintf(w, fmt.Sprintf(format, a.Value, b.Value))
		},
		Vec2: func(a, b dsl.Vec2) {
			fmt.Fprintf(w, fmt.Sprintf(format, a.Value, b.Value))
		},
		Vec3: func(a, b dsl.Vec3) {
			fmt.Fprintf(w, fmt.Sprintf(format, a.Value, b.Value))
		},
		Vec4: func(a, b dsl.Vec4) {
			fmt.Fprintf(w, fmt.Sprintf(format, a.Value, b.Value))
		},
		Mat2: func(a, b dsl.Mat2) {
			fmt.Fprintf(w, fmt.Sprintf(format, a.Value, b.Value))
		},
		Mat3: func(a, b dsl.Mat3) {
			fmt.Fprintf(w, fmt.Sprintf(format, a.Value, b.Value))
		},
		Mat4: func(a, b dsl.Mat4) {
			fmt.Fprintf(w, fmt.Sprintf(format, a.Value, b.Value))
		},
		RGB: func(a, b dsl.RGB) {
			fmt.Fprintf(w, fmt.Sprintf(format, a.Value, b.Value))
		},
		RGBA: func(a, b dsl.RGBA) {
			fmt.Fprintf(w, fmt.Sprintf(format, a.Value, b.Value))
		},
	}
}

type Constructor struct {
	Bool  string
	Int   string
	Uint  string
	Float string

	Vec2 string
	Vec3 string
	Vec4 string

	RGB  string
	RGBA string
}

func (s State) NewConstructor(helper Constructor) dsl.Constructor {
	return dsl.Constructor{
		Bool: func(b bool) dsl.Bool {
			return s.NewBool(fmt.Sprintf(helper.Bool, b))
		},
		Int: func(i int) dsl.Int {
			return s.NewInt(fmt.Sprintf(helper.Int, i))
		},
		Uint: func(i uint) dsl.Uint {
			return s.NewUint(fmt.Sprintf(helper.Uint, i))
		},
		Float: func(f float32) dsl.Float {
			return s.NewFloat(fmt.Sprintf(helper.Float, f))
		},
		Vec2: func(x, y dsl.Float) dsl.Vec2 {
			return s.NewVec2(fmt.Sprintf(helper.Vec2, x, y))
		},
		Vec3: func(x, y, z dsl.Float) dsl.Vec3 {
			return s.NewVec3(fmt.Sprintf(helper.Vec3, x, y, z))
		},
		Vec4: func(x, y, z, w dsl.Float) dsl.Vec4 {
			return s.NewVec4(fmt.Sprintf(helper.Vec4, x, y, z, w))
		},
		RGB: func(r, g, b dsl.Float) dsl.RGB {
			return s.NewRGB(fmt.Sprintf(helper.RGB, r, g, b))
		},
		RGBA: func(r, g, b, a dsl.Float) dsl.RGBA {
			return s.NewRGBA(fmt.Sprintf(helper.RGBA, r, g, b, a))
		},
	}
}

type Uniform struct {
	Name    string
	Pointer interface{}
}

type State struct {
	TypeSystem

	counts map[string]int
	indent int

	Uniforms []Uniform
}

func (s *State) Indent(fn func()) {
	s.indent++
	fn()
	s.indent--
}

func (s *State) Indentation() int {
	return s.indent
}

func (s *State) GetVariableName() string {
	if s.counts == nil {
		s.counts = make(map[string]int)
	}
	s.counts["var"]++
	return fmt.Sprintf("var_%v", s.counts["var"])
}

func (s *State) GetUniformName() string {
	if s.counts == nil {
		s.counts = make(map[string]int)
	}
	s.counts["uniform"]++
	return fmt.Sprintf("uniform_%v", s.counts["uniform"])
}

func (s *State) NewDefiner(w io.Writer, format string) dsl.Definer {

	definer := func(t dsl.Type) string {
		var name = s.GetVariableName()
		fmt.Fprintf(w, format, s.TypeOf(t), name, t)
		return name
	}

	return dsl.Definer{
		Bool:  func(b dsl.Bool) dsl.Bool { return s.NewBool(definer(b)) },
		Int:   func(i dsl.Int) dsl.Int { return s.NewInt(definer(i)) },
		Uint:  func(u dsl.Uint) dsl.Uint { return s.NewUint(definer(u)) },
		Float: func(f dsl.Float) dsl.Float { return s.NewFloat(definer(f)) },
		Vec2:  func(v dsl.Vec2) dsl.Vec2 { return s.NewVec2(definer(v)) },
		Vec3:  func(v dsl.Vec3) dsl.Vec3 { return s.NewVec3(definer(v)) },
		Vec4:  func(v dsl.Vec4) dsl.Vec4 { return s.NewVec4(definer(v)) },
		Mat2:  func(m dsl.Mat2) dsl.Mat2 { return s.NewMat2(definer(m)) },
		Mat3:  func(m dsl.Mat3) dsl.Mat3 { return s.NewMat3(definer(m)) },
		Mat4:  func(m dsl.Mat4) dsl.Mat4 { return s.NewMat4(definer(m)) },
		RGB:   func(r dsl.RGB) dsl.RGB { return s.NewRGB(definer(r)) },
		RGBA:  func(r dsl.RGBA) dsl.RGBA { return s.NewRGBA(definer(r)) },
	}
}

func (s *State) NewUniforms(w io.Writer, uformat, vformat string) dsl.Uniforms {

	uniform := func(pointer interface{}, t dsl.Type) string {
		var name = s.GetUniformName()
		fmt.Fprintf(w, fmt.Sprintf(uformat, name, s.TypeOf(t), s.counts["uniform"]))
		s.Uniforms = append(s.Uniforms, Uniform{name, pointer})
		return fmt.Sprintf(vformat, name)
	}

	return dsl.Uniforms{
		Bool:      func(v *bool) (t dsl.Bool) { return s.NewBool(uniform(v, t)) },
		Int:       func(v *int) (t dsl.Int) { return s.NewInt(uniform(v, t)) },
		Uint:      func(v *uint) (t dsl.Uint) { return s.NewUint(uniform(v, t)) },
		Float:     func(v *float32) (t dsl.Float) { return s.NewFloat(uniform(v, t)) },
		Vec2:      func(v *vec2.Float32) (t dsl.Vec2) { return s.NewVec2(uniform(v, t)) },
		Vec3:      func(v *vec3.Float32) (t dsl.Vec3) { return s.NewVec3(uniform(v, t)) },
		Vec4:      func(v *vec4.Float32) (t dsl.Vec4) { return s.NewVec4(uniform(v, t)) },
		Mat2:      func(v *mat2.Float32) (t dsl.Mat2) { return s.NewMat2(uniform(v, t)) },
		Mat3:      func(v *mat3.Float32) (t dsl.Mat3) { return s.NewMat3(uniform(v, t)) },
		Mat4:      func(v *mat4.Float32) (t dsl.Mat4) { return s.NewMat4(uniform(v, t)) },
		RGB:       func(v *rgb.Color) (t dsl.RGB) { return s.NewRGB(uniform(v, t)) },
		RGBA:      func(v *rgba.Color) (t dsl.RGBA) { return s.NewRGBA(uniform(v, t)) },
		Texture1D: func(v *core.Texture) (t dsl.Texture1D) { return s.NewTexture1D(uniform(v, t)) },
		Texture2D: func(v *core.Texture) (t dsl.Texture2D) { return s.NewTexture2D(uniform(v, t)) },
		Texture3D: func(v *core.Texture) (t dsl.Texture3D) { return s.NewTexture3D(uniform(v, t)) },
	}
}

type Bool struct{}

func NewBool(name string, sl TypeSystem, helper Bool) dsl.Bool {
	return dsl.Bool{Value: dsl.Value(name)}
}

type Int struct {
	LessThan  string
	MoreThan  string
	Plus      string
	Minus     string
	Times     string
	DividedBy string

	Float string
	Uint  string
	Bool  string
}

func NewInt(name string, sl TypeSystem, helper Int) dsl.Int {
	return dsl.Int{
		Value: dsl.Value(name),

		LessThan: func(i dsl.Int) dsl.Bool {
			return sl.NewBool(fmt.Sprintf(helper.LessThan, name, i.Value))
		},
		MoreThan: func(i dsl.Int) dsl.Bool {
			return sl.NewBool(fmt.Sprintf(helper.MoreThan, name, i.Value))
		},
		Plus: func(i dsl.Int) dsl.Int {
			return sl.NewInt(fmt.Sprintf(helper.Plus, name, i.Value))
		},
		Minus: func(i dsl.Int) dsl.Int {
			return sl.NewInt(fmt.Sprintf(helper.Minus, name, i.Value))
		},
		Times: func(i dsl.Int) dsl.Int {
			return sl.NewInt(fmt.Sprintf(helper.Times, name, i.Value))
		},
		DividedBy: func(i dsl.Int) dsl.Int {
			return sl.NewInt(fmt.Sprintf(helper.DividedBy, name, i.Value))
		},

		Float: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Float, name))
		},
		Uint: func() dsl.Uint {
			return sl.NewUint(fmt.Sprintf(helper.Uint, name))
		},
		Bool: func() dsl.Bool {
			return sl.NewBool(fmt.Sprintf(helper.Bool, name))
		},
	}
}

type Uint struct {
	LessThan  string
	MoreThan  string
	Plus      string
	Minus     string
	Times     string
	DividedBy string

	Float string
	Int   string
	Bool  string
}

func NewUint(name string, sl TypeSystem, helper Uint) dsl.Uint {
	return dsl.Uint{
		Value: dsl.Value(name),

		LessThan: func(i dsl.Uint) dsl.Bool {
			return sl.NewBool(fmt.Sprintf(helper.LessThan, name, i.Value))
		},
		MoreThan: func(i dsl.Uint) dsl.Bool {
			return sl.NewBool(fmt.Sprintf(helper.MoreThan, name, i.Value))
		},
		Plus: func(i dsl.Uint) dsl.Uint {
			return sl.NewUint(fmt.Sprintf(helper.Plus, name, i.Value))
		},
		Minus: func(i dsl.Uint) dsl.Uint {
			return sl.NewUint(fmt.Sprintf(helper.Minus, name, i.Value))
		},
		Times: func(i dsl.Uint) dsl.Uint {
			return sl.NewUint(fmt.Sprintf(helper.Times, name, i.Value))
		},
		DividedBy: func(i dsl.Uint) dsl.Uint {
			return sl.NewUint(fmt.Sprintf(helper.DividedBy, name, i.Value))
		},

		Float: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Float, name))
		},
		Int: func() dsl.Int {
			return sl.NewInt(fmt.Sprintf(helper.Int, name))
		},
		Bool: func() dsl.Bool {
			return sl.NewBool(fmt.Sprintf(helper.Bool, name))
		},
	}
}

type Float struct {
	LessThan string
	MoreThan string

	Plus      string
	Minus     string
	Times     string
	DividedBy string

	Uint string
	Int  string
	Bool string

	Radians     string
	Degrees     string
	Sine        string
	Cos         string
	Tan         string
	Asin        string
	Acos        string
	Atan        string
	Pow         string
	Exp         string
	Exp2        string
	Log         string
	Log2        string
	Sqrt        string
	InverseSqrt string
	Mod         string

	Min        string
	Max        string
	Clamp      string
	Lerp       string
	Step       string
	SmoothStep string
}

func NewFloat(name string, sl TypeSystem, helper Float) dsl.Float {
	return dsl.Float{
		Value: dsl.Value(name),

		LessThan: func(i dsl.Float) dsl.Bool {
			return sl.NewBool(fmt.Sprintf(helper.LessThan, name, i.Value))
		},
		MoreThan: func(i dsl.Float) dsl.Bool {
			return sl.NewBool(fmt.Sprintf(helper.MoreThan, name, i.Value))
		},
		Plus: func(i dsl.Float) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Plus, name, i.Value))
		},
		Minus: func(i dsl.Float) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Minus, name, i.Value))
		},
		Times: func(i dsl.Float) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Times, name, i.Value))
		},
		DividedBy: func(i dsl.Float) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.DividedBy, name, i.Value))
		},

		Uint: func() dsl.Uint {
			return sl.NewUint(fmt.Sprintf(helper.Uint, name))
		},
		Int: func() dsl.Int {
			return sl.NewInt(fmt.Sprintf(helper.Int, name))
		},
		Bool: func() dsl.Bool {
			return sl.NewBool(fmt.Sprintf(helper.Bool, name))
		},

		Radians: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Radians, name))
		},
		Degrees: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Degrees, name))
		},
		Sine: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Sine, name))
		},
		Cos: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Cos, name))
		},
		Tan: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Tan, name))
		},
		Asin: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Asin, name))
		},
		Acos: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Acos, name))
		},
		Atan: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Atan, name))
		},
		Pow: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Pow, name))
		},
		Exp: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Exp, name))
		},
		Exp2: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Exp2, name))
		},
		Log: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Log, name))
		},
		Log2: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Log2, name))
		},
		Sqrt: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Sqrt, name))
		},
		InverseSqrt: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.InverseSqrt, name))
		},
		Mod: func(y dsl.Float) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Mod, name, y.Value))
		},

		Min: func(i dsl.Float) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Min, name, i.Value))
		},
		Max: func(i dsl.Float) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Max, name, i.Value))
		},
		Clamp: func(min, max dsl.Float) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Clamp, name, min.Value, max.Value))
		},
		Lerp: func(to, t dsl.Float) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Lerp, name, to.Value, t.Value))
		},
		Step: func(edge dsl.Float) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Step, name, edge.Value))
		},
		SmoothStep: func(a, b dsl.Float) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.SmoothStep, name, a.Value, b.Value))
		},
	}
}

type Vec2 struct {
	X, Y string

	Length    string
	Distance  string
	Dot       string
	Normalize string
}

func NewVec2(name string, sl TypeSystem, helper Vec2) dsl.Vec2 {
	return dsl.Vec2{
		Value: dsl.Value(name),

		X: sl.NewFloat(fmt.Sprintf(helper.X, name)),
		Y: sl.NewFloat(fmt.Sprintf(helper.Y, name)),

		Length: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Length, name))
		},
		DistanceTo: func(to dsl.Vec2) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Distance, name, to.Value))
		},
		Dot: func(to dsl.Vec2) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Dot, name, to.Value))
		},
		Normalize: func() dsl.Vec2 {
			return sl.NewVec2(fmt.Sprintf(helper.Normalize, name))
		},
	}
}

type Vec3 struct {
	X, Y, Z string

	RGB string

	Length    string
	Distance  string
	Dot       string
	Normalize string
	Cross     string
}

func NewVec3(name string, sl TypeSystem, helper Vec3) dsl.Vec3 {
	return dsl.Vec3{
		Value: dsl.Value(name),

		X: sl.NewFloat(fmt.Sprintf(helper.X, name)),
		Y: sl.NewFloat(fmt.Sprintf(helper.Y, name)),
		Z: sl.NewFloat(fmt.Sprintf(helper.Z, name)),

		RGB: func() dsl.RGB {
			return sl.NewRGB(fmt.Sprintf(helper.RGB, name))
		},

		Length: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Length, name))
		},
		DistanceTo: func(to dsl.Vec3) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Distance, name, to.Value))
		},
		Dot: func(to dsl.Vec3) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Dot, name, to.Value))
		},
		Normalize: func() dsl.Vec3 {
			return sl.NewVec3(fmt.Sprintf(helper.Normalize, name))
		},
		Cross: func(to dsl.Vec3) dsl.Vec3 {
			return sl.NewVec3(fmt.Sprintf(helper.Cross, name, to.Value))
		},
	}
}

type Vec4 struct {
	X, Y, Z, W string

	RGBA string

	Length    string
	Distance  string
	Dot       string
	Normalize string
}

func NewVec4(name string, sl TypeSystem, helper Vec4) dsl.Vec4 {
	return dsl.Vec4{
		Value: dsl.Value(name),

		X: sl.NewFloat(fmt.Sprintf(helper.X, name)),
		Y: sl.NewFloat(fmt.Sprintf(helper.Y, name)),
		Z: sl.NewFloat(fmt.Sprintf(helper.Z, name)),
		W: sl.NewFloat(fmt.Sprintf(helper.W, name)),

		RGBA: func() dsl.RGBA {
			return sl.NewRGBA(fmt.Sprintf(helper.RGBA, name))
		},

		Length: func() dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Length, name))
		},
		DistanceTo: func(to dsl.Vec4) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Distance, name, to.Value))
		},
		Dot: func(to dsl.Vec4) dsl.Float {
			return sl.NewFloat(fmt.Sprintf(helper.Dot, name, to.Value))
		},
		Normalize: func() dsl.Vec4 {
			return sl.NewVec4(fmt.Sprintf(helper.Normalize, name))
		},
	}
}

type Mat2 struct {
	Times     string
	Transform string
}

func NewMat2(name string, sl TypeSystem, helper Mat2) dsl.Mat2 {
	return dsl.Mat2{
		Value: dsl.Value(name),

		Times: func(to dsl.Mat2) dsl.Mat2 {
			return sl.NewMat2(fmt.Sprintf(helper.Times, name, to.Value))
		},
		Transform: func(to dsl.Vec2) dsl.Vec2 {
			return sl.NewVec2(fmt.Sprintf(helper.Transform, name, to.Value))
		},
	}
}

type Mat3 struct {
	Transform string
}

func NewMat3(name string, sl TypeSystem, helper Mat3) dsl.Mat3 {
	return dsl.Mat3{
		Value: dsl.Value(name),

		Transform: func(to dsl.Vec3) dsl.Vec3 {
			return sl.NewVec3(fmt.Sprintf(helper.Transform, name, to.Value))
		},
	}
}

type Mat4 struct {
	Times     string
	Transform string
}

func NewMat4(name string, sl TypeSystem, helper Mat4) dsl.Mat4 {
	return dsl.Mat4{
		Value: dsl.Value(name),

		Times: func(to dsl.Mat4) dsl.Mat4 {
			return sl.NewMat4(fmt.Sprintf(helper.Times, name, to.Value))
		},
		Transform: func(to dsl.Vec4) dsl.Vec4 {
			return sl.NewVec4(fmt.Sprintf(helper.Transform, name, to.Value))
		},
	}
}

type RGB struct {
	R, G, B string

	Vec3 string
}

func NewRGB(name string, sl TypeSystem, helper RGB) dsl.RGB {
	return dsl.RGB{
		Value: dsl.Value(name),

		R: sl.NewFloat(fmt.Sprintf(helper.R, name)),
		G: sl.NewFloat(fmt.Sprintf(helper.G, name)),
		B: sl.NewFloat(fmt.Sprintf(helper.B, name)),

		Vec3: func() dsl.Vec3 {
			return sl.NewVec3(fmt.Sprintf(helper.Vec3, name))
		},
	}
}

type RGBA struct {
	R, G, B, A string

	Vec4 string
}

func NewRGBA(name string, sl TypeSystem, helper RGBA) dsl.RGBA {
	return dsl.RGBA{
		Value: dsl.Value(name),

		R: sl.NewFloat(fmt.Sprintf(helper.R, name)),
		G: sl.NewFloat(fmt.Sprintf(helper.G, name)),
		B: sl.NewFloat(fmt.Sprintf(helper.B, name)),
		A: sl.NewFloat(fmt.Sprintf(helper.A, name)),

		Vec4: func() dsl.Vec4 {
			return sl.NewVec4(fmt.Sprintf(helper.Vec4, name))
		},
	}
}

type Texture1D struct {
	Sample string
}

func NewTexture1D(name string, sl TypeSystem, helper Texture1D) dsl.Texture1D {
	return dsl.Texture1D{
		Value: dsl.Value(name),

		Sample: func(to dsl.Float) dsl.RGBA {
			return sl.NewRGBA(fmt.Sprintf(helper.Sample, name, to.Value))
		},
	}
}

type Texture2D struct {
	Sample string
}

func NewTexture2D(name string, sl TypeSystem, helper Texture2D) dsl.Texture2D {
	return dsl.Texture2D{
		Value: dsl.Value(name),

		Sample: func(to dsl.Vec2) dsl.RGBA {
			return sl.NewRGBA(fmt.Sprintf(helper.Sample, name, to.Value))
		},
	}
}

type Texture3D struct {
	Sample string
}

func NewTexture3D(name string, sl TypeSystem, helper Texture3D) dsl.Texture3D {
	return dsl.Texture3D{
		Value: dsl.Value(name),

		Sample: func(to dsl.Vec3) dsl.RGBA {
			return sl.NewRGBA(fmt.Sprintf(helper.Sample, name, to.Value))
		},
	}
}

type TextureCube struct {
	Sample string
}

func NewTextureCube(name string, sl TypeSystem, helper TextureCube) dsl.TextureCube {
	return dsl.TextureCube{
		Value: dsl.Value(name),

		Sample: func(to dsl.Vec3) dsl.RGBA {
			return sl.NewRGBA(fmt.Sprintf(helper.Sample, name, to.Value))
		},
	}
}
