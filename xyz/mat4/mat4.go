package mat4

import "qlova.tech/xyz/f32"

//effective members of a Transform
const (
	x = 12
	y = 13
	z = 14
)

type Float32 [4 * 4]float32

func Identity() Float32 {
	return Float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func (a *Float32) Inverse() Float32 {
	n11 := a[0]
	n12 := a[4]
	n13 := a[8]
	n14 := a[12]
	n21 := a[1]
	n22 := a[5]
	n23 := a[9]
	n24 := a[13]
	n31 := a[2]
	n32 := a[6]
	n33 := a[10]
	n34 := a[14]
	n41 := a[3]
	n42 := a[7]
	n43 := a[11]
	n44 := a[15]

	t11 := n23*n34*n42 - n24*n33*n42 + n24*n32*n43 - n22*n34*n43 - n23*n32*n44 + n22*n33*n44
	t12 := n14*n33*n42 - n13*n34*n42 - n14*n32*n43 + n12*n34*n43 + n13*n32*n44 - n12*n33*n44
	t13 := n13*n24*n42 - n14*n23*n42 + n14*n22*n43 - n12*n24*n43 - n13*n22*n44 + n12*n23*n44
	t14 := n14*n23*n32 - n13*n24*n32 - n14*n22*n33 + n12*n24*n33 + n13*n22*n34 - n12*n23*n34

	det := n11*t11 + n21*t12 + n31*t13 + n41*t14

	if det == 0 {
		return Float32{}
	}

	var m Float32

	m[0] = t11
	m[1] = n24*n33*n41 - n23*n34*n41 - n24*n31*n43 + n21*n34*n43 + n23*n31*n44 - n21*n33*n44
	m[2] = n22*n34*n41 - n24*n32*n41 + n24*n31*n42 - n21*n34*n42 - n22*n31*n44 + n21*n32*n44
	m[3] = n23*n32*n41 - n22*n33*n41 - n23*n31*n42 + n21*n33*n42 + n22*n31*n43 - n21*n32*n43
	m[4] = t12
	m[5] = n13*n34*n41 - n14*n33*n41 + n14*n31*n43 - n11*n34*n43 - n13*n31*n44 + n11*n33*n44
	m[6] = n14*n32*n41 - n12*n34*n41 - n14*n31*n42 + n11*n34*n42 + n12*n31*n44 - n11*n32*n44
	m[7] = n12*n33*n41 - n13*n32*n41 + n13*n31*n42 - n11*n33*n42 - n12*n31*n43 + n11*n32*n43
	m[8] = t13
	m[9] = n14*n23*n41 - n13*n24*n41 - n14*n21*n43 + n11*n24*n43 + n13*n21*n44 - n11*n23*n44
	m[10] = n12*n24*n41 - n14*n22*n41 + n14*n21*n42 - n11*n24*n42 - n12*n21*n44 + n11*n22*n44
	m[11] = n13*n22*n41 - n12*n23*n41 - n13*n21*n42 + n11*n23*n42 + n12*n21*n43 - n11*n22*n43
	m[12] = t14
	m[13] = n13*n24*n31 - n14*n23*n31 + n14*n21*n33 - n11*n24*n33 - n13*n21*n34 + n11*n23*n34
	m[14] = n14*n22*n31 - n12*n24*n31 - n14*n21*n32 + n11*n24*n32 + n12*n21*n34 - n11*n22*n34
	m[15] = n12*n23*n31 - n13*n22*n31 + n13*n21*n32 - n11*n23*n32 - n12*n21*n33 + n11*n22*n33

	for i := range m {
		m[i] *= 1.0 / det
	}

	return m
}

func (a *Float32) SetScale(X, Y, Z float32) {
	a[0] = X
	a[5] = Y
	a[10] = Z
}

func (a Float32) Mul(b Float32) Float32 {
	a11 := a[0]
	a12 := a[4]
	a13 := a[8]
	a14 := a[12]
	a21 := a[1]
	a22 := a[5]
	a23 := a[9]
	a24 := a[13]
	a31 := a[2]
	a32 := a[6]
	a33 := a[10]
	a34 := a[14]
	a41 := a[3]
	a42 := a[7]
	a43 := a[11]
	a44 := a[15]

	b11 := b[0]
	b12 := b[4]
	b13 := b[8]
	b14 := b[12]
	b21 := b[1]
	b22 := b[5]
	b23 := b[9]
	b24 := b[13]
	b31 := b[2]
	b32 := b[6]
	b33 := b[10]
	b34 := b[14]
	b41 := b[3]
	b42 := b[7]
	b43 := b[11]
	b44 := b[15]

	var m Float32

	m[0] = a11*b11 + a12*b21 + a13*b31 + a14*b41
	m[4] = a11*b12 + a12*b22 + a13*b32 + a14*b42
	m[8] = a11*b13 + a12*b23 + a13*b33 + a14*b43
	m[12] = a11*b14 + a12*b24 + a13*b34 + a14*b44

	m[1] = a21*b11 + a22*b21 + a23*b31 + a24*b41
	m[5] = a21*b12 + a22*b22 + a23*b32 + a24*b42
	m[9] = a21*b13 + a22*b23 + a23*b33 + a24*b43
	m[13] = a21*b14 + a22*b24 + a23*b34 + a24*b44

	m[2] = a31*b11 + a32*b21 + a33*b31 + a34*b41
	m[6] = a31*b12 + a32*b22 + a33*b32 + a34*b42
	m[10] = a31*b13 + a32*b23 + a33*b33 + a34*b43
	m[14] = a31*b14 + a32*b24 + a33*b34 + a34*b44

	m[3] = a41*b11 + a42*b21 + a43*b31 + a44*b41
	m[7] = a41*b12 + a42*b22 + a43*b32 + a44*b42
	m[11] = a41*b13 + a42*b23 + a43*b33 + a44*b43
	m[15] = a41*b14 + a42*b24 + a43*b34 + a44*b44

	return m
}

//Translate translates the transfrom by the given x,y,z coordinates.
func (t *Float32) Translate(X, Y, Z float32) {
	t[x] += X
	t[y] += Y
	t[z] += Z
}

//SetPosition sets the position of the transform.
func (t *Float32) SetPosition(X, Y, Z float32) {
	t[x] = X
	t[y] = Y
	t[z] = Z
}

//SetRotation sets the rotation component of the matrix.
func (t *Float32) Rotate(X, Y, Z float32) {
	m := Float32{}
	x := X
	y := Y
	z := Z
	a := f32.Cos(x)
	b := f32.Sin(x)
	c := f32.Cos(y)
	d := f32.Sin(y)
	e := f32.Cos(z)
	f := f32.Sin(z)

	ae := a * e
	af := a * f
	be := b * e
	bf := b * f
	m[0] = c * e
	m[4] = -c * f
	m[8] = d
	m[1] = af + be*d
	m[5] = ae - bf*d
	m[9] = -b * c
	m[2] = bf - ae*d
	m[6] = be + af*d
	m[10] = a * c

	// bottom row
	m[3] = 0
	m[7] = 0
	m[11] = 0

	// last column
	m[12] = 0
	m[13] = 0
	m[14] = 0
	m[15] = 1

	*t = t.Mul(m)
}
