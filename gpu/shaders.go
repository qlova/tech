package gpu

import (
	"qlova.tech/gpu/shader"
	"qlova.tech/gpu/vertex"
	"qlova.tech/mat/mat4"
)

type Textured struct {
	Texture   Texture
	Transform mat4.Type
}

func (c *Textured) Variables() []interface{} {
	return []interface{}{&c.Texture}
}

func (t *Textured) Vertex(core *shader.Core) {
	position := core.In.Vec4(vertex.Position)
	uv := core.In.Vec2(vertex.UV)

	camera := core.Uniform.Mat4(&Camera)
	transform := core.Current.Mat4(&t.Transform)

	frag_uv := core.Out.Vec2(vertex.UV)

	core.Main(func() {
		core.Set.Vec4(core.Position, camera.Times(transform).Transform(position))
		core.Set.Vec2(frag_uv, uv)
	})
}

func (t *Textured) Fragment(core *shader.Core) {
	f := core.New.Float

	uv := core.In.Vec2(vertex.UV)
	pixel := core.Out.RGBA("")

	texture := core.Current.Sampler(&t.Texture)

	core.Main(func() {
		c := core.RGBA(core.Sample(texture, uv))
		core.If(c.A.LessThan(f(0.85)), func() {
			core.Discard()
		})
		core.Set.RGBA(pixel, c)
	})
}
