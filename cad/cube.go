package cad

type Cube struct {
	Size float64
}

func NewCube(size float64) Cube {
	return Cube{size}
}

func (cube Cube) Vertices() []float32 {
	size := float32(cube.Size)
	return []float32{
		-size, -size, -size,
		-size, -size, size,
		-size, size, size,
		size, size, -size,
		-size, -size, -size,
		-size, size, -size,
		size, -size, size,
		-size, -size, -size,
		size, -size, -size,
		size, size, -size,
		size, -size, -size,
		-size, -size, -size,
		-size, -size, -size,
		-size, size, size,
		-size, size, -size,
		size, -size, size,
		-size, -size, size,
		-size, -size, -size,
		-size, size, size,
		-size, -size, size,
		size, -size, size,
		size, size, size,
		size, -size, -size,
		size, size, -size,
		size, -size, -size,
		size, size, size,
		size, -size, size,
		size, size, size,
		size, size, -size,
		-size, size, -size,
		size, size, size,
		-size, size, -size,
		-size, size, size,
		size, size, size,
		-size, size, size,
		size, -size, size,
	}
}
