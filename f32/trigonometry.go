//Package f32 provides math
package f32

import "math"

//Acos returns the angle whose trigonometric cosine is x.
//The range of values returned by acos is [0,π].
//The result is undefined if |x|>1.
func Acos(x float32) float32 {
	return float32(math.Acos(float64(x)))
}

//Acosh returns the arc hyperbolic cosine of x; the non-negative inverse of cosh.
//Results are undefined if x<1.
func Acosh(x float32) float32 {
	return float32(math.Acosh(float64(x)))
}

//Asin returns the angle whose trigonometric sine is x.
//The range of values returned by asin is [−π2,π2]. The result is undefined if |x|>1.
func Asin(x float32) float32 {
	return float32(math.Asin(float64(x)))
}

//Asinh returns the arc hyperbolic sine of x; the inverse of sinh.
func Asinh(x float32) float32 {
	return float32(math.Asinh(float64(x)))
}

//Atan returns the angle whose trigonometric arctangent is x.
//Values returned in this case are in the range [−π2,π2].
func Atan(x float32) float32 {
	return float32(math.Atan(float64(x)))
}

//Atan2 returns the angle whose trigonometric arctangent is y/x
//the signs of y and x are used to determine the quadrant that the angle lies in.
//The values returned by atan in this case are in the range [−π,π].
//Results are undefined if x is zero.
func Atan2(y, x float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}

//Atanh returns the arc hyperbolic tangent of x; the inverse of tanh.
//Results are undefined if |x|>1.
func Atanh(x float32) float32 {
	return float32(math.Atanh(float64(x)))
}

//Cos returns the trigonometric cosine of angle.
func Cos(angle float32) float32 {
	return float32(math.Cos(float64(angle)))
}

//Cosh returns the hyperbolic cosine of x.
func Cosh(x float32) float32 {
	return float32(math.Cosh(float64(x)))
}

//Degrees converts a quantity, specified in radians into degrees.
func Degrees(radians float32) float32 {
	return radians * (180.0 / math.Pi)
}

//Radians converts a quantity, specified in degrees into radians.
func Radians(degrees float32) float32 {
	return degrees * (math.Pi / 180)
}

//Sin returns the trigonometric sine of angle.
func Sin(angle float32) float32 {
	return float32(math.Sin(float64(angle)))
}

//Sinh returns the hyperbolic sine of x.
func Sinh(x float32) float32 {
	return float32(math.Sinh(float64(x)))
}

//Tan returns the trigonometric tangent of angle.
func Tan(angle float32) float32 {
	return float32(math.Tan(float64(angle)))
}

//Tanh returns the hyperbolic tangent of x.
func Tanh(x float32) float32 {
	return float32(math.Tanh(float64(x)))
}
