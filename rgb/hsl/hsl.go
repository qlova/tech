package hsl

import (
	"image/color"
	"math"

	"qlova.tech/rgb/srgb"
)

//Model is the color.Model for hsl
var Model color.Model = color.ModelFunc(func(c color.Color) color.Color {
	return From(c)
})

//Color is an immutable hsl color.
type Color struct {
	h, s, l float32
}

//New returns a Color with the given hsl values.
func New(hue, saturation, lightness float32) Color {
	return Color{hue, saturation, lightness}
}

//From converts the color.Color to a Color.
func From(c color.Color) Color {
	f := srgb.From(c)
	r, g, b := float64(f.Red()), float64(f.Green()), float64(f.Blue())

	min := math.Min(math.Min(r, g), b)
	max := math.Max(math.Max(r, g), b)

	var s, h, l float64
	l = (max + min) / 2

	if min == max {
		s = 0
		h = 0
	} else {
		if l < 0.5 {
			s = (max - min) / (max + min)
		} else {
			s = (max - min) / (2.0 - max - min)
		}

		if max == r {
			h = (g - b) / (max - min)
		} else if max == g {
			h = 2.0 + (b-r)/(max-min)
		} else {
			h = 4.0 + (r-g)/(max-min)
		}

		h *= 60

		if h < 0 {
			h += 360
		}
	}

	return New(float32(s), float32(h), float32(l))
}

//Hue returns the hue component of the color.
func (c Color) Hue() float32 {
	return c.h
}

//Saturation returns the saturation component of the color.
func (c Color) Saturation() float32 {
	return c.s
}

//Lightness returns the lightness component of the color.
func (c Color) Lightness() float32 {
	return c.l
}

//RGBA implements color.Color
func (c Color) RGBA() (uint32, uint32, uint32, uint32) {
	if c.s == 0 {
		return srgb.New(c.l, c.l, c.l).RGBA()
	}

	var h, s, l = float64(c.h), float64(c.s), float64(c.l)

	var r, g, b float64
	var t1 float64
	var t2 float64
	var tr float64
	var tg float64
	var tb float64

	if c.l < 0.5 {
		t1 = l * (1.0 + s)
	} else {
		t1 = l + s - l*s
	}

	t2 = 2*l - t1
	h /= 360
	tr = h + 1.0/3.0
	tg = h
	tb = h - 1.0/3.0

	if tr < 0 {
		tr++
	}
	if tr > 1 {
		tr--
	}
	if tg < 0 {
		tg++
	}
	if tg > 1 {
		tg--
	}
	if tb < 0 {
		tb++
	}
	if tb > 1 {
		tb--
	}

	// Red
	if 6*tr < 1 {
		r = t2 + (t1-t2)*6*tr
	} else if 2*tr < 1 {
		r = t1
	} else if 3*tr < 2 {
		r = t2 + (t1-t2)*(2.0/3.0-tr)*6
	} else {
		r = t2
	}

	// Green
	if 6*tg < 1 {
		g = t2 + (t1-t2)*6*tg
	} else if 2*tg < 1 {
		g = t1
	} else if 3*tg < 2 {
		g = t2 + (t1-t2)*(2.0/3.0-tg)*6
	} else {
		g = t2
	}

	// Blue
	if 6*tb < 1 {
		b = t2 + (t1-t2)*6*tb
	} else if 2*tb < 1 {
		b = t1
	} else if 3*tb < 2 {
		b = t2 + (t1-t2)*(2.0/3.0-tb)*6
	} else {
		b = t2
	}

	return srgb.New(float32(r), float32(g), float32(b)).RGBA()
}
