package main

import (
	"log"

	"qlova.tech/app"
	"qlova.tech/dsl"
	"qlova.tech/gpu"
	"qlova.tech/rgb"
	"qlova.tech/xyz"

	_ "qlova.tech/gpu/opengl/2.1"
	_ "qlova.tech/gpu/opengl/es/2.0"
	_ "qlova.tech/gpu/webgl/1.0"
)

type HelloTriangle struct {
	app.System
	gpu.Program

	color    rgb.Color
	triangle gpu.Mesh
}

func (hello *HelloTriangle) Load() (err error) {
	hello.triangle, err = gpu.NewMesh(gpu.Attributes{
		"position": gpu.Data([]xyz.Vector{
			{-1, -1, 0},
			{1, -1, 0},
			{0, 1, 0},
		}),
	})
	hello.color = rgb.Bytes(255, 255, 0, 255)
	return err
}

func (*HelloTriangle) Vertex(core dsl.Core) {
	core.Set.Vec3(core.Position, core.In.Vec3("position"))
}

func (hello *HelloTriangle) Fragment(core dsl.Core) {
	core.Set.RGB(core.Fragment, core.Get.RGB(&hello.color))
}

func (hello *HelloTriangle) Update() {
	gpu.NewFrame(0)

	hello.Draw(hello.triangle)
}

func main() {
	if err := app.Open("Hello Triangle",
		new(HelloTriangle),
	); err != nil {
		log.Fatal(err)
	}
}
