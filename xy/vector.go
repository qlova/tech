package xy

import (
	"math"
	"unsafe"
)

// Vector is a 2D vector of floats.
type Vector struct {
	X, Y float32
}

// Add returns the sum of the two vectors.
func Add(a, b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y}
}

// Distance returns the distance between the two vectors.
func Distance(a, b Vector) float32 {
	return float32(math.Hypot(float64(a.X-b.X), float64(a.Y-b.Y)))
}

// Lerp returns the linear interpolation of
// 't' between the two vectors.
func Lerp(a, b Vector, t float32) Vector {
	return Vector{
		X: a.X + (b.X-a.X)*t,
		Y: a.Y + (b.Y-a.Y)*t,
	}
}

// Dot returns the dot product of the two vectors.
func Dot(a, b Vector) float32 {
	return a.X*b.X + a.Y*b.Y
}

// Normalize returns the normalized vector
// with distance to origin of 1.
func Normalize(v Vector) Vector {
	l := Distance(Vector{0, 0}, v)
	return Vector{v.X / l, v.Y / l}
}

// Array returns the underlying float32 array.
func (v *Vector) Array() *[2]float32 {
	const assert = unsafe.Sizeof(*v) / unsafe.Sizeof(float32(0))
	return (*[assert]float32)(unsafe.Pointer(v))
}
