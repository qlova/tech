package xyz

import "math"

//Transform is a 4x4 3D transformation matrix
//in row-column order.
type Transform [4 * 4]float32

//NewTransform returns a new identity Transform.
func NewTransform() Transform {
	return Transform{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func (a Transform) Mul(b Transform) Transform {
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

	var m Transform

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

// At returns a transform that translates
// to the given position.
func At(x, y, z float32) Transform {
	return Transform{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		x, y, z, 1,
	}
}

func Rotate(angle, x, y, z float32) Transform {
	var m Transform

	s := float32(math.Sin(float64(angle)))
	c := float32(math.Cos(float64(angle)))

	m[0] = x*x*(1-c) + c
	m[1] = y*x*(1-c) + z*s
	m[2] = x*z*(1-c) - y*s
	m[3] = 0

	m[4] = x*y*(1-c) - z*s
	m[5] = y*y*(1-c) + c
	m[6] = y*z*(1-c) + x*s
	m[7] = 0

	m[8] = x*z*(1-c) + y*s
	m[9] = y*z*(1-c) - x*s
	m[10] = z*z*(1-c) + c
	m[11] = 0

	m[12] = 0
	m[13] = 0
	m[14] = 0
	m[15] = 1

	return m
}

// LookAt adds a rotation to the transformation matrix that looks
// at the given position.
func LookAt(eye, at, up Vector) Transform {
	const x, y, z = 0, 1, 2

	zaxis := eye.Sub(at).Normalize()
	xaxis := up.Cross(zaxis).Normalize()
	yaxis := zaxis.Cross(xaxis)

	return Transform{
		xaxis[x], xaxis[y], xaxis[z], 0,
		yaxis[x], yaxis[y], yaxis[z], 0,
		zaxis[x], zaxis[y], zaxis[z], 0,
		eye[x], eye[y], eye[z], 1,
	}
}

// MakeFrustum sets this matrix to a projection frustum matrix bounded by the specified planes.
// Returns pointer to this updated matrix.
func makeFrustum(left, right, bottom, top, near, far float32) Transform {
	var m Transform

	m[0] = 2 * near / (right - left)
	m[1] = 0
	m[2] = 0
	m[3] = 0
	m[4] = 0
	m[5] = 2 * near / (top - bottom)
	m[6] = 0
	m[7] = 0
	m[8] = (right + left) / (right - left)
	m[9] = (top + bottom) / (top - bottom)
	m[10] = -(far + near) / (far - near)
	m[11] = -1
	m[12] = 0
	m[13] = 0
	m[14] = -(2 * far * near) / (far - near)
	m[15] = 0

	return m
}

// MakePerspective sets this matrix to a perspective projection matrix
// with the specified vertical field of view in degrees,
// aspect ratio (width/height) and near and far planes.
// Returns pointer to this updated matrix.
func makePerspective(fov, aspect, near, far float32) Transform {
	ymax := near * float32(math.Tan(float64(fov*0.5*math.Pi/180)))
	ymin := -ymax
	xmin := ymin * aspect
	xmax := ymax * aspect
	return makeFrustum(xmin, xmax, ymin, ymax, near, far)
}

//NewProjection returns a new projection transform.
func NewProjection(fov, aspect, near, far float32) Transform {
	return makePerspective(fov, aspect, near, far)
}

func (a Transform) Inverse() Transform {
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
		return Transform{}
	}

	var m Transform

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
