package glsl

import (
	"fmt"

	"qlova.tech/gpu/dsl"
	"qlova.tech/gpu/dsl/dslutil"
)

func (s Program) TypeOf(t dsl.Type) string {
	switch t.(type) {
	case dsl.Bool:
		return "bool"
	case dsl.Int:
		return "int"
	case dsl.Uint:
		return "uint"
	case dsl.Float:
		return "float"
	case dsl.Vec2:
		return "vec2"
	case dsl.Vec3:
		return "vec3"
	case dsl.Vec4:
		return "vec4"
	case dsl.Mat2:
		return "mat2"
	case dsl.Mat3:
		return "mat3"
	case dsl.Mat4:
		return "mat4"
	case dsl.RGB:
		return "vec3"
	case dsl.RGBA:
		return "vec4"
	case dsl.Texture1D:
		return "sampler1D"
	case dsl.Texture2D:
		return "sampler2D"
	case dsl.Texture3D:
		return "sampler3D"
	case dsl.TextureCube:
		return "samplerCube"
	}
	panic(fmt.Sprintf("unknown dsl type: %T", t))
}

func (s Program) NewBool(name string) dsl.Bool {
	return dslutil.NewBool(name, s, dslutil.Bool{})
}

func (s Program) NewInt(name string) dsl.Int {
	return dslutil.NewInt(name, s, dslutil.Int{
		LessThan:  "(%v < %v)",
		MoreThan:  "(%v > %v)",
		Plus:      "(%v + %v)",
		Minus:     "(%v - %v)",
		Times:     "(%v * %v)",
		DividedBy: "(%v / %v)",

		Float: "((float)%v)",
		Uint:  "((uint)%v)",
		Bool:  "((bool)%v)",
	})
}

func (s Program) NewUint(name string) dsl.Uint {
	return dslutil.NewUint(name, s, dslutil.Uint{
		LessThan:  "(%v < %v)",
		MoreThan:  "(%v > %v)",
		Plus:      "(%v + %v)",
		Minus:     "(%v - %v)",
		Times:     "(%v * %v)",
		DividedBy: "(%v / %v)",

		Float: "((float)%v)",
		Int:   "((int)%v)",
		Bool:  "((bool)%v)",
	})
}

func (s Program) NewFloat(name string) dsl.Float {
	return dslutil.NewFloat(name, s, dslutil.Float{
		LessThan: "(%v < %v)",
		MoreThan: "(%v > %v)",

		Plus:      "(%v + %v)",
		Minus:     "(%v - %v)",
		Times:     "(%v * %v)",
		DividedBy: "(%v / %v)",

		Uint: "((uint)%v)",
		Int:  "((int)%v)",
		Bool: "((bool)%v)",

		Radians: "radians(%v)",
		Degrees: "degrees(%v)",
		Sine:    "sin(%v)",
		Cos:     "cos(%v)",
		Tan:     "tan(%v)",
		Asin:    "asin(%v)",
		Acos:    "acos(%v)",
		Atan:    "atan(%v)",

		Pow:         "pow(%v, %v)",
		Exp:         "exp(%v)",
		Exp2:        "exp2(%v)",
		Sqrt:        "sqrt(%v)",
		InverseSqrt: "inversesqrt(%v)",
		Mod:         "mod(%v, %v)",

		Min:        "min(%v, %v)",
		Max:        "max(%v, %v)",
		Clamp:      "clamp(%v, %v, %v)",
		Lerp:       "mix(%v, %v, %v)",
		Step:       "step(%[2]v, %[1]v)",
		SmoothStep: "smoothstep(%[2]v, %[3]v, %[1]v)",
	})
}

func (s Program) NewVec2(name string) dsl.Vec2 {
	return dslutil.NewVec2(name, s, dslutil.Vec2{
		X: "%v.x", Y: "%v.y",

		Length:    "length(%v)",
		Distance:  "distance(%v, %v)",
		Dot:       "dot(%v, %v)",
		Normalize: "normalize(%v)",
	})
}

func (s Program) NewVec3(name string) dsl.Vec3 {
	return dslutil.NewVec3(name, s, dslutil.Vec3{
		X: "%v.x", Y: "%v.y", Z: "%v.z",

		Length:    "length(%v)",
		Distance:  "distance(%v, %v)",
		Dot:       "dot(%v, %v)",
		Normalize: "normalize(%v)",
		Cross:     "cross(%v, %v)",
	})
}

func (s Program) NewVec4(name string) dsl.Vec4 {
	return dslutil.NewVec4(name, s, dslutil.Vec4{
		X: "%v.x", Y: "%v.y", Z: "%v.z", W: "%v.w",

		Length:    "length(%v)",
		Distance:  "distance(%v, %v)",
		Dot:       "dot(%v, %v)",
		Normalize: "normalize(%v)",
	})
}

func (s Program) NewMat2(name string) dsl.Mat2 {
	return dslutil.NewMat2(name, s, dslutil.Mat2{
		Times:     "(%v * %v)",
		Transform: "(%v * %v)",
	})
}

func (s Program) NewMat3(name string) dsl.Mat3 {
	return dslutil.NewMat3(name, s, dslutil.Mat3{
		Transform: "(%v * %v)",
	})
}

func (s Program) NewMat4(name string) dsl.Mat4 {
	return dslutil.NewMat4(name, s, dslutil.Mat4{
		Times:     "(%v * %v)",
		Transform: "(%v * %v)",
	})
}

func (s Program) NewRGB(name string) dsl.RGB {
	return dslutil.NewRGB(name, s, dslutil.RGB{
		R: "%v.r", G: "%v.g", B: "%v.b",

		Vec3: "%v",
	})
}

func (s Program) NewRGBA(name string) dsl.RGBA {
	return dslutil.NewRGBA(name, s, dslutil.RGBA{
		R: "%v.r", G: "%v.g", B: "%v.b", A: "%v.a",

		Vec4: "%v",
	})
}

func (s Program) NewTexture1D(name string) dsl.Texture1D {
	return dslutil.NewTexture1D(name, s, dslutil.Texture1D{
		Sample: "sample1D(%v, %v)",
	})
}

func (s Program) NewTexture2D(name string) dsl.Texture2D {
	return dslutil.NewTexture2D(name, s, dslutil.Texture2D{
		Sample: "sample2D(%v, %v)",
	})
}

func (s Program) NewTexture3D(name string) dsl.Texture3D {
	return dslutil.NewTexture3D(name, s, dslutil.Texture3D{
		Sample: "sample3D(%v, %v)",
	})
}

func (s Program) NewTextureCube(name string) dsl.TextureCube {
	return dslutil.NewTextureCube(name, s, dslutil.TextureCube{
		Sample: "sampleCube(%v, %v)",
	})
}
