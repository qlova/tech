package gpu

import (
	"qlova.tech/dsl"
	"qlova.tech/mat/mat4"
	"qlova.tech/rgb/rgba"
	"qlova.tech/vtx"
)

// standardised uniforms/variables.
var (

	// Camera is the transformation of the camera.
	Camera mat4.Float32

	// Transform is the transformation of the current object.
	Transform mat4.Float32
)

// Textured is a basic program that draws a textured mesh.
type Textured struct {
	Program

	Texture Texture
}

// Vertex function.
func (t *Textured) Vertex(core dsl.Core) {
	position := core.In.Vec4(vtx.Position)

	camera := core.Uniform.Mat4(&Camera)
	transform := core.Get.Mat4(&Transform)

	core.Set.Vec4(core.Position, camera.Times(transform).Transform(position))
}

// Fragment function.
func (t *Textured) Fragment(core dsl.Core) {
	f := core.Float

	uv := core.In.Vec2(vtx.UV)
	texture := core.Get.Texture2D(&t.Texture)

	c := core.Var.RGBA(texture.Sample(uv))
	core.If(c.A.LessThan(f(0.85)), func() {
		core.Discard()
	})
	core.Set.RGBA(core.Fragment, c)
}

// Colored is a basic program that draws a colored mesh.
type Colored struct {
	Program

	rgba.Color
}

// Vertex function.
func (c *Colored) Vertex(core dsl.Core) {
	position := core.In.Vec4(vtx.Position)

	camera := core.Uniform.Mat4(&Camera)
	transform := core.Get.Mat4(&Transform)

	core.Set.Vec4(core.Position, camera.Times(transform).Transform(position))
}

// Fragment function.
func (c *Colored) Fragment(core dsl.Core) {
	core.Set.RGBA(core.Fragment, core.Get.RGBA(&c.Color))
}
