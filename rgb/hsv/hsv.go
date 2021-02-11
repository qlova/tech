package hsv

import (
	"image/color"
	"math"

	"qlova.tech/rgb/srgb"
)

//Model is the color.Model for hsv
var Model color.Model = color.ModelFunc(func(c color.Color) color.Color {
	return From(c)
})

//Color is an immutable hsv color.
type Color struct {
	h, s, v float32
}

//New returns a Color with the given hsv values.
func New(hue, saturation, value float32) Color {
	return Color{hue, saturation, value}
}

//From converts the color.Color to a Color.
func From(c color.Color) Color {
	f := srgb.From(c)
	r, g, b := float64(f.Red()), float64(f.Green()), float64(f.Blue())

	var h, s, v float64

	min := math.Min(math.Min(r, g), b)
	v = math.Max(math.Max(r, g), b)
	C := v - min

	s = 0.0
	if v != 0.0 {
		s = C / v
	}

	h = 0.0 // We use 0 instead of undefined as in wp.
	if min != v {
		if v == r {
			h = math.Mod((g-b)/C, 6.0)
		}
		if v == g {
			h = (b-r)/C + 2.0
		}
		if v == b {
			h = (r-g)/C + 4.0
		}
		h *= 60.0
		if h < 0.0 {
			h += 360.0
		}
	}

	return New(float32(h), float32(s), float32(v))
}

//Hue returns the hue component of the color.
func (c Color) Hue() float32 {
	return c.h
}

//Saturation returns the saturation component of the color.
func (c Color) Saturation() float32 {
	return c.s
}

//Value returns the value component of the color.
func (c Color) Value() float32 {
	return c.v
}

//RGBA implements color.Color
func (c Color) RGBA() (uint32, uint32, uint32, uint32) {
	H, S, V := float64(c.h), float64(c.s), float64(c.v)

	Hp := H / 60.0
	C := V * S
	X := C * (1.0 - math.Abs(math.Mod(Hp, 2.0)-1.0))

	m := V - C
	r, g, b := 0.0, 0.0, 0.0

	switch {
	case 0.0 <= Hp && Hp < 1.0:
		r = C
		g = X
	case 1.0 <= Hp && Hp < 2.0:
		r = X
		g = C
	case 2.0 <= Hp && Hp < 3.0:
		g = C
		b = X
	case 3.0 <= Hp && Hp < 4.0:
		g = X
		b = C
	case 4.0 <= Hp && Hp < 5.0:
		r = X
		b = C
	case 5.0 <= Hp && Hp < 6.0:
		r = C
		b = X
	}

	return srgb.New(float32(m+r), float32(m+g), float32(m+b)).RGBA()
}
