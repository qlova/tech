package gpu

import (
	"qlova.tech/gpu/vertex"
	"qlova.tech/mat/mat4"
	"qlova.tech/rgb/rgba"
)

// standardised uniforms/variables.
var (

	// Camera is the transformation of the camera.
	Camera mat4.Type

	// Transform is the transformation of the current object.
	Transform mat4.Type
)

// Textured is a basic program that draws a textured mesh.
type Textured struct {
	Program

	Texture Texture
}

// Vertex function.
func (t *Textured) Vertex(core *Core) {
	position := core.Arg.Vec4(vertex.Position)

	camera := core.Uniform.Mat4(&Camera)
	transform := core.Get.Mat4(&Transform)

	core.Main(func() {
		core.Set.Vec4(core.Position, camera.Times(transform).Transform(position))
	})
}

// Fragment function.
func (t *Textured) Fragment(core *Core) {
	f := core.New.Float

	uv := core.Arg.Vec2(vertex.UV)

	texture := core.Get.Sampler(&t.Texture)

	core.Main(func() {
		c := core.RGBA(core.Sample(texture, uv))
		core.If(c.A.LessThan(f(0.85)), func() {
			core.Discard()
		})
		core.Set.RGBA(core.Fragment, c)
	})
}

// Colored is a basic program that draws a colored mesh.
type Colored struct {
	Program

	rgba.Color
}

// Vertex function.
func (c *Colored) Vertex(core *Core) {
	position := core.Arg.Vec4(vertex.Position)

	camera := core.Uniform.Mat4(&Camera)
	transform := core.Get.Mat4(&Transform)

	core.Main(func() {
		core.Set.Vec4(core.Position, camera.Times(transform).Transform(position))
	})
}

// Fragment function.
func (c *Colored) Fragment(core *Core) {
	color := core.Get.RGBA(&c.Color)

	core.Main(func() {
		core.Set.RGBA(core.Fragment, color)
	})
}
