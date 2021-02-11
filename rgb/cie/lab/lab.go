package lab

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
	l, a, b float32
}

//New returns a Color with the given lab values.
func New(lightness, green2red, blue2yellow float32) Color {
	return Color{lightness, green2red, blue2yellow}
}

//From converts the color.Color to a Color.
func From(c color.Color) Color {
	from := xyz.From(c)

	x, y, z := from.X(), from.Y(), from.Z()

	f := func(t float32) float32 {
		if t > 6.0/29.0*6.0/29.0*6.0/29.0 {
			return float32(math.Cbrt(float64(t)))
		}
		return t/3.0*29.0/6.0*29.0/6.0 + 4.0/29.0
	}

	fy := f(y / d65[1])
	l := 1.16*fy - 0.16
	a := 5.0 * (f(x/d65[0]) - fy)
	b := 2.0 * (fy - f(z/d65[2]))

	return New(l, a, b)
}

//Lightness returns the lightness component of the color.
func (c Color) Lightness() float32 {
	return c.l
}

//GreenToRed returns the green-to-red component of the color.
func (c Color) GreenToRed() float32 {
	return c.a
}

//BlueToYellow returns the blue-to-yellow component of the color.
func (c Color) BlueToYellow() float32 {
	return c.b
}

//RGBA implements color.Color
func (c Color) RGBA() (uint32, uint32, uint32, uint32) {

	l, a, b := c.l, c.a, c.b

	finv := func(t float32) float32 {
		if t > 6.0/29.0 {
			return t * t * t
		}
		return 3.0 * 6.0 / 29.0 * 6.0 / 29.0 * (t - 4.0/29.0)
	}

	l2 := (l + 0.16) / 1.16
	x := d65[0] * finv(l2+a/5.0)
	y := d65[1] * finv(l2)
	z := d65[2] * finv(l2-b/2.0)

	return xyz.New(x, y, z).RGBA()
}
