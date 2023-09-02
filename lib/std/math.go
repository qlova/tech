package std

import (
	"math"

	"qlova.tech/abi"
	"qlova.tech/ffi"
)

var Int struct {
	ffi.Header `linux:"libm.so.6" darwin:"libSystem.dylib"`

	Abs func(abi.Int) abi.Int               `cgo:"abs"`
	Div func(abi.Int, abi.Int) Div[abi.Int] `cgo:"div"`

	Rand        func() abi.Int `cgo:"rand"`
	SetRandSeed func(abi.Int)  `cgo:"srand"`
}

var Long struct {
	ffi.Header `linux:"libm.so.6" darwin:"libSystem.dylib"`

	Abs func(abi.Long) abi.Long                `cgo:"labs"`
	Div func(abi.Long, abi.Long) Div[abi.Long] `cgo:"ldiv"`

	RoundFloat func(abi.Float) abi.Long      `cgo:"lround"`
	Round      func(abi.Double) abi.Long     `cgo:"lround"`
	RoundLong  func(abi.DoubleLong) abi.Long `cgo:"lround"`
}

var LongLong struct {
	ffi.Header `linux:"libm.so.6" darwin:"libSystem.dylib"`

	Abs func(abi.LongLong) abi.LongLong                    `cgo:"llabs"`
	Div func(abi.LongLong, abi.LongLong) Div[abi.LongLong] `cgo:"lldiv"`

	RoundFloat func(abi.Float) abi.Long      `cgo:"llround"`
	Round      func(abi.Double) abi.Long     `cgo:"llround"`
	RoundLong  func(abi.DoubleLong) abi.Long `cgo:"llround"`
}

var IntMax struct {
	ffi.Header `linux:"libm.so.6" darwin:"libSystem.dylib"`

	Abs func(abi.IntMax) abi.IntMax                  `cgo:"imaxabs"`
	Div func(abi.IntMax, abi.IntMax) Div[abi.IntMax] `cgo:"imaxdiv"`
}

var Double struct {
	ffi.Header `linux:"libm.so.6" darwin:"libSystem.dylib"`

	Abs                func(abi.Double) abi.Double                         `cgo:"fabs"`
	Mod                func(abi.Double, abi.Double) abi.Double             `cgo:"fmod"`
	Remainder          func(abi.Double, abi.Double) abi.Double             `cgo:"remainder"`
	RemainderQuotient  func(abi.Double, abi.Double) (abi.Double, abi.Int)  `cgo:"remquo"`
	FusedMuliplyAdd    func(abi.Double, abi.Double, abi.Double) abi.Double `cgo:"fma"`
	Max                func(abi.Double, abi.Double) abi.Double             `cgo:"fmax"`
	Min                func(abi.Double, abi.Double) abi.Double             `cgo:"fmin"`
	PositiveDifference func(abi.Double, abi.Double) abi.Double             `cgo:"fdim"`
	Nan                func(abi.String) abi.Double                         `cgo:"nan"`

	Exp   func(abi.Double) abi.Double `cgo:"exp"`
	Exp2  func(abi.Double) abi.Double `cgo:"exp2"`
	Expm1 func(abi.Double) abi.Double `cgo:"expm1"`
	Log   func(abi.Double) abi.Double `cgo:"log"`
	Log10 func(abi.Double) abi.Double `cgo:"log10"`
	Log2  func(abi.Double) abi.Double `cgo:"log2"`
	Log1p func(abi.Double) abi.Double `cgo:"log1p"`

	Pow   func(abi.Double, abi.Double) abi.Double `cgo:"pow"`
	Sqrt  func(abi.Double) abi.Double             `cgo:"sqrt"`
	Cbrt  func(abi.Double) abi.Double             `cgo:"cbrt"`
	Hypot func(abi.Double, abi.Double) abi.Double `cgo:"hypot"`

	Sin   func(abi.Double) abi.Double             `cgo:"sin"`
	Cos   func(abi.Double) abi.Double             `cgo:"cos"`
	Tan   func(abi.Double) abi.Double             `cgo:"tan"`
	Asin  func(abi.Double) abi.Double             `cgo:"asin"`
	Acos  func(abi.Double) abi.Double             `cgo:"acos"`
	Atan  func(abi.Double) abi.Double             `cgo:"atan"`
	Atan2 func(abi.Double, abi.Double) abi.Double `cgo:"atan2"`

	Sinh  func(abi.Double) abi.Double `cgo:"sinh"`
	Cosh  func(abi.Double) abi.Double `cgo:"cosh"`
	Tanh  func(abi.Double) abi.Double `cgo:"tanh"`
	Asinh func(abi.Double) abi.Double `cgo:"asinh"`
	Acosh func(abi.Double) abi.Double `cgo:"acosh"`
	Atanh func(abi.Double) abi.Double `cgo:"atanh"`

	Erf    func(abi.Double) abi.Double `cgo:"erf"`
	Erfc   func(abi.Double) abi.Double `cgo:"erfc"`
	GammaT func(abi.Double) abi.Double `cgo:"tgamma"`
	GammaL func(abi.Double) abi.Double `cgo:"lgamma"`

	Ceil      func(abi.Double) abi.Double   `cgo:"ceil"`
	Floor     func(abi.Double) abi.Double   `cgo:"floor"`
	Trunc     func(abi.Double) abi.Double   `cgo:"trunc"`
	Round     func(abi.Double) abi.Double   `cgo:"round"`
	NearbyInt func(abi.Double) abi.Double   `cgo:"nearbyint"`
	Int       func(abi.Double) abi.Double   `cgo:"rint"`
	Long      func(abi.Double) abi.Long     `cgo:"lrint"`
	LongLong  func(abi.Double) abi.LongLong `cgo:"llrint"`

	Frexp      func(abi.Double) (abi.Double, abi.Int)      `cgo:"frexp"`
	Ldexp      func(abi.Double, abi.Int) abi.Double        `cgo:"ldexp"`
	Modf       func(abi.Double) (abi.Double, abi.Double)   `cgo:"modf"`
	Scale      func(abi.Double, abi.Double) abi.Double     `cgo:"scalbn"`
	ScaleLong  func(abi.Double, abi.Long) abi.Double       `cgo:"scalbln"`
	LogInt     func(abi.Double) abi.Int                    `cgo:"logb"`
	Logb       func(abi.Double) abi.Double                 `cgo:"logb"`
	NextAfter  func(abi.Double, abi.Double) abi.Double     `cgo:"nextafter"`
	NextToward func(abi.Double, abi.DoubleLong) abi.Double `cgo:"nexttoward"`
	CopySign   func(abi.Double, abi.Double) abi.Double     `cgo:"copysign"`
}

