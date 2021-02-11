package ycbcr

import "image/color"

//Model is the color.Model for ycbcr
var Model color.Model = color.ModelFunc(func(c color.Color) color.Color {
	return From(c)
})

//Color is an immutable ycbcr color.
type Color struct {
	y, cb, cr uint8
}

//New returns a Color with the given ycbcr values.
func New(luma, chromablue, chromared uint8) Color {
	return Color{luma, chromablue, chromared}
}

//From converts the color.Color to a Color.
func From(c color.Color) Color {
	if ycbcr, ok := c.(Color); ok {
		return ycbcr
	}
	r, g, b, _ := c.RGBA()
	return New(color.RGBToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8)))
}

//Luma returns the luma component of the color.
func (c Color) Luma() uint8 {
	return c.y
}

//ChromaBlue returns the blue-difference chroma component of the color.
func (c Color) ChromaBlue() uint8 {
	return c.cb
}

//ChromaRed returns the red-difference chroma component of the color.
func (c Color) ChromaRed() uint8 {
	return c.cr
}

//RGBA implements color.Color
func (c Color) RGBA() (uint32, uint32, uint32, uint32) {
	yy1 := int32(c.y) * 0x10101
	cb1 := int32(c.cb) - 128
	cr1 := int32(c.cr) - 128

	r := yy1 + 91881*cr1
	if uint32(r)&0xff000000 == 0 {
		r >>= 8
	} else {
		r = ^(r >> 31) & 0xffff
	}

	g := yy1 - 22554*cb1 - 46802*cr1
	if uint32(g)&0xff000000 == 0 {
		g >>= 8
	} else {
		g = ^(g >> 31) & 0xffff
	}

	b := yy1 + 116130*cb1
	if uint32(b)&0xff000000 == 0 {
		b >>= 8
	} else {
		b = ^(b >> 31) & 0xffff
	}

	return uint32(r), uint32(g), uint32(b), 0xffff
}
