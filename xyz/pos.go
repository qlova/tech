//Package pos provides simple methods to create a two-dimensional transform.
package xyz

import (
	"qlova.tech/xyz/mat4"
)

//At returns a transform to the given (x,y) position
func At(x, y float32) *mat4.Float32 {
	var t = mat4.Identity()
	t.SetPosition(x, y, 0)
	return &t
}

//WithScale returns a transform to the given (x, y) position
//and uniformly scaled with the given scale.
func WithScale(x, y float32, scale float32) *mat4.Float32 {
	var t = mat4.Identity()
	t.SetScale(scale, scale, 1)
	t.SetPosition(x, y, 0)
	return &t
}

//WithScales returns a transform to the given (x, y) position
//and scaled in the x and y direction.
func WithScales(x, y float32, sx, sy float32) *mat4.Float32 {
	var t = mat4.Identity()
	t.SetScale(sx, sy, 1)
	t.SetPosition(x, y, 0)
	return &t
}
