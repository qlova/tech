package xy

import (
	"math"
	"unsafe"
)

// Rotation represents a precomputed 2D rotation.
type Rotation struct {
	cos, sin, angle float32
}

// Angle returns a Rotation corresponding to
// the specified number of radians.
func Angle(radians float32) Rotation {
	return Rotation{
		cos:   float32(math.Cos(float64(radians))),
		sin:   float32(math.Sin(float64(radians))),
		angle: radians,
	}
}

// Degrees returns a Rotation corresponding to
// the specified number of degrees.
func Degrees(degrees float32) Rotation {
	return Angle(degrees * (math.Pi / 180))
}

// Angle returns the angle of the Rotation.
func (r Rotation) Angle() float32 {
	return r.angle
}

// Rotate rotates the specified vector around (0,0)
func (r Rotation) Rotate(v Vector) Vector {
	return Vector{
		v.X*r.cos - v.Y*r.sin,
		v.X*r.sin + v.Y*r.cos,
	}
}

// Array returns the underlying float32 array.
func (r *Rotation) Array() *[3]float32 {
	const assert = unsafe.Sizeof(*r) / unsafe.Sizeof(float32(0))
	return (*[3]float32)(unsafe.Pointer(r))
}
