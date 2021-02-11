package luv

import (
	"image/color"
	"math"

	"qlova.tech/rgb/cie/xyz"
)

var d65 = [...]float32{0.95047, 1.00000, 1.08883}

//Model is the color.Model for lab
var Model color.Model = color.ModelFunc(func(c color.Color) color.Color {
	return From(c)
})

//Color is an immutable lab color.
type Color struct {
	l, u, v float32
}

//New returns a Color with the given luv values.
func New(lightness, u, v float32) Color {
	return Color{lightness, u, v}
}

//From converts the color.Color to a Color.
func From(c color.Color) Color {
	from := xyz.From(c)

	x, y, z := from.X(), from.Y(), from.Z()

	var l, u, v float32

	f := func(x, y, z float32) (u, v float32) {
		denom := x + 15.0*y + 3.0*z
		if denom == 0.0 {
			u, v = 0.0, 0.0
		} else {
			u = 4.0 * x / denom
			v = 9.0 * y / denom
		}
		return
	}

	if y/d65[1] <= 6.0/29.0*6.0/29.0*6.0/29.0 {
		l = y / d65[1] * 29.0 / 3.0 * 29.0 / 3.0 * 29.0 / 3.0
	} else {
		l = 1.16*float32(math.Cbrt(float64(y/d65[1]))) - 0.16
	}
	ubis, vbis := f(x, y, z)
	un, vn := f(d65[0], d65[1], d65[2])
	u = 13.0 * l * (ubis - un)
	v = 13.0 * l * (vbis - vn)

	return New(l, u, v)
}

//Lightness returns the lightness component of the color.
func (c Color) Lightness() float32 {
	return c.l
}

//U returns the u component of the color.
func (c Color) U() float32 {
	return c.u
}

//V returns the v component of the color.
func (c Color) V() float32 {
	return c.v
}

//RGBA implements color.Color
func (c Color) RGBA() (uint32, uint32, uint32, uint32) {

	l, u, v := c.l, c.u, c.v

	var x, y, z float32

	f := func(x, y, z float32) (u, v float32) {
		denom := x + 15.0*y + 3.0*z
		if denom == 0.0 {
			u, v = 0.0, 0.0
		} else {
			u = 4.0 * x / denom
			v = 9.0 * y / denom
		}
		return
	}

	cube := func(v float32) float32 {
		return v * v * v
	}

	if l <= 0.08 {
		y = d65[1] * l * 100.0 * 3.0 / 29.0 * 3.0 / 29.0 * 3.0 / 29.0
	} else {
		y = d65[1] * cube((l+0.16)/1.16)
	}
	un, vn := f(d65[0], d65[1], d65[2])
	if l != 0.0 {
		ubis := u/(13.0*l) + un
		vbis := v/(13.0*l) + vn
		x = y * 9.0 * ubis / (4.0 * vbis)
		z = y * (12.0 - 3.0*ubis - 20.0*vbis) / (4.0 * vbis)
	} else {
		x, y = 0.0, 0.0
	}

	return xyz.New(x, y, z).RGBA()
}
