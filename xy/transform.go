package xy

import "unsafe"

// Transform represents a 2D rotation, scale and position.
// (applied in that order).
type Transform struct {
	Position Vector
	Scale    float32
	Rotation Rotation
}

// NewTransform returns a new Transform with
// rotation set to zero, scale set to one and
// translation set to zero.
func NewTransform() Transform {
	return Transform{
		Scale:    1,
		Rotation: Rotation{cos: 1},
	}
}

// Transform applies this transform to the vector
// and returns a transformed vector.
func (t Transform) Transform(v Vector) Vector {
	r := &t.Rotation
	return Vector{
		v.X*r.cos - v.Y*r.sin + t.Position.X,
		v.X*r.sin + v.Y*r.cos + t.Position.Y,
	}
}

// transformation returns the transform's Transformation.
func (t Transform) transformation() Transformation {
	return Transformation{
		matrix: [3][3]float32{
			{t.Rotation.cos * t.Scale, t.Rotation.sin * t.Scale, 0},
			{-t.Rotation.sin * t.Scale, t.Rotation.cos * t.Scale, 0},
			{t.Position.X, t.Position.Y, 1},
		},
	}
}

// Array returns the underlying float32 array.
func (t *Transform) Array() *[6]float32 {
	const assert = unsafe.Sizeof(*t) / unsafe.Sizeof(float32(0))
	return (*[assert]float32)(unsafe.Pointer(t))
}
