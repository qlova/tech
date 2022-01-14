package vec3

import "qlova.tech/xyz/f32"

//Add adds x to v.
func (v *Float32) Add(x Float32) {
	for i := range v {
		v[i] = v[i] + x[i]
	}
	return
}

//Sub subtracts x from v.
func (v *Float32) Sub(x Float32) {
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
func (v *Float32) Normalize() {
	var length = Length(*v)

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
