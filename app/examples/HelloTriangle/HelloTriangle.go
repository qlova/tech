package main

import (
	"log"

	"qlova.tech/app"
	"qlova.tech/dsl"
	"qlova.tech/gpu"
	"qlova.tech/rgb"
	"qlova.tech/xyz/vec3"
	"qlova.tech/xyz/vertex"

	_ "qlova.tech/gpu/opengl/2.1"
)

type HelloTriangle struct {
	app.System
	gpu.Program

	triangle gpu.Mesh
}

func (hello *HelloTriangle) Load() (err error) {
	hello.triangle, err = gpu.NewMesh(vertex.Attributes{
		vertex.Position: vertex.Data([]vec3.Float32{
			{-1, -1, 0},
			{1, -1, 0},
			{0, 1, 0},
		}),
	})
	return err
}

func (*HelloTriangle) Vertex(core dsl.Core) {
	core.Set.Vec4(core.Position, core.In.Vec4(vertex.Position))
}

func (*HelloTriangle) Fragment(core dsl.Core) {
	f := core.Float
	red := core.RGBA(f(1), f(0), f(0), f(1))

	core.Set.RGBA(core.Fragment, red)
}

func (hello *HelloTriangle) Update() {
	gpu.NewFrame(rgb.New(0, 0, 0))

	hello.Draw(hello.triangle)
}

func main() {
	err := app.New("Hello Triangle",
		new(HelloTriangle),
	).Launch()
	if err != nil {
		log.Fatal(err)
	}
}