var Float struct {
	ffi.Header `linux:"libm.so.6" darwin:"libSystem.dylib"`

	Abs                func(abi.Float) abi.Float                       `cgo:"fabsf"`
	Mod                func(abi.Float, abi.Float) abi.Float            `cgo:"fmodf"`
	Remainder          func(abi.Float, abi.Float) abi.Float            `cgo:"remainderf"`
	RemainderQuotient  func(abi.Float, abi.Float) (abi.Float, abi.Int) `cgo:"remquof"`
	FusedMuliplyAdd    func(abi.Float, abi.Float, abi.Float) abi.Float `cgo:"fmaf"`
	Max                func(abi.Float, abi.Float) abi.Float            `cgo:"fmaxf"`
	Min                func(abi.Float, abi.Float) abi.Float            `cgo:"fminf"`
	PositiveDifference func(abi.Float, abi.Float) abi.Float            `cgo:"fdimf"`
	Nan                func(abi.String) abi.Float                      `cgo:"nanf"`

	Exp   func(abi.Float) abi.Float `cgo:"expf"`
	Exp2  func(abi.Float) abi.Float `cgo:"exp2f"`
	Expm1 func(abi.Float) abi.Float `cgo:"expm1f"`
	Log   func(abi.Float) abi.Float `cgo:"logf"`
	Log10 func(abi.Float) abi.Float `cgo:"log10f"`
	Log2  func(abi.Float) abi.Float `cgo:"log2f"`
	Log1p func(abi.Float) abi.Float `cgo:"log1pf"`

	Pow   func(abi.Float, abi.Float) abi.Float `cgo:"powf"`
	Sqrt  func(abi.Float) abi.Float            `cgo:"sqrtf"`
	Cbrt  func(abi.Float) abi.Float            `cgo:"cbrtf"`
	Hypot func(abi.Float, abi.Float) abi.Float `cgo:"hypotf"`

	Sin   func(abi.Float) abi.Float            `cgo:"sinf"`
	Cos   func(abi.Float) abi.Float            `cgo:"cosf"`
	Tan   func(abi.Float) abi.Float            `cgo:"tanf"`
	Asin  func(abi.Float) abi.Float            `cgo:"asinf"`
	Acos  func(abi.Float) abi.Float            `cgo:"acosf"`
	Atan  func(abi.Float) abi.Float            `cgo:"atanf"`
	Atan2 func(abi.Float, abi.Float) abi.Float `cgo:"atan2f"`

	Sinh  func(abi.Float) abi.Float `cgo:"sinhf"`
	Cosh  func(abi.Float) abi.Float `cgo:"coshf"`
	Tanh  func(abi.Float) abi.Float `cgo:"tanhf"`
	Asinh func(abi.Float) abi.Float `cgo:"asinhf"`
	Acosh func(abi.Float) abi.Float `cgo:"acoshf"`
	Atanh func(abi.Float) abi.Float `cgo:"atanhf"`

	Erf    func(abi.Float) abi.Float `cgo:"erff"`
	Erfc   func(abi.Float) abi.Float `cgo:"erfcf"`
	GammaT func(abi.Float) abi.Float `cgo:"tgammaf"`
	GammaL func(abi.Float) abi.Float `cgo:"lgammaf"`

	Ceil      func(abi.Float) abi.Float    `cgo:"ceilf"`
	Floor     func(abi.Float) abi.Float    `cgo:"floorf"`
	Trunc     func(abi.Float) abi.Float    `cgo:"truncf"`
	Round     func(abi.Float) abi.Float    `cgo:"roundf"`
	NearbyInt func(abi.Float) abi.Float    `cgo:"nearbyintf"`
	Int       func(abi.Float) abi.Float    `cgo:"rintf"`
	Long      func(abi.Float) abi.Long     `cgo:"lrintf"`
	LongLong  func(abi.Float) abi.LongLong `cgo:"llrintf"`

	Frexp      func(abi.Float) (abi.Float, abi.Int)      `cgo:"frexpf"`
	Ldexp      func(abi.Float, abi.Int) abi.Float        `cgo:"ldexpf"`
	Modf       func(abi.Float) (abi.Float, abi.Float)    `cgo:"modff"`
	Scale      func(abi.Float, abi.Float) abi.Float      `cgo:"scalbnf"`
	ScaleLong  func(abi.Float, abi.Long) abi.Float       `cgo:"scalblnf"`
	LogInt     func(abi.Float) abi.Int                   `cgo:"logbf"`
	Logb       func(abi.Float) abi.Float                 `cgo:"logbf"`
	NextAfter  func(abi.Float, abi.Float) abi.Float      `cgo:"nextafterf"`
	NextToward func(abi.Float, abi.DoubleLong) abi.Float `cgo:"nexttowardf"`
	CopySign   func(abi.Float, abi.Float) abi.Float      `cgo:"copysignf"`
}

