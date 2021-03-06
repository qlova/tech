package rgb

import "qlova.tech/rgb/internal/int"

//Int is an immutable rgb color in integer representation, as in 0xffffff
//invalid values are treated as a zero value. hashtag is optional.
type Int uint32

//Int returns the color as an int.
func (c Color) Int() Int {
	return Int(int.Encode3(c.r, c.g, c.b))
}

//RGBA implements color.Color
func (i Int) RGBA() (uint32, uint32, uint32, uint32) {
	return New(int.Decode3(uint32(i))).RGBA()
}
