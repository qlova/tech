package shaders

import (
	"qlova.tech/gpu"
	"qlova.tech/gpu/mat4"
	"qlova.tech/gpu/rgba"
	"qlova.tech/gpu/vec4"

	_ "embed"
)

//go:embed texture.go
var texturedSource string

type TexturedVertex struct {
	Position gpu.Vec4
	UV       gpu.Vec2
}

type TexturedFragment struct {
	UV gpu.Vec2
}

type TexturedMeshRenderer struct {
	gpu.Program

	Vert TexturedVertex
	Frag TexturedFragment

	DrawID int
	Camera gpu.Mat4

	Transform []gpu.Mat4
	Texture   []gpu.Texture
}

func (t *TexturedMeshRenderer) SourceProgram() (string, *gpu.Program) {
	return texturedSource, &t.Program
}

func (r *TexturedMeshRenderer) Render(mesh gpu.Mesh, transform gpu.Mat4, texture gpu.Texture) {
	r.Transform = append(r.Transform, transform)
	r.Texture = append(r.Texture, texture)
	r.Draw(mesh)
}

func (t *TexturedMeshRenderer) Vertex() gpu.Vec4 {
	t.Frag.UV = t.Vert.UV
	return vec4.Mul(t.Vert.Position, mat4.Mul(t.Camera, t.Transform[t.DrawID]))
}

func (t *TexturedMeshRenderer) Fragment() gpu.RGBA {
	var c = rgba.Sample(t.Texture[t.DrawID], t.Frag.UV)
	if c[3] == 0 {
		return rgba.Discard()
	}
	return c
}