var DoubleLong struct {
	ffi.Header `linux:"libm.so.6" darwin:"libSystem.dylib"`

	Abs                func(abi.DoubleLong) abi.DoubleLong                                 `cgo:"fabsl"`
	Mod                func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong                 `cgo:"fmodl"`
	Remainder          func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong                 `cgo:"remainderl"`
	RemainderQuotient  func(abi.DoubleLong, abi.DoubleLong) (abi.DoubleLong, abi.Int)      `cgo:"remquol"`
	FusedMuliplyAdd    func(abi.DoubleLong, abi.DoubleLong, abi.DoubleLong) abi.DoubleLong `cgo:"fmal"`
	Max                func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong                 `cgo:"fmaxl"`
	Min                func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong                 `cgo:"fminl"`
	PositiveDifference func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong                 `cgo:"fdiml"`
	Nan                func(abi.String) abi.DoubleLong                                     `cgo:"nanl"`

	Exp   func(abi.DoubleLong) abi.DoubleLong `cgo:"expl"`
	Exp2  func(abi.DoubleLong) abi.DoubleLong `cgo:"exp2l"`
	Expm1 func(abi.DoubleLong) abi.DoubleLong `cgo:"expm1l"`
	Log   func(abi.DoubleLong) abi.DoubleLong `cgo:"logl"`
	Log10 func(abi.DoubleLong) abi.DoubleLong `cgo:"log10l"`
	Log2  func(abi.DoubleLong) abi.DoubleLong `cgo:"log2l"`
	Log1p func(abi.DoubleLong) abi.DoubleLong `cgo:"log1pl"`

	Pow   func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong `cgo:"powl"`
	Sqrt  func(abi.DoubleLong) abi.DoubleLong                 `cgo:"sqrtl"`
	Cbrt  func(abi.DoubleLong) abi.DoubleLong                 `cgo:"cbrtl"`
	Hypot func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong `cgo:"hypotl"`

	Sin   func(abi.DoubleLong) abi.DoubleLong                 `cgo:"sinl"`
	Cos   func(abi.DoubleLong) abi.DoubleLong                 `cgo:"cosl"`
	Tan   func(abi.DoubleLong) abi.DoubleLong                 `cgo:"tanl"`
	Asin  func(abi.DoubleLong) abi.DoubleLong                 `cgo:"asinl"`
	Acos  func(abi.DoubleLong) abi.DoubleLong                 `cgo:"acosl"`
	Atan  func(abi.DoubleLong) abi.DoubleLong                 `cgo:"atanl"`
	Atan2 func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong `cgo:"atan2l"`

	Sinh  func(abi.DoubleLong) abi.DoubleLong `cgo:"sinhl"`
	Cosh  func(abi.DoubleLong) abi.DoubleLong `cgo:"coshl"`
	Tanh  func(abi.DoubleLong) abi.DoubleLong `cgo:"tanhl"`
	Asinh func(abi.DoubleLong) abi.DoubleLong `cgo:"asinhl"`
	Acosh func(abi.DoubleLong) abi.DoubleLong `cgo:"acoshl"`
	Atanh func(abi.DoubleLong) abi.DoubleLong `cgo:"atanhl"`

	Erf    func(abi.DoubleLong) abi.DoubleLong `cgo:"erfl"`
	Erfc   func(abi.DoubleLong) abi.DoubleLong `cgo:"erfcl"`
	GammaT func(abi.DoubleLong) abi.DoubleLong `cgo:"tgammal"`
	GammaL func(abi.DoubleLong) abi.DoubleLong `cgo:"lgammal"`

	Ceil      func(abi.DoubleLong) abi.DoubleLong `cgo:"ceill"`
	Floor     func(abi.DoubleLong) abi.DoubleLong `cgo:"floorl"`
	Trunc     func(abi.DoubleLong) abi.DoubleLong `cgo:"truncl"`
	Round     func(abi.DoubleLong) abi.DoubleLong `cgo:"roundl"`
	NearbyInt func(abi.DoubleLong) abi.DoubleLong `cgo:"nearbyintl"`
	Int       func(abi.DoubleLong) abi.DoubleLong `cgo:"rintl"`
	Long      func(abi.DoubleLong) abi.Long       `cgo:"lrintl"`
	LongLong  func(abi.DoubleLong) abi.LongLong   `cgo:"llrintl"`

	Frexp      func(abi.DoubleLong) (abi.DoubleLong, abi.Int)        `cgo:"frexpl"`
	Ldexp      func(abi.DoubleLong, abi.Int) abi.DoubleLong          `cgo:"ldexpl"`
	Modf       func(abi.DoubleLong) (abi.DoubleLong, abi.DoubleLong) `cgo:"modfl"`
	Scale      func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong   `cgo:"scalbnl"`
	ScaleLong  func(abi.DoubleLong, abi.Long) abi.DoubleLong         `cgo:"scalblnl"`
	LogInt     func(abi.DoubleLong) abi.Int                          `cgo:"logbl"`
	Logb       func(abi.DoubleLong) abi.DoubleLong                   `cgo:"logbl"`
	NextAfter  func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong   `cgo:"nextafterl"`
	NextToward func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong   `cgo:"nexttowardl"`
	CopySign   func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong   `cgo:"copysignl"`
}

