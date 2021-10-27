package rgba

import (
	"qlova.tech/gpu"

	_ "embed"
)

//go:embed rgba.go
var Source string

const maxFloat32 = 3.40282e+38

func New(r, g, b, a gpu.Float) gpu.RGBA {
	return gpu.RGBA{r, g, b, a}
}

func Sample(tex gpu.Texture, uv gpu.Vec2) gpu.RGBA {
	//TODO
	return gpu.RGBA{}
}

func Discard() gpu.RGBA {
	return New(0, 0, 0, -maxFloat32)
}
