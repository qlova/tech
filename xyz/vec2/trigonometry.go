package vec2

import "qlova.tech/xyz/f32"

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
