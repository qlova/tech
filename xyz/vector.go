package xyz

import "math"

//Vector is a 3D vector of floats.
type Vector [3]float32

func (a Vector) Sub(b Vector) Vector {
	return Vector{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func (a Vector) Cross(b Vector) Vector {
	return Vector{a[1]*b[2] - a[2]*b[1], a[2]*b[0] - a[0]*b[2], a[0]*b[1] - a[1]*b[0]}
}

func (a Vector) Dot(b Vector) float32 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

func (a Vector) Normalize() Vector {
	return a.Scale(1 / a.Length())
}

func (a Vector) Scale(b float32) Vector {
	return Vector{a[0] * b, a[1] * b, a[2] * b}
}

func (a Vector) Length() float32 {
	return float32(math.Sqrt(float64(a.Dot(a))))
}
