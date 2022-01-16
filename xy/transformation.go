package xy

import "unsafe"

// Transformer is either a Transform or a Transformation.
type Transformer interface {
	Transform | Transformation

	transformation() Transformation
}

// Transformation is a more expansive Transform
// that can describe any affine 2D transformation.
// Stored as a 3x3 matrix in column-major order.
type Transformation struct {
	matrix [3][3]float32
}

// NewTransformation returns a new identity
// Transformation.
func NewTransformation() Transformation {
	return Transformation{
		matrix: [3][3]float32{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
	}
}

// Transform applies this transform to the vector
// and returns a transformed vector.
func (t *Transformation) Transform(v Vector) Vector {
	return Vector{
		v.X*t.matrix[0][0] + v.Y*t.matrix[0][1] + t.matrix[0][2],
		v.X*t.matrix[1][0] + v.Y*t.matrix[1][1] + t.matrix[1][2],
	}
}

// Array returns the underlying float32 array.
func (t *Transformation) Array() *[9]float32 {
	const assert = unsafe.Sizeof(*t) / unsafe.Sizeof(float32(0))
	return (*[assert]float32)(unsafe.Pointer(t))
}

//transformation implements Transformer.
func (t Transformation) transformation() Transformation { return t }

// Compose composes two transformations together,
// The transformations are applied in order.
func Compose[A, B Transformer](first A, then B) Transformation {
	a := first.transformation().matrix
	b := then.transformation().matrix

	return Transformation{
		matrix: [3][3]float32{
			{ //first column
				a[0][0]*b[0][0] + a[1][0]*b[0][1] + a[2][0]*b[0][2],
				a[0][1]*b[0][0] + a[1][1]*b[0][1] + a[2][1]*b[0][2],
				a[0][2]*b[0][0] + a[1][2]*b[0][1] + a[2][2]*b[0][2],
			},
			{ //second column
				a[0][0]*b[1][0] + a[1][0]*b[1][1] + a[2][0]*b[1][2],
				a[0][1]*b[1][0] + a[1][1]*b[1][1] + a[2][1]*b[1][2],
				a[0][2]*b[1][0] + a[1][2]*b[1][1] + a[2][2]*b[1][2],
			},
			{ //third column
				a[0][0]*b[2][0] + a[1][0]*b[2][1] + a[2][0]*b[2][2],
				a[0][1]*b[2][0] + a[1][1]*b[2][1] + a[2][1]*b[2][2],
				a[0][2]*b[2][0] + a[1][2]*b[2][1] + a[2][2]*b[2][2],
			},
		},
	}
}
