package cmyk

import "qlova.tech/rgb/internal/int"

//Int is an immutable cmyk color in integer representation, as in 0xffffffff
//invalid values are treated as a zero value. hashtag is optional.
type Int uint32

//Int returns the color as an int.
func (c Color) Int() Int {
	return Int(int.Encode4(c.c, c.m, c.y, c.k))
}

//RGBA implements color.Color
func (i Int) RGBA() (uint32, uint32, uint32, uint32) {
	return New(int.Decode4(uint32(i))).RGBA()
}
