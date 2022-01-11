package vec4

//Member variable indicies.
//ie v[X] = 2
const (
	X = iota
	Y
	Z
	W
)

//Member variable indicies.
//ie v[R] = 2
const (
	R = iota
	G
	B
	A
)

//Float32 is a 4D vector type.
type Float32 [4]float32
