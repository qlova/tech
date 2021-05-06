//Package pos provides simple methods to create a two-dimensional transform.
package pos

import "qlova.tech/gpu"

//At returns a transform to the given (x,y) position
func At(x, y float32) *gpu.Transform {
	var t = gpu.NewTransform()
	t.SetPosition(x, y, 0)
	return &t
}

//WithScale returns a transform to the given (x, y) position
//and uniformly scaled with the given scale.
func WithScale(x, y float32, scale float32) *gpu.Transform {
	var t = gpu.NewTransform()
	t.SetScale(scale, scale, scale)
	t.SetPosition(x, y, 0)
	return &t
}
