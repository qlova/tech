package rgb

// Color is a four-component (red, green, blue, alpha) color.
type Color uint32

// Bytes returns a Color from the given (red, green, blue, alpha) bytes.
func Bytes(r, g, b, a uint8) Color {
	return Color(uint32(r)<<24 | uint32(g)<<16 | uint32(b)<<8 | uint32(a))
}

func (c Color) Red() uint8 {
	return uint8(c >> 24)
}

func (c Color) Green() uint8 {
	return uint8(c >> 16)
}

func (c Color) Blue() uint8 {
	return uint8(c >> 8)
}

func (c Color) Alpha() uint8 {
	return uint8(c)
}

//RGBA implements color.Color
func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.Red())
	r |= r << 8
	r *= uint32(c.Alpha())
	r /= 0xff
	g = uint32(c.Green())
	g |= g << 8
	g *= uint32(c.Alpha())
	g /= 0xff
	b = uint32(c.Blue())
	b |= b << 8
	b *= uint32(c.Alpha())
	b /= 0xff
	a = uint32(c.Alpha())
	a |= a << 8
	return
}

// HTML Color constants.
const (
	AliceBlue            Color = 0xF0F8FF
	AntiqueWhite         Color = 0xFAEBD7
	Aqua                 Color = 0x00FFFF
	Aquamarine           Color = 0x7FFFD4
	Azure                Color = 0xF0FFFF
	Beige                Color = 0xF5F5DC
	Bisque               Color = 0xFFE4C4
	Black                Color = 0x000000
	BlanchedAlmond       Color = 0xFFEBCD
	Blue                 Color = 0x0000FF
	BlueViolet           Color = 0x8A2BE2
	Brown                Color = 0xA52A2A
	BurlyWood            Color = 0xDEB887
	CadetBlue            Color = 0x5F9EA0
	Chartreuse           Color = 0x7FFF00
	Chocolate            Color = 0xD2691E
	Coral                Color = 0xFF7F50
	CornflowerBlue       Color = 0x6495ED
	Cornsilk             Color = 0xFFF8DC
	Crimson              Color = 0xDC143C
	Cyan                 Color = 0x00FFFF
	DarkBlue             Color = 0x00008B
	DarkCyan             Color = 0x008B8B
	DarkGoldenRod        Color = 0xB8860B
	DarkGray             Color = 0xA9A9A9
	DarkGrey             Color = 0xA9A9A9
	DarkGreen            Color = 0x006400
	DarkKhaki            Color = 0xBDB76B
	DarkMagenta          Color = 0x8B008B
	DarkOliveGreen       Color = 0x556B2F
	Darkorange           Color = 0xFF8C00
	DarkOrchid           Color = 0x9932CC
	DarkRed              Color = 0x8B0000
	DarkSalmon           Color = 0xE9967A
	DarkSeaGreen         Color = 0x8FBC8F
	DarkSlateBlue        Color = 0x483D8B
	DarkSlateGray        Color = 0x2F4F4F
	DarkSlateGrey        Color = 0x2F4F4F
	DarkTurquoise        Color = 0x00CED1
	DarkViolet           Color = 0x9400D3
	DeepPink             Color = 0xFF1493
	DeepSkyBlue          Color = 0x00BFFF
	DimGray              Color = 0x696969
	DimGrey              Color = 0x696969
	DodgerBlue           Color = 0x1E90FF
	FireBrick            Color = 0xB22222
	FloralWhite          Color = 0xFFFAF0
	ForestGreen          Color = 0x228B22
	Fuchsia              Color = 0xFF00FF
	Gainsboro            Color = 0xDCDCDC
	GhostWhite           Color = 0xF8F8FF
	Gold                 Color = 0xFFD700
	GoldenRod            Color = 0xDAA520
	Gray                 Color = 0x808080
	Grey                 Color = 0x808080
	Green                Color = 0x008000
	GreenYellow          Color = 0xADFF2F
	HoneyDew             Color = 0xF0FFF0
	HotPink              Color = 0xFF69B4
	IndianRed            Color = 0xCD5C5C
	Indigo               Color = 0x4B0082
	Ivory                Color = 0xFFFFF0
	Khaki                Color = 0xF0E68C
	Lavender             Color = 0xE6E6FA
	LavenderBlush        Color = 0xFFF0F5
	LawnGreen            Color = 0x7CFC00
	LemonChiffon         Color = 0xFFFACD
	LightBlue            Color = 0xADD8E6
	LightCoral           Color = 0xF08080
	LightCyan            Color = 0xE0FFFF
	LightGoldenRodYellow Color = 0xFAFAD2
	LightGray            Color = 0xD3D3D3
	LightGrey            Color = 0xD3D3D3
	LightGreen           Color = 0x90EE90
	LightPink            Color = 0xFFB6C1
	LightSalmon          Color = 0xFFA07A
	LightSeaGreen        Color = 0x20B2AA
	LightSkyBlue         Color = 0x87CEFA
	LightSlateGray       Color = 0x778899
	LightSlateGrey       Color = 0x778899
	LightSteelBlue       Color = 0xB0C4DE
	LightYellow          Color = 0xFFFFE0
	Lime                 Color = 0x00FF00
	LimeGreen            Color = 0x32CD32
	Linen                Color = 0xFAF0E6
	Magenta              Color = 0xFF00FF
	Maroon               Color = 0x800000
	MediumAquaMarine     Color = 0x66CDAA
	MediumBlue           Color = 0x0000CD
	MediumOrchid         Color = 0xBA55D3
	MediumPurple         Color = 0x9370D8
	MediumSeaGreen       Color = 0x3CB371
	MediumSlateBlue      Color = 0x7B68EE
	MediumSpringGreen    Color = 0x00FA9A
	MediumTurquoise      Color = 0x48D1CC
	MediumVioletRed      Color = 0xC71585
	MidnightBlue         Color = 0x191970
	MintCream            Color = 0xF5FFFA
	MistyRose            Color = 0xFFE4E1
	Moccasin             Color = 0xFFE4B5
	NavajoWhite          Color = 0xFFDEAD
	Navy                 Color = 0x000080
	OldLace              Color = 0xFDF5E6
	Olive                Color = 0x808000
	OliveDrab            Color = 0x6B8E23
	Orange               Color = 0xFFA500
	OrangeRed            Color = 0xFF4500
	Orchid               Color = 0xDA70D6
	PaleGoldenRod        Color = 0xEEE8AA
	PaleGreen            Color = 0x98FB98
	PaleTurquoise        Color = 0xAFEEEE
	PaleVioletRed        Color = 0xD87093
	PapayaWhip           Color = 0xFFEFD5
	PeachPuff            Color = 0xFFDAB9
	Peru                 Color = 0xCD853F
	Pink                 Color = 0xFFC0CB
	Plum                 Color = 0xDDA0DD
	PowderBlue           Color = 0xB0E0E6
	Purple               Color = 0x800080
	Red                  Color = 0xFF0000
	RosyBrown            Color = 0xBC8F8F
	RoyalBlue            Color = 0x4169E1
	SaddleBrown          Color = 0x8B4513
	Salmon               Color = 0xFA8072
	SandyBrown           Color = 0xF4A460
	SeaGreen             Color = 0x2E8B57
	SeaShell             Color = 0xFFF5EE
	Sienna               Color = 0xA0522D
	Silver               Color = 0xC0C0C0
	SkyBlue              Color = 0x87CEEB
	SlateBlue            Color = 0x6A5ACD
	SlateGray            Color = 0x708090
	SlateGrey            Color = 0x708090
	Snow                 Color = 0xFFFAFA
	SpringGreen          Color = 0x00FF7F
	SteelBlue            Color = 0x4682B4
	Tan                  Color = 0xD2B48C
	Teal                 Color = 0x008080
	Thistle              Color = 0xD8BFD8
	Tomato               Color = 0xFF6347
	Turquoise            Color = 0x40E0D0
	Violet               Color = 0xEE82EE
	Wheat                Color = 0xF5DEB3
	White                Color = 0xFFFFFF
	WhiteSmoke           Color = 0xF5F5F5
	Yellow               Color = 0xFFFF00
	YellowGreen          Color = 0x9ACD32
)
