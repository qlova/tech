package hcl

import (
	"image/color"
	"math"

	"qlova.tech/rgb/cie/lab"
)

//Model is the color.Model for hcl
var Model color.Model = color.ModelFunc(func(c color.Color) color.Color {
	return From(c)
})

//Color is an immutable hcl color.
type Color struct {
	h, c, l float32
}

//New returns a Color with the given hcl values.
func New(hue, chroma, lightness float32) Color {
	return Color{hue, chroma, lightness}
}

//From converts the color.Color to a Color.
func From(col color.Color) Color {
	from := lab.From(col)

	L, a, b := float64(from.Lightness()),
		float64(from.GreenToRed()),
		float64(from.BlueToYellow())

	sq := func(v float64) float64 {
		return v * v
	}

	var h, c, l float64

	// Oops, floating point workaround necessary if a ~= b and both are very small (i.e. almost zero).
	if math.Abs(b-a) > 1e-4 && math.Abs(a) > 1e-4 {
		h = math.Mod(57.29577951308232087721*math.Atan2(b, a)+360.0, 360.0) // Rad2Deg
	} else {
		h = 0.0
	}
	c = math.Sqrt(sq(a) + sq(b))
	l = L

	return New(float32(h), float32(c), float32(l))
}

//Hue returns the hue component of the color.
func (c Color) Hue() float32 {
	return c.h
}

//Chroma returns the chroma component of the color.
func (c Color) Chroma() float32 {
	return c.c
}

//Lightness returns the lightness component of the color.
func (c Color) Lightness() float32 {
	return c.l
}

//RGBA implements color.Color
func (c Color) RGBA() (uint32, uint32, uint32, uint32) {

	h, C, l := c.h, c.c, c.l

	var L, a, b float32

	H := 0.01745329251994329576 * h // Deg2Rad
	a = C * float32(math.Cos(float64(H)))
	b = C * float32(math.Sin(float64(H)))
	L = l

	return lab.New(L, a, b).RGBA()
}
