package std

import (
	"math"

	"qlova.tech/abi"
)

var Int struct {
	LibM

	Abs func(abi.Int) abi.Int               `ffi:"abs"`
	Div func(abi.Int, abi.Int) Div[abi.Int] `ffi:"div"`

	Rand        func() abi.Int `ffi:"rand"`
	SetRandSeed func(abi.Int)  `ffi:"srand"`
}

var Long struct {
	LibM

	Abs func(abi.Long) abi.Long                `ffi:"labs"`
	Div func(abi.Long, abi.Long) Div[abi.Long] `ffi:"ldiv"`

	RoundFloat func(abi.Float) abi.Long      `ffi:"lround"`
	Round      func(abi.Double) abi.Long     `ffi:"lround"`
	RoundLong  func(abi.DoubleLong) abi.Long `ffi:"lround"`
}

var LongLong struct {
	LibM

	Abs func(abi.LongLong) abi.LongLong                    `ffi:"llabs"`
	Div func(abi.LongLong, abi.LongLong) Div[abi.LongLong] `ffi:"lldiv"`

	RoundFloat func(abi.Float) abi.Long      `ffi:"llround"`
	Round      func(abi.Double) abi.Long     `ffi:"llround"`
	RoundLong  func(abi.DoubleLong) abi.Long `ffi:"llround"`
}

var IntMax struct {
	LibM

	Abs func(abi.IntMax) abi.IntMax                  `ffi:"imaxabs"`
	Div func(abi.IntMax, abi.IntMax) Div[abi.IntMax] `ffi:"imaxdiv"`
}

var Double struct {
	LibM

	Abs                func(abi.Double) abi.Double                         `ffi:"fabs"`
	Mod                func(abi.Double, abi.Double) abi.Double             `ffi:"fmod"`
	Remainder          func(abi.Double, abi.Double) abi.Double             `ffi:"remainder"`
	RemainderQuotient  func(abi.Double, abi.Double) (abi.Double, abi.Int)  `ffi:"remquo"`
	FusedMuliplyAdd    func(abi.Double, abi.Double, abi.Double) abi.Double `ffi:"fma"`
	Max                func(abi.Double, abi.Double) abi.Double             `ffi:"fmax"`
	Min                func(abi.Double, abi.Double) abi.Double             `ffi:"fmin"`
	PositiveDifference func(abi.Double, abi.Double) abi.Double             `ffi:"fdim"`
	Nan                func(abi.String) abi.Double                         `ffi:"nan"`

	Exp   func(abi.Double) abi.Double `ffi:"exp"`
	Exp2  func(abi.Double) abi.Double `ffi:"exp2"`
	Expm1 func(abi.Double) abi.Double `ffi:"expm1"`
	Log   func(abi.Double) abi.Double `ffi:"log"`
	Log10 func(abi.Double) abi.Double `ffi:"log10"`
	Log2  func(abi.Double) abi.Double `ffi:"log2"`
	Log1p func(abi.Double) abi.Double `ffi:"log1p"`

	Pow   func(abi.Double, abi.Double) abi.Double `ffi:"pow"`
	Sqrt  func(abi.Double) abi.Double             `ffi:"sqrt"`
	Cbrt  func(abi.Double) abi.Double             `ffi:"cbrt"`
	Hypot func(abi.Double, abi.Double) abi.Double `ffi:"hypot"`

	Sin   func(abi.Double) abi.Double             `ffi:"sin"`
	Cos   func(abi.Double) abi.Double             `ffi:"cos"`
	Tan   func(abi.Double) abi.Double             `ffi:"tan"`
	Asin  func(abi.Double) abi.Double             `ffi:"asin"`
	Acos  func(abi.Double) abi.Double             `ffi:"acos"`
	Atan  func(abi.Double) abi.Double             `ffi:"atan"`
	Atan2 func(abi.Double, abi.Double) abi.Double `ffi:"atan2"`

	Sinh  func(abi.Double) abi.Double `ffi:"sinh"`
	Cosh  func(abi.Double) abi.Double `ffi:"cosh"`
	Tanh  func(abi.Double) abi.Double `ffi:"tanh"`
	Asinh func(abi.Double) abi.Double `ffi:"asinh"`
	Acosh func(abi.Double) abi.Double `ffi:"acosh"`
	Atanh func(abi.Double) abi.Double `ffi:"atanh"`

	Erf    func(abi.Double) abi.Double `ffi:"erf"`
	Erfc   func(abi.Double) abi.Double `ffi:"erfc"`
	GammaT func(abi.Double) abi.Double `ffi:"tgamma"`
	GammaL func(abi.Double) abi.Double `ffi:"lgamma"`

	Ceil      func(abi.Double) abi.Double   `ffi:"ceil"`
	Floor     func(abi.Double) abi.Double   `ffi:"floor"`
	Trunc     func(abi.Double) abi.Double   `ffi:"trunc"`
	Round     func(abi.Double) abi.Double   `ffi:"round"`
	NearbyInt func(abi.Double) abi.Double   `ffi:"nearbyint"`
	Int       func(abi.Double) abi.Double   `ffi:"rint"`
	Long      func(abi.Double) abi.Long     `ffi:"lrint"`
	LongLong  func(abi.Double) abi.LongLong `ffi:"llrint"`

	Frexp      func(abi.Double) (abi.Double, abi.Int)      `ffi:"frexp"`
	Ldexp      func(abi.Double, abi.Int) abi.Double        `ffi:"ldexp"`
	Modf       func(abi.Double) (abi.Double, abi.Double)   `ffi:"modf"`
	Scale      func(abi.Double, abi.Double) abi.Double     `ffi:"scalbn"`
	ScaleLong  func(abi.Double, abi.Long) abi.Double       `ffi:"scalbln"`
	LogInt     func(abi.Double) abi.Int                    `ffi:"logb"`
	Logb       func(abi.Double) abi.Double                 `ffi:"logb"`
	NextAfter  func(abi.Double, abi.Double) abi.Double     `ffi:"nextafter"`
	NextToward func(abi.Double, abi.DoubleLong) abi.Double `ffi:"nexttoward"`
	CopySign   func(abi.Double, abi.Double) abi.Double     `ffi:"copysign"`
}

