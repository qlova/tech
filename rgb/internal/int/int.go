package int

//Encode3 encodes 3 bytes into a uint32
func Encode3(a, b, c byte) (i uint32) {
	return uint32(a)<<16 | uint32(b)<<8 | uint32(c)
}

//Decode3 decodes 3 bytes from a uint32.
func Decode3(i uint32) (a, b, c byte) {
	if i > 0xffffff {
		return 0, 0, 0
	}

	a = byte(i & 0xff0000 >> 16)
	b = byte(i & 0x00ff00 >> 8)
	c = byte(i & 0x0000ff)
	return
}

//Encode4 encodes 4 bytes into a uint32
func Encode4(a, b, c, d byte) (i uint32) {
	return uint32(a)<<24 | uint32(b)<<16 | uint32(c)<<8 | uint32(d)
}

//Decode4 decodes 4 bytes from a uint32.
func Decode4(i uint32) (a, b, c, d byte) {
	a = byte(i & 0xff000000 >> 24)
	b = byte(i & 0x00ff0000 >> 16)
	c = byte(i & 0x0000ff00 >> 8)
	d = byte(i & 0x000000ff)
	return
}
