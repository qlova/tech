package rgb

import "image/color"

//Model is the color.Model for rgb
var Model color.Model = color.ModelFunc(func(c color.Color) color.Color {
	return From(c)
})

//Color is an immutable rgb color consisting of one byte for each of the red, green and blue channels.
type Color struct {
	r, g, b uint8
}

//New returns a Color with the given rgb values.
func New(r, g, b uint8) Color {
	return Color{r, g, b}
}

//From converts the color.Color to a Color.
func From(c color.Color) Color {
	if rgb, ok := c.(Color); ok {
		return rgb
	}
	r, g, b, a := c.RGBA()
	if a == 0xffff {
		return New(uint8(r>>8), uint8(g>>8), uint8(b>>8))
	}
	if a == 0 {
		return Color{}
	}
	// Since Color.RGBA returns an alpha-premultiplied color, we should have r <= a && g <= a && b <= a.
	r = (r * 0xffff) / a
	g = (g * 0xffff) / a
	b = (b * 0xffff) / a
	return New(uint8(r>>8), uint8(g>>8), uint8(b>>8))
}

//Red returns the red component of the color.
func (c Color) Red() uint8 {
	return c.r
}

//Green returns the green component of the color.
func (c Color) Green() uint8 {
	return c.g
}

//Blue returns the blue component of the color.
func (c Color) Blue() uint8 {
	return c.b
}

//RGBA implements color.Color
func (c Color) RGBA() (r, g, b, a uint32) {
	return color.NRGBA{c.r, c.g, c.b, 255}.RGBA()
}