var Float struct {
	LibM

	Abs                func(abi.Float) abi.Float                       `ffi:"fabsf"`
	Mod                func(abi.Float, abi.Float) abi.Float            `ffi:"fmodf"`
	Remainder          func(abi.Float, abi.Float) abi.Float            `ffi:"remainderf"`
	RemainderQuotient  func(abi.Float, abi.Float) (abi.Float, abi.Int) `ffi:"remquof"`
	FusedMuliplyAdd    func(abi.Float, abi.Float, abi.Float) abi.Float `ffi:"fmaf"`
	Max                func(abi.Float, abi.Float) abi.Float            `ffi:"fmaxf"`
	Min                func(abi.Float, abi.Float) abi.Float            `ffi:"fminf"`
	PositiveDifference func(abi.Float, abi.Float) abi.Float            `ffi:"fdimf"`
	Nan                func(abi.String) abi.Float                      `ffi:"nanf"`

	Exp   func(abi.Float) abi.Float `ffi:"expf"`
	Exp2  func(abi.Float) abi.Float `ffi:"exp2f"`
	Expm1 func(abi.Float) abi.Float `ffi:"expm1f"`
	Log   func(abi.Float) abi.Float `ffi:"logf"`
	Log10 func(abi.Float) abi.Float `ffi:"log10f"`
	Log2  func(abi.Float) abi.Float `ffi:"log2f"`
	Log1p func(abi.Float) abi.Float `ffi:"log1pf"`

	Pow   func(abi.Float, abi.Float) abi.Float `ffi:"powf"`
	Sqrt  func(abi.Float) abi.Float            `ffi:"sqrtf"`
	Cbrt  func(abi.Float) abi.Float            `ffi:"cbrtf"`
	Hypot func(abi.Float, abi.Float) abi.Float `ffi:"hypotf"`

	Sin   func(abi.Float) abi.Float            `ffi:"sinf"`
	Cos   func(abi.Float) abi.Float            `ffi:"cosf"`
	Tan   func(abi.Float) abi.Float            `ffi:"tanf"`
	Asin  func(abi.Float) abi.Float            `ffi:"asinf"`
	Acos  func(abi.Float) abi.Float            `ffi:"acosf"`
	Atan  func(abi.Float) abi.Float            `ffi:"atanf"`
	Atan2 func(abi.Float, abi.Float) abi.Float `ffi:"atan2f"`

	Sinh  func(abi.Float) abi.Float `ffi:"sinhf"`
	Cosh  func(abi.Float) abi.Float `ffi:"coshf"`
	Tanh  func(abi.Float) abi.Float `ffi:"tanhf"`
	Asinh func(abi.Float) abi.Float `ffi:"asinhf"`
	Acosh func(abi.Float) abi.Float `ffi:"acoshf"`
	Atanh func(abi.Float) abi.Float `ffi:"atanhf"`

	Erf    func(abi.Float) abi.Float `ffi:"erff"`
	Erfc   func(abi.Float) abi.Float `ffi:"erfcf"`
	GammaT func(abi.Float) abi.Float `ffi:"tgammaf"`
	GammaL func(abi.Float) abi.Float `ffi:"lgammaf"`

	Ceil      func(abi.Float) abi.Float    `ffi:"ceilf"`
	Floor     func(abi.Float) abi.Float    `ffi:"floorf"`
	Trunc     func(abi.Float) abi.Float    `ffi:"truncf"`
	Round     func(abi.Float) abi.Float    `ffi:"roundf"`
	NearbyInt func(abi.Float) abi.Float    `ffi:"nearbyintf"`
	Int       func(abi.Float) abi.Float    `ffi:"rintf"`
	Long      func(abi.Float) abi.Long     `ffi:"lrintf"`
	LongLong  func(abi.Float) abi.LongLong `ffi:"llrintf"`

	Frexp      func(abi.Float) (abi.Float, abi.Int)      `ffi:"frexpf"`
	Ldexp      func(abi.Float, abi.Int) abi.Float        `ffi:"ldexpf"`
	Modf       func(abi.Float) (abi.Float, abi.Float)    `ffi:"modff"`
	Scale      func(abi.Float, abi.Float) abi.Float      `ffi:"scalbnf"`
	ScaleLong  func(abi.Float, abi.Long) abi.Float       `ffi:"scalblnf"`
	LogInt     func(abi.Float) abi.Int                   `ffi:"logbf"`
	Logb       func(abi.Float) abi.Float                 `ffi:"logbf"`
	NextAfter  func(abi.Float, abi.Float) abi.Float      `ffi:"nextafterf"`
	NextToward func(abi.Float, abi.DoubleLong) abi.Float `ffi:"nexttowardf"`
	CopySign   func(abi.Float, abi.Float) abi.Float      `ffi:"copysignf"`
}

