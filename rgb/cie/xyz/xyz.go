package xyz

import (
	"image/color"

	"qlova.tech/rgb/srgb"
)

//Model is the color.Model for xyz
var Model color.Model = color.ModelFunc(func(c color.Color) color.Color {
	return From(c)
})

//Color is an immutable xyz color.
type Color struct {
	x, y, z float32
}

//New returns a Color with the given xyz values.
func New(x, y, z float32) Color {
	return Color{x, y, z}
}

//From converts the color.Color to a Color.
func From(c color.Color) Color {
	f := srgb.From(c).Linear()
	r, g, b := f.Red(), f.Green(), f.Blue()

	x := 0.4124564*r + 0.3575761*g + 0.1804375*b
	y := 0.2126729*r + 0.7151522*g + 0.0721750*b
	z := 0.0193339*r + 0.1191920*g + 0.9503041*b

	return New(x, y, z)
}

//X returns the x component of the color.
func (c Color) X() float32 {
	return c.x
}

//Y returns the y component of the color.
func (c Color) Y() float32 {
	return c.y
}

//Z returns the z component of the color.
func (c Color) Z() float32 {
	return c.z
}

//RGBA implements color.Color
func (c Color) RGBA() (uint32, uint32, uint32, uint32) {

	x, y, z := c.x, c.y, c.z

	r := 3.2404542*x - 1.5371385*y - 0.4985314*z
	g := -0.9692660*x + 1.8760108*y + 0.0415560*z
	b := 0.0556434*x - 0.2040259*y + 1.0572252*z

	return srgb.New(float32(r), float32(g), float32(b)).Linear().RGBA()
}
