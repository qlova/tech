package css

import (
	"fmt"
	"image/color"
)

// String containing CSS.
type String string

// MediaQuery is a media query.
type MediaQuery string

type Size string

// Color represents a color.
type Color string

// NewColor returns a Color out of the given color.
func NewColor(c color.Color) Color {
	var col = color.NRGBAModel.Convert(c).(color.NRGBA)
	var r, g, b, a = col.R, col.G, col.B, col.A
	if a != 255 {
		return Color(fmt.Sprint("rgba(", r, ",", g, ",", b, ",", float64(a)/255, ")"))
	} else {
		return Color(fmt.Sprint("rgb(", r, ",", g, ",", b, ")"))
	}
}
