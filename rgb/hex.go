package rgb

import "qlova.tech/rgb/internal/hex"

//Hex is an immutable rgb color in hexadecimal representation, as in #ffffff
//invalid values are treated as a zero value. hashtag is optional.
type Hex string

//Hex returns the Color as a Hex.
func (c Color) Hex() Hex {
	return Hex(hex.Encode3ToString(c.r, c.g, c.b))
}

func (h Hex) String() string {
	return hex.String3(string(h))
}

//RGBA implements color.Color
func (h Hex) RGBA() (uint32, uint32, uint32, uint32) {
	return New(hex.Decode3(string(h))).RGBA()
}