func ClassifyFloat(f abi.Float) abi.FloatClass {
	switch {
	case f == 0:
		return abi.FloatIsZero
	case math.IsNaN(float64(f)):
		return abi.FloatIsNaN
	case math.IsInf(float64(f), 1), math.IsInf(float64(f), -1):
		return abi.FloatIsInfinite
	case f == -0 || f == +0:
		return abi.FloatIsSubnormal
	default:
		return abi.FloatIsNormal
	}
}

func IsFinite(f abi.Float) bool {
	return !math.IsNaN(float64(f)) && !math.IsInf(float64(f), 1) && !math.IsInf(float64(f), -1)
}

func IsInf(f abi.Float) bool {
	return math.IsInf(float64(f), 1) || math.IsInf(float64(f), -1)
}

func IsNaN(f abi.Float) bool {
	return math.IsNaN(float64(f))
}

func IsNormal(f abi.Float) bool {
	return !math.IsNaN(float64(f)) && !math.IsInf(float64(f), 1) && !math.IsInf(float64(f), -1) && f != -0 && f != +0
}

func SignBit(f abi.Float) bool {
	return math.Signbit(float64(f))
}

func IsGreater(a, b abi.Float) bool {
	return a > b
}

func IsGreaterEqual(a, b abi.Float) bool {
	return a >= b
}

func IsLess(a, b abi.Float) bool {
	return a < b
}

func IsLessEqual(a, b abi.Float) bool {
	return a <= b
}

func IsLessGreater(a, b abi.Float) bool {
	return a != b
}

func IsUnordered(a, b abi.Float) bool {
	return math.IsNaN(float64(a)) || math.IsNaN(float64(b))
}

func HugeFloat() abi.Float {
	return abi.Float(math.Inf(1))
}

func HugeDouble() abi.Double {
	return abi.Double(math.Inf(1))
}

func HugeDoubleLong() abi.DoubleLong {
	return abi.DoubleLong(math.Inf(1))
}

func NaN() abi.Float {
	return abi.Float(math.NaN())
}

func Infinity() abi.Float {
	return abi.Float(math.Inf(1))
}
