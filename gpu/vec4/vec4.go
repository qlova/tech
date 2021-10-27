package vec4

import (
	_ "embed"

	"qlova.tech/gpu"
)

//go:embed vec4.go
var Source string

func Mul(a gpu.Vec4, b gpu.Mat4) gpu.Vec4 {
	return gpu.Vec4{
		a[0]*b[0] + a[1]*b[4] + a[2]*b[8] + a[3]*b[12],
		a[0]*b[1] + a[1]*b[5] + a[2]*b[9] + a[3]*b[13],
		a[0]*b[2] + a[1]*b[6] + a[2]*b[10] + a[3]*b[14],
		a[0]*b[3] + a[1]*b[7] + a[2]*b[11] + a[3]*b[15],
	}
}
