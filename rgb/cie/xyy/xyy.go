package xyy

import (
	"image/color"
	"math"

	"qlova.tech/rgb/cie/xyz"
)

var d65 = [...]float32{0.95047, 1.00000, 1.08883}

//Model is the color.Model for xyy
var Model color.Model = color.ModelFunc(func(c color.Color) color.Color {
	return From(c)
})

//Color is an immutable xyy color.
type Color struct {
	x, y, yy float32
}

//New returns a Color with the given xyy values.
func New(x, y, yy float32) Color {
	return Color{x, y, yy}
}

//From converts the color.Color to a Color.
func From(c color.Color) Color {
	from := xyz.From(c)

	X, Y, Z := from.X(), from.Y(), from.Z()

	var x, y, yy float32

	yy = Y
	N := X + Y + Z
	if math.Abs(float64(N)) < 1e-14 {
		// When we have black, Bruce Lindbloom recommends to use
		// the reference white's chromacity for x and y.
		x = d65[0] / (d65[0] + d65[1] + d65[2])
		y = d65[1] / (d65[0] + d65[1] + d65[2])
	} else {
		x = X / N
		y = Y / N
	}

	return New(x, y, yy)
}

//X returns the x component of the color.
func (c Color) X() float32 {
	return c.x
}

//Y returns the two y components of the color.
func (c Color) Y() (y float32, Y float32) {
	return c.y, c.yy
}

//RGBA implements color.Color
func (c Color) RGBA() (uint32, uint32, uint32, uint32) {

	x := c.X()
	y, Y := c.Y()

	var X, Z float32
	Yout := Y

	if -1e-14 < y && y < 1e-14 {
		X = 0.0
		Z = 0.0
	} else {
		X = Y / y * x
		Z = Y / y * (1.0 - x - y)
	}

	return xyz.New(float32(X), float32(Yout), float32(Z)).RGBA()
}
