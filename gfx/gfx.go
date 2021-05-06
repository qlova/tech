//Package gfx is a package for drawing two-dimensional graphics.
package gfx

import "qlova.tech/gpu"

//Drawable is any type that can draw itself using the GPU.
type Drawable interface {
	Draw(gpu.DrawOptions, *gpu.Transform)
}

//Draw draws a drawable object with the given transform.
func Draw(d Drawable, t *gpu.Transform) {
	d.Draw(0, t)
}
