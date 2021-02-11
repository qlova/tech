package srgb

import "image/color"

//Model is the color.Model for srgb
var Model color.Model = color.ModelFunc(func(c color.Color) color.Color {
	return From(c)
})

//Color is an immutable rgb color consisting of a floating point value between 0 and 1 for each channel.
type Color struct {
	r, g, b float32
}

//New returns a Color with the given srgb values.
func New(r, g, b float32) Color {
	if r > 1 {
		r = 1
	}
	if g > 1 {
		g = 1
	}
	if b > 1 {
		b = 1
	}
	if r < 0 {
		r = 0
	}
	if g < 0 {
		g = 0
	}
	if b < 0 {
		b = 0
	}
	return Color{r, g, b}
}

//From converts the color.Color to a Color.
func From(c color.Color) Color {
	if rgb, ok := c.(Color); ok {
		return rgb
	}

	r, g, b, a := c.RGBA()
	if a == 0 {
		return Color{}
	}

	r *= 0xffff
	r /= a
	g *= 0xffff
	g /= a
	b *= 0xffff
	b /= a

	return New(float32(r)/65535.0, float32(g)/65535.0, float32(b)/65535.0)
}

//Red returns the red component of the color.
func (c Color) Red() float32 {
	return c.r
}

//Green returns the green component of the color.
func (c Color) Green() float32 {
	return c.g
}

//Blue returns the blue component of the color.
func (c Color) Blue() float32 {
	return c.b
}

//RGBA implements color.Color
func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.r*65535.0 + 0.5)
	g = uint32(c.g*65535.0 + 0.5)
	b = uint32(c.b*65535.0 + 0.5)
	a = 0xFFFF
	return
}
