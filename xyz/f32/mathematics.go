package f32

import "math"

//Abs returns the absolute value of x.
func Abs(x float32) float32 {
	if x > 0 {
		return x
	}
	return -x
}

//Ceil returns a value equal to the nearest integer that is greater than or equal to x.
func Ceil(x float32) float32 {
	return float32(math.Ceil(float64(x)))
}

//Clamp returns the value of x constrained to the range min to max.
func Clamp(x, min, max float32) float32 {
	return Min(Max(x, min), max)
}

//Exp returns the natural exponentiation of x. i.e., e^x.
func Exp(x float32) float32 {
	return float32(math.Exp(float64(x)))
}

//Exp2 returns 2 raised to the power of x. i.e., 2^x.
func Exp2(x float32) float32 {
	return float32(math.Exp2(float64(x)))
}

//Floor returns a value equal to the nearest integer that is less than or equal to x.
func Floor(x float32) float32 {
	return float32(math.Floor(float64(x)))
}

/*
FMA performs, where possible, a fused multiply-add operation, returning a * b + c.
In use cases where the return value is eventually consumed by a variable declared as precise:

fma() is considered a single operation, whereas the expression a * b + c consumed by a variable declared as precise is considered two operations.

The precision of fma() can differ from the precision of the expression a * b + c.

fma() will be computed with the same precision as any other fma() consumed by a precise variable, giving invariant results for the same input values of a, b and c.

Otherwise, in the absense of precise consumption, there are no special constraints on the number of operations or difference in precision between fma() and the expression a * b + c.
*/
func FMA(a, b, c float32) float32 {
	return a*b + c
}

//Fract returns the fractional part of x.
func Fract(x float32) float32 {
	return x - Floor(x)
}

//InverseSqrt returns the inverse of the square root of x. i.e., the value 1/(√x). Results are undefined if x≤0.
func InverseSqrt(x float32) float32 {
	const magic32 = 0x5F375A86

	// If n is negative return NaN
	if x < 0 {
		return float32(math.NaN())
	}
	// n2 and th are for one iteration of Newton's method later
	n2, th := x*0.5, float32(1.5)
	// Use math.Float32bits to represent the float32, n, as
	// an uint32 without modification.
	b := math.Float32bits(x)
	// Use the new uint32 view of the float32 to shift the bits
	// of the float32 1 to the right, chopping off 1 bit from
	// the fraction part of the float32.
	b = magic32 - (b >> 1)
	// Use math.Float32frombits to convert the uint32 bits back
	// into their float32 representation, again no actual change
	// in the bits, just a change in how we treat them in memory.
	// f is now our answer of 1 / sqrt(n)
	f := math.Float32frombits(b)
	// Perform one iteration of Newton's method on f to improve
	// accuracy
	f *= th - (n2 * f * f)

	// And return our fast inverse square root result
	return f
}

//IsInf returns true if x is posititve or negative floating point infinity and false otherwise.
func IsInf(x float32) bool {
	return math.IsInf(float64(x), 0)
}

//IsNaN returns true if x[i] is posititve or negative floating point NaN (Not a Number) and false otherwise.
func IsNaN(x float32) bool {
	return math.IsNaN(float64(x))
}

//Log returns the natural logarithm of x. i.e., the value y which satisfies x=e^y.
func Log(x float32) float32 {
	return float32(math.Log(float64(x)))
}

//Log2 returns the base 2 logarithm of x. i.e., the value y which satisfies x=2^y. Results are undefined if x≤0.
func Log2(x float32) float32 {
	return float32(math.Log2(float64(x)))
}

//Max returns the maximum of the two parameters.
func Max(x, y float32) float32 {
	if y > x {
		return y
	}
	return x
}

//Min returns the minimum of the two parameters.
func Min(x, y float32) float32 {
	if y < x {
		return y
	}
	return x
}

//Mix performs a linear interpolation between x and y using a to weight between them.
func Mix(x, y, a float32) float32 {
	return x*(1-a) + y*a
}

//Mod returns the value of x modulo y.
func Mod(x, y float32) float32 {
	return x - y*Floor(x/y)
}

//Modf separates a floating point value x into its integer and fractional parts.
func Modf(x float32) (int32, float32) {
	i, f := math.Modf(float64(x))
	return int32(i), float32(f)
}

//Pow returns the value of x raised to the y power. i.e., x^y. Results are undefined if x< or if x=0 and y=0.
func Pow(x, y float32) float32 {
	return float32(math.Pow(float64(x), float64(y)))
}

//Round returns a value equal to the nearest integer to x.
//The fraction 0.5 will round in a direction chosen by the implementation, presumably the direction that is fastest.
//This includes the possibility that Round(x) returns the same value as RoundEven(x) for all values of x.
func Round(x float32) float32 {
	return float32(math.Round(float64(x)))
}

//RoundEven returns a value equal to the nearest integer to x.
//The fractional part of 0.5 will round toward the nearest even integer. For example, both 3.5 and 4.5 will round to 4.0
func RoundEven(x float32) float32 {
	return float32(math.RoundToEven(float64(x)))
}

//Sign returns -1.0 if x is less than 0.0, 0.0 if x is equal to 0.0, and +1.0 if x is greater than 0.0.
func Sign(x float32) float32 {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	}
	return 0
}

//SmoothStep performs smooth Hermite interpolation between 0 and 1 when a < x < b.
//This is useful in cases where a threshold function with a smooth transition is desired.
func SmoothStep(a, b, x float32) float32 {
	t := Clamp((x-a)/(b-a), 0.0, 1.0)
	return t * t * (3.0 - 2.0*t)
}

//Sqrt returns the square root of x. i.e., the value √x.
//Results are undefined if x<0.
func Sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

//Step generates a step function by comparing x to edge.
func Step(edge, x float32) float32 {
	if x < edge {
		return 0
	}
	return 1
}

//Trunc returns a a value equal to the nearest integer to x whose absolute value is not larger than the absolute value of x.
func Trunc(x float32) float32 {
	if x < 0 {
		return Ceil(x)
	}
	return Floor(x)
}
