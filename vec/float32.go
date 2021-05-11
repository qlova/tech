//Package vec provides variable sized
package vec

import "qlova.tech/f32"

//Float32 is a variable sized vector of float32 values.
type Float32 []float32

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

//Acos returns the angle whose trigonometric cosine is x.
//The range of values returned by acos is [0,π].
//The result is undefined if |x|>1.
func Acos(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Acos(x[i])
	}
	return
}

//Acosh returns the arc hyperbolic cosine of x; the non-negative inverse of cosh.
//Results are undefined if x<1.
func Acosh(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Acosh(x[i])
	}
	return
}

//Asin returns the angle whose trigonometric sine is x.
//The range of values returned by asin is [−π2,π2]. The result is undefined if |x|>1.
func Asin(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Asin(x[i])
	}
	return
}

//Asinh returns the arc hyperbolic sine of x; the inverse of sinh.
func Asinh(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Asinh(x[i])
	}
	return
}

//Atan returns the angle whose trigonometric arctangent is x.
//Values returned in this case are in the range [−π2,π2].
func Atan(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Atan(x[i])
	}
	return
}

//Atan2 returns the angle whose trigonometric arctangent is y/x
//the signs of y and x are used to determine the quadrant that the angle lies in.
//The values returned by atan in this case are in the range [−π,π].
//Results are undefined if x is zero.
func Atan2(y, x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Atan2(y[i], x[i])
	}
	return
}

//Atanh returns the arc hyperbolic tangent of x; the inverse of tanh.
//Results are undefined if |x|>1.
func Atanh(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Atanh(x[i])
	}
	return
}

//Cos returns the trigonometric cosine of angle.
func Cos(angle Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Cos(angle[i])
	}
	return
}

//Cosh returns the hyperbolic cosine of x. The hyperbolic cosine of x is computed as (e^x+e^−x)/2
func Cosh(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Cosh(x[i])
	}
	return
}

//Degrees converts a quantity, specified in radians into degrees.
func Degrees(radians Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Degrees(radians[i])
	}
	return
}

//Radians converts a quantity, specified in degrees into radians.
func Radians(degrees Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Radians(degrees[i])
	}
	return
}

//Sin returns the trigonometric cosine of angle.
func Sin(angle Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Sin(angle[i])
	}
	return
}

//Sinh returns the hyperbolic cosine of x. The hyperbolic cosine of x is computed as (e^x+e^−x)/2
func Sinh(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Sinh(x[i])
	}
	return
}

//Tan returns the trigonometric cosine of angle.
func Tan(angle Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Tan(angle[i])
	}
	return
}

//Tanh returns the hyperbolic cosine of x. The hyperbolic cosine of x is computed as (e^x+e^−x)/2
func Tanh(x Float32) (out Float32) {
	for i := range out {
		out[i] = f32.Tanh(x[i])
	}
	return
}

//Add adds x to v.
func (v Float32) Add(x Float32) {
	for i := range v {
		v[i] = v[i] + x[i]
	}
	return
}

//Sub subtracts x from v.
func (v Float32) Sub(x Float32) {
	for i := range v {
		v[i] = v[i] - x[i]
	}
	return
}

//Minus returns v - x
func (v Float32) Minus(x float32) (out Float32) {
	for i := range out {
		out[i] = v[i] - x
	}
	return
}

//Plus returns v + x
func (v Float32) Plus(x float32) (out Float32) {
	for i := range out {
		out[i] = v[i] + x
	}
	return
}

//Times returns v * x
func (v Float32) Times(x float32) (out Float32) {
	for i := range out {
		out[i] = v[i] * x
	}
	return
}

//DividedBy returns v / x
func (v Float32) DividedBy(x float32) (out Float32) {
	for i := range out {
		out[i] = v[i] / x
	}
	return
}

//Sub returns a - b
func Sub(a, b Float32) (out Float32) {
	for i := range out {
		out[i] = a[i] - b[i]
	}
	return
}

//Add returns a + b
func Add(a, b Float32) (out Float32) {
	for i := range out {
		out[i] = a[i] + b[i]
	}
	return
}

//Mul returns a * b
func Mul(a, b Float32) (out Float32) {
	for i := range out {
		out[i] = a[i] * b[i]
	}
	return
}

//Div returns a / b
func Div(a, b Float32) (out Float32) {
	for i := range out {
		out[i] = a[i] / b[i]
	}
	return
}

//Distance returns the distance between the two points p0 and p1.
func Distance(a, b Float32) float32 {
	var diffs float32
	for i := range a {
		diffs += (a[i] - b[i]) * (a[i] - b[i])
	}
	return f32.Sqrt(diffs)
}

//Dot returns the dot product of two vectors, x and y.
func Dot(x, y Float32) float32 {
	var product float32
	for i := range x {
		product += (x[i] * y[i])
	}
	return product
}

//Invert returns -x
func Invert(x Float32) (out Float32) {
	for i := range out {
		out[i] = -x[i]
	}
	return
}

//FaceForward orients a vector to point away from a surface as defined by its normal.
func FaceForward(N, I, Nref Float32) Float32 {
	if Dot(Nref, I) > 0 {
		return N
	}
	return Invert(N)
}

//Length returns the length of the vector
func Length(x Float32) float32 {
	var product float32
	for i := range x {
		product += (x[i] * x[i])
	}
	return f32.Sqrt(product)
}

//Normalize normalizes the vector.
func (v Float32) Normalize() {
	var length = Length(v)

	for i := range v {
		v[i] = v[i] / length
	}

	return
}

//Normalize returns a vector with the same direction as its parameter, v, but with length 1.
func Normalize(x Float32) (out Float32) {
	var length = Length(x)

	for i := range out {
		out[i] = out[i] / length
	}

	return
}

//Reflect returns the reflection direction for a given incident vector I and surface normal N.
func Reflect(I, N Float32) Float32 {
	return Sub(I, N.Times(2*Dot(N, I)))
}

//Refract returns the refraction vector for a given incident vector I and surface normal N and ratio of indices of refraction.
func Refract(I, N Float32, eta float32) Float32 {
	k := 1.0 - eta*eta*(1.0-Dot(N, I)*Dot(N, I))
	if k < 0.0 {
		return Float32{}
	}
	return Sub(I.Times(eta), N.Times(eta*Dot(N, I)+f32.Sqrt(k)))
}
