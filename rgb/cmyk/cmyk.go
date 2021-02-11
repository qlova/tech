package cmyk

import "image/color"

//Model is the color.Model for cmyk
var Model color.Model = color.ModelFunc(func(c color.Color) color.Color {
	return From(c)
})

//Color is an immutable cmyk Color.
type Color struct {
	c, m, y, k uint8
}

//New returns a new cmyk color from the given values.
func New(cyan, magenta, yellow, black uint8) Color {
	return Color{cyan, magenta, yellow, black}
}

//From converts the color.Color to a Color.
func From(c color.Color) Color {
	if cmyk, ok := c.(Color); ok {
		return cmyk
	}

	r, g, b, _ := c.RGBA()
	return New(color.RGBToCMYK(uint8(r>>8), uint8(g>>8), uint8(b>>8)))
}

//Cyan returns the cyan component of the color.
func (c Color) Cyan() uint8 {
	return c.c
}

//Magenta returns the magenta component of the color.
func (c Color) Magenta() uint8 {
	return c.m
}

//Yellow returns the yellow component of the color.
func (c Color) Yellow() uint8 {
	return c.y
}

//Black returns the black component of the color.
func (c Color) Black() uint8 {
	return c.k
}

//RGBA implements color.Color
func (c Color) RGBA() (uint32, uint32, uint32, uint32) {
	w := 0xffff - uint32(c.k)*0x101
	r := (0xffff - uint32(c.c)*0x101) * w / 0xffff
	g := (0xffff - uint32(c.m)*0x101) * w / 0xffff
	b := (0xffff - uint32(c.y)*0x101) * w / 0xffff
	return r, g, b, 0xffff
}
