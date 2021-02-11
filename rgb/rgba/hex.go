package rgba

import "qlova.tech/rgb/internal/hex"

//Hex is an immutable rgba color in hexadecimal representation, as in #ffffff
//invalid values are treated as a zero value. hashtag is optional.
type Hex string

//Hex returns the Color as a Hex.
func (c Color) Hex() Hex {
	return Hex(hex.Encode4ToString(c.r, c.g, c.b, c.a))
}

func (h Hex) String() string {
	return hex.String4(string(h))
}

//RGBA implements color.Color
func (h Hex) RGBA() (uint32, uint32, uint32, uint32) {
	return New(hex.Decode4(string(h))).RGBA()
}
