package dsl

type Texture struct {
	uint64
}

func NewTexture(v uint64) Texture {
	return Texture{v}
}

func (t Texture) Value() uint64 {
	return t.uint64
}

//Internal types available to shader cores running on
//the GPU.
type (
	Bool struct {
		Value string
	}

	Int struct {
		Value string

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
		Value string

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
		Value string

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
		Value string
		X, Y  Float

		Length     func() Float
		DistanceTo func(Vec2) Float
		Dot        func(Vec2) Float
		Normalize  func() Vec2
	}

	Vec3 struct {
		Value   string
		X, Y, Z Float

		RGB func() RGB

		Length     func() Float
		DistanceTo func(Vec3) Float
		Dot        func(Vec3) Float
		Cross      func(Vec3) Vec3
		Normalize  func() Vec3
	}

	Vec4 struct {
		Value      string
		X, Y, Z, W Float

		RGBA func() RGBA

		Length     func() Float
		DistanceTo func(Vec4) Float
		Dot        func(Vec4) Float
		Normalize  func() Vec4
	}

	Mat3 struct {
		Value string

		Transform func(Vec3) Vec3
	}

	Mat4 struct {
		Value string

		Times     func(Mat4) Mat4
		Transform func(Vec4) Vec4
	}

	RGB struct {
		Value   string
		R, G, B Float

		Vec3 func() Vec3
	}

	RGBA struct {
		Value      string
		R, G, B, A Float

		Vec4 func() Vec4
	}

	Sampler struct {
		Value string
	}
)
