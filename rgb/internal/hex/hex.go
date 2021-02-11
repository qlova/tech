package hex

import (
	"encoding/hex"
)

//Encode3ToString encodes 3 bytes to a string
func Encode3ToString(a, b, c byte) string {
	var bytes = [3]byte{a, b, c}
	return hex.EncodeToString(bytes[:])
}

//Encode4ToString encodes 4 bytes to a string
func Encode4ToString(a, b, c, d byte) string {
	var bytes = [4]byte{a, b, c, d}
	return hex.EncodeToString(bytes[:])
}

//Decode3 decodes a 3-byte hex string.
func Decode3(s string) (a, b, c byte) {
	if len(s) == 0 {
		return 0, 0, 0
	}

	if s[0] == '#' {
		s = s[1:]
	}

	if len(s) != 6 {
		return 0, 0, 0
	}

	var bytes [3]byte

	hex.Decode(bytes[:], []byte(s))

	return bytes[0], bytes[1], bytes[2]
}

//Decode4 decodes a 4-byte hex string.
func Decode4(s string) (a, b, c, d byte) {
	if len(s) == 0 {
		return 0, 0, 0, 0
	}

	if s[0] == '#' {
		s = s[1:]
	}

	if len(s) != 8 {
		return 0, 0, 0, 0
	}

	var bytes [4]byte

	hex.Decode(bytes[:], []byte(s))

	return bytes[0], bytes[1], bytes[2], bytes[3]
}

//String3 converts a hex string of 3 bytes with an optional hashtag to a readable string.
func String3(s string) string {
	if len(s) == 0 {
		return "#000000"
	}

	if s[0] == '#' {
		s = s[1:]
	}

	if len(s) != 6 {
		return "#000000"
	}

	var c [3]byte

	_, err := hex.Decode(c[:], []byte(s))
	if err != nil {
		return "#000000"
	}

	return "#" + string(s)
}

//String4 converts a hex string of 4 bytes with an optional hashtag to a readable string.
func String4(s string) string {
	if len(s) == 0 {
		return "#00000000"
	}

	if s[0] == '#' {
		s = s[1:]
	}

	if len(s) != 8 {
		return "#00000000"
	}

	var c [4]byte

	_, err := hex.Decode(c[:], []byte(s))
	if err != nil {
		return "#00000000"
	}

	return "#" + string(s)
}
