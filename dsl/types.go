package dsl

type Type interface {
	dslType()
}

type Value string

func (v Value) String() string {
	return string(v)
}

//Internal types available to shader cores running on
//the GPU.
type (
	Bool struct {
		Type
		Value
	}

	Int struct {
		Type
		Value

		LessThan func(Int) Bool
		MoreThan func(Int) Bool

		Plus      func(Int) Int
		Minus     func(Int) Int
		Times     func(Int) Int
		DividedBy func(Int) Int

		Float func() Float
		Uint  func() Uint
		Bool  func() Bool
	}

	Uint struct {
		Type
		Value

		LessThan func(Uint) Bool
		MoreThan func(Uint) Bool

		Plus      func(Uint) Uint
		Minus     func(Uint) Uint
		Times     func(Uint) Uint
		DividedBy func(Uint) Uint

		Float func() Float
		Int   func() Int
		Bool  func() Bool
	}

	Float struct {
		Type
		Value

		LessThan func(Float) Bool
		MoreThan func(Float) Bool

		Plus      func(Float) Float
		Minus     func(Float) Float
		Times     func(Float) Float
		DividedBy func(Float) Float

		Uint func() Uint
		Int  func() Int
		Bool func() Bool

		Radians func() Float
		Degrees func() Float
		Sine    func() Float
		Cos     func() Float
		Tan     func() Float
		Asin    func() Float
		Acos    func() Float
		Atan    func() Float

		Pow         func() Float
		Exp         func() Float
		Exp2        func() Float
		Log         func() Float
		Log2        func() Float
		Sqrt        func() Float
		InverseSqrt func() Float
		Mod         func(y Float) Float

		Min        func(Float) Float
		Max        func(Float) Float
		Clamp      func(min, max Float) Float
		Lerp       func(to Float, t Float) Float
		Step       func(edge Float) Float
		SmoothStep func(edge1, edge2 Float) Float
	}

	Vec2 struct {
		Type
		Value

		X, Y Float

		Length     func() Float
		DistanceTo func(Vec2) Float
		Dot        func(Vec2) Float
		Normalize  func() Vec2
	}

	Vec3 struct {
		Type
		Value

		X, Y, Z Float

		Length     func() Float
		DistanceTo func(Vec3) Float
		Dot        func(Vec3) Float
		Cross      func(Vec3) Vec3
		Normalize  func() Vec3
	}

	Vec4 struct {
		Type
		Value

		X, Y, Z, W Float

		RGB func() RGB

		Length     func() Float
		DistanceTo func(Vec4) Float
		Dot        func(Vec4) Float
		Normalize  func() Vec4
	}

	Mat2 struct {
		Type
		Value

		Times     func(Mat2) Mat2
		Transform func(Vec2) Vec2
	}

	Mat3 struct {
		Type
		Value

		Transform func(Vec3) Vec3
	}

	Mat4 struct {
		Type
		Value

		Times func(Mat4) Mat4

		TransformNormal func(Vec3) Vec3
		Transform       func(Vec3) Vec3
	}

	RGB struct {
		Type
		Value

		R, G, B, A Float

		Vec4 func() Vec4
	}

	Texture1D struct {
		Type
		Value

		Sample func(Float) RGB
	}
	Texture2D struct {
		Type
		Value

		Sample func(Vec2) RGB
	}
	Texture3D struct {
		Type
		Value

		Sample func(Vec3) RGB
	}
	TextureCube struct {
		Type
		Value

		Sample func(Vec3) RGB
	}
)
