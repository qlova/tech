package vec3

import "qlova.tech/mat/mat4"

//Member variable indicies.
//ie v[X] = 2
const (
	X = iota
	Y
	Z
)

//Member variable indicies.
//ie v[R] = 2
const (
	R = iota
	G
	B
)

//Type is a 3D vector type.
type Type [3]float32

//X returns the x component of the vector.
func (v Type) X() float32 { return v[X] }

//Y returns the y component of the vector.
func (v Type) Y() float32 { return v[Y] }

//Z returns the z component of the vector.
func (v Type) Z() float32 { return v[Z] }

//Transform this vector by the given matrix.
func (v *Type) Transform(m *mat4.Type) {
	x := v[X]
	y := v[Y]
	z := v[Z]

	v[X] = m[0]*x + m[4]*y + m[8]*z + m[12]
	v[Y] = m[1]*x + m[5]*y + m[9]*z + m[13]
	v[Z] = m[2]*x + m[6]*y + m[10]*z + m[14]
}
