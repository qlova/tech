package srgb

import "math"

//Linear is an rgb color (with floating-point values) in linear space.
type Linear struct {
	Color
}

func linearize(v float64) float64 {
	if v <= 0.04045 {
		return v / 12.92
	}
	return math.Pow((v+0.055)/1.055, 2.4)
}

//Linear converts the color to Linear space.
func (c Color) Linear() Linear {
	var l Linear
	l.r = float32(linearize(float64(c.r)))
	l.g = float32(linearize(float64(c.g)))
	l.b = float32(linearize(float64(c.b)))
	return l
}