var DoubleLong struct {
	LibM

	Abs                func(abi.DoubleLong) abi.DoubleLong                                 `ffi:"fabsl"`
	Mod                func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong                 `ffi:"fmodl"`
	Remainder          func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong                 `ffi:"remainderl"`
	RemainderQuotient  func(abi.DoubleLong, abi.DoubleLong) (abi.DoubleLong, abi.Int)      `ffi:"remquol"`
	FusedMuliplyAdd    func(abi.DoubleLong, abi.DoubleLong, abi.DoubleLong) abi.DoubleLong `ffi:"fmal"`
	Max                func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong                 `ffi:"fmaxl"`
	Min                func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong                 `ffi:"fminl"`
	PositiveDifference func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong                 `ffi:"fdiml"`
	Nan                func(abi.String) abi.DoubleLong                                     `ffi:"nanl"`

	Exp   func(abi.DoubleLong) abi.DoubleLong `ffi:"expl"`
	Exp2  func(abi.DoubleLong) abi.DoubleLong `ffi:"exp2l"`
	Expm1 func(abi.DoubleLong) abi.DoubleLong `ffi:"expm1l"`
	Log   func(abi.DoubleLong) abi.DoubleLong `ffi:"logl"`
	Log10 func(abi.DoubleLong) abi.DoubleLong `ffi:"log10l"`
	Log2  func(abi.DoubleLong) abi.DoubleLong `ffi:"log2l"`
	Log1p func(abi.DoubleLong) abi.DoubleLong `ffi:"log1pl"`

	Pow   func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong `ffi:"powl"`
	Sqrt  func(abi.DoubleLong) abi.DoubleLong                 `ffi:"sqrtl"`
	Cbrt  func(abi.DoubleLong) abi.DoubleLong                 `ffi:"cbrtl"`
	Hypot func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong `ffi:"hypotl"`

	Sin   func(abi.DoubleLong) abi.DoubleLong                 `ffi:"sinl"`
	Cos   func(abi.DoubleLong) abi.DoubleLong                 `ffi:"cosl"`
	Tan   func(abi.DoubleLong) abi.DoubleLong                 `ffi:"tanl"`
	Asin  func(abi.DoubleLong) abi.DoubleLong                 `ffi:"asinl"`
	Acos  func(abi.DoubleLong) abi.DoubleLong                 `ffi:"acosl"`
	Atan  func(abi.DoubleLong) abi.DoubleLong                 `ffi:"atanl"`
	Atan2 func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong `ffi:"atan2l"`

	Sinh  func(abi.DoubleLong) abi.DoubleLong `ffi:"sinhl"`
	Cosh  func(abi.DoubleLong) abi.DoubleLong `ffi:"coshl"`
	Tanh  func(abi.DoubleLong) abi.DoubleLong `ffi:"tanhl"`
	Asinh func(abi.DoubleLong) abi.DoubleLong `ffi:"asinhl"`
	Acosh func(abi.DoubleLong) abi.DoubleLong `ffi:"acoshl"`
	Atanh func(abi.DoubleLong) abi.DoubleLong `ffi:"atanhl"`

	Erf    func(abi.DoubleLong) abi.DoubleLong `ffi:"erfl"`
	Erfc   func(abi.DoubleLong) abi.DoubleLong `ffi:"erfcl"`
	GammaT func(abi.DoubleLong) abi.DoubleLong `ffi:"tgammal"`
	GammaL func(abi.DoubleLong) abi.DoubleLong `ffi:"lgammal"`

	Ceil      func(abi.DoubleLong) abi.DoubleLong `ffi:"ceill"`
	Floor     func(abi.DoubleLong) abi.DoubleLong `ffi:"floorl"`
	Trunc     func(abi.DoubleLong) abi.DoubleLong `ffi:"truncl"`
	Round     func(abi.DoubleLong) abi.DoubleLong `ffi:"roundl"`
	NearbyInt func(abi.DoubleLong) abi.DoubleLong `ffi:"nearbyintl"`
	Int       func(abi.DoubleLong) abi.DoubleLong `ffi:"rintl"`
	Long      func(abi.DoubleLong) abi.Long       `ffi:"lrintl"`
	LongLong  func(abi.DoubleLong) abi.LongLong   `ffi:"llrintl"`

	Frexp      func(abi.DoubleLong) (abi.DoubleLong, abi.Int)        `ffi:"frexpl"`
	Ldexp      func(abi.DoubleLong, abi.Int) abi.DoubleLong          `ffi:"ldexpl"`
	Modf       func(abi.DoubleLong) (abi.DoubleLong, abi.DoubleLong) `ffi:"modfl"`
	Scale      func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong   `ffi:"scalbnl"`
	ScaleLong  func(abi.DoubleLong, abi.Long) abi.DoubleLong         `ffi:"scalblnl"`
	LogInt     func(abi.DoubleLong) abi.Int                          `ffi:"logbl"`
	Logb       func(abi.DoubleLong) abi.DoubleLong                   `ffi:"logbl"`
	NextAfter  func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong   `ffi:"nextafterl"`
	NextToward func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong   `ffi:"nexttowardl"`
	CopySign   func(abi.DoubleLong, abi.DoubleLong) abi.DoubleLong   `ffi:"copysignl"`
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

func NaN() abi.Float {
	return abi.Float(math.NaN())
}

func Infinity() abi.Float {
	return abi.Float(math.Inf(1))
}
