package shaders

import (
	"qlova.tech/gpu"
	"qlova.tech/gpu/mat4"
	"qlova.tech/gpu/rgba"
	"qlova.tech/gpu/vec4"
)

type ColouredVertex struct {
	Position gpu.Vec4
}

type ColouredFragment struct{}

type Coloured struct {
	Camera gpu.Mat4

	Transform []gpu.Mat4
	Color     []gpu.RGB
}

func (c Coloured) Vertex(v ColouredVertex, out *ColouredFragment) gpu.Vec4 {
	return vec4.Mul(v.Position, mat4.Mul(c.Camera, c.Transform[0]))
}

func (c Coloured) Fragment(f ColouredFragment) gpu.RGBA {
	return rgba.New(c.Color[0][R], c.Color[0][G], c.Color[0][B], 1)
}
