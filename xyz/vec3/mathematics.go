package vec3

import "qlova.tech/xyz/f32"

//Abs returns the absolute value of x.
func Abs(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Abs(x[i])
	}
	return
}

//Ceil returns a value equal to the nearest integer that is greater than or equal to x.
func Ceil(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Ceil(x[i])
	}
	return
}

//Clamp returns the value of x constrained to the range min to max.
func Clamp(x Float32, min, max float32) (out Float32) {
	for i := range out {
		out[i] = f32.Clamp(x[i], min, max)
	}
	return
}

//Exp returns the natural exponentiation of x. i.e., ex.
func Exp(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Exp(x[i])
	}
	return
}

//Exp2 returns the natural exponentiation of x. i.e., ex.
func Exp2(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Exp2(x[i])
	}
	return
}

//Floor returns a value equal to the nearest integer that is less than or equal to x.
func Floor(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Floor(x[i])
	}
	return
}

/*
FMA performs, where possible, a fused multiply-add operation, returning a * b + c.
In use cases where the return value is eventually consumed by a variable declared as precise:

fma() is considered a single operation, whereas the expression a * b + c consumed by a variable declared as precise is considered two operations.

The precision of fma() can differ from the precision of the expression a * b + c.

fma() will be computed with the same precision as any other fma() consumed by a precise variable, giving invariant results for the same input values of a, b and c.

Otherwise, in the absense of precise consumption, there are no special constraints on the number of operations or difference in precision between fma() and the expression a * b + c.
*/
func FMA(a, b, c Float32) (out Float32) {
	for i := range out {
		out[i] = f32.FMA(a[i], b[i], c[i])
	}
	return
}

//Fract returns the fractional part of x.
func Fract(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Fract(x[i])
	}
	return
}

//InverseSqrt returns the inverse of the square root of x. i.e., the value 1/(√x). Results are undefined if x≤0.
func InverseSqrt(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.InverseSqrt(x[i])
	}
	return
}

//Log returns the natural logarithm of x. i.e., the value y which satisfies x=e^y.
func Log(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Log(x[i])
	}
	return
}

//Log2 returns the base 2 logarithm of x. i.e., the value y which satisfies x=2^y. Results are undefined if x≤0.
func Log2(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Log2(x[i])
	}
	return
}

//Max returns the maximum of the two parameters.
func Max(x, y Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Max(x[i], y[i])
	}
	return
}

//Min returns the minimum of the two parameters.
func Min(x, y Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Min(x[i], y[i])
	}
	return
}

//Mix performs a linear interpolation between x and y using a to weight between them.
func Mix(x, y Float32, a float32) (out Float32) {
	for i := range out {
		out[i] = f32.Mix(x[i], y[i], a)
	}
	return
}

//Mod returns the value of x modulo y.
func Mod(x, y Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Mod(x[i], y[i])
	}
	return
}

//Pow returns the value of x raised to the y power. i.e., x^y. Results are undefined if x< or if x=0 and y=0.
func Pow(x, y Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Pow(x[i], y[i])
	}
	return
}

//Round returns a value equal to the nearest integer to x.
//The fraction 0.5 will round in a direction chosen by the implementation, presumably the direction that is fastest.
//This includes the possibility that Round(x) returns the same value as RoundEven(x) for all values of x.
func Round(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Round(x[i])
	}
	return
}

//RoundEven returns a value equal to the nearest integer to x.
//The fractional part of 0.5 will round toward the nearest even integer. For example, both 3.5 and 4.5 will round to 4.0
func RoundEven(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.RoundEven(x[i])
	}
	return
}

//Sign returns -1.0 if x is less than 0.0, 0.0 if x is equal to 0.0, and +1.0 if x is greater than 0.0.
func Sign(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Sign(x[i])
	}
	return
}

//SmoothStep performs smooth Hermite interpolation between 0 and 1 when a < x < b.
//This is useful in cases where a threshold function with a smooth transition is desired.
func SmoothStep(a, b float32, x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.SmoothStep(a, b, x[i])
	}
	return
}

//Sqrt returns the square root of x. i.e., the value √x.
//Results are undefined if x<0.
func Sqrt(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Sqrt(x[i])
	}
	return
}

//Step generates a step function by comparing x to edge.
func Step(edge float32, x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Step(edge, x[i])
	}
	return
}

//Trunc returns a a value equal to the nearest integer to x whose absolute value is not larger than the absolute value of x.
func Trunc(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Trunc(x[i])
	}
	return
}
