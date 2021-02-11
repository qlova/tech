package rgba

import (
	"image/color"

	"qlova.tech/rgb"
)

//Model is the color.Model for rgba
var Model color.Model = color.ModelFunc(func(c color.Color) color.Color {
	return From(c)
})

//Color is an immutable rgb Color with an alpha channel.
type Color struct {
	r, g, b, a uint8
}

//RGBA implements color.Color
func (c Color) RGBA() (r, g, b, a uint32) {
	return color.NRGBA{c.r, c.g, c.b, c.a}.RGBA()
}

//New returns a Color with the given rgba values.
func New(red, green, blue, alpha uint8) Color {
	return Color{red, green, blue, alpha}
}

//With returns a Color with the given rgb value and alpha value.
func With(color rgb.Color, alpha uint8) Color {
	return Color{color.Red(), color.Green(), color.Blue(), alpha}
}

//From converts the color.Color to a Color.
func From(c color.Color) Color {
	if rgb, ok := c.(Color); ok {
		return rgb
	}

	r, g, b, a := c.RGBA()
	if a == 0xffff {
		return New(uint8(r>>8), uint8(g>>8), uint8(b>>8), 0xff)
	}
	if a == 0 {
		return Color{}
	}

	// Since Color.RGBA returns an alpha-premultiplied color, we should have r <= a && g <= a && b <= a.
	r = (r * 0xffff) / a
	g = (g * 0xffff) / a
	b = (b * 0xffff) / a

	return New(uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
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

//Alpha returns the alpha component of the color.
func (c Color) Alpha() uint8 {
	return c.a
}
