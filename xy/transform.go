package xy

// Transform is a 3x3 2D transformation matrix
type Transform [3 * 3]float32

// NewTransform returns a new identity Transform.
func NewTransform() Transform {
	return Transform{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}
