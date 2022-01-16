package main

import (
	"log"

	"qlova.tech/app"
	"qlova.tech/dsl"
	"qlova.tech/gpu"
	"qlova.tech/rgb"
	"qlova.tech/xyz"

	_ "qlova.tech/gpu/opengl/2.1"
)

type RotatingCube struct {
	app.System
	gpu.Program

	color rgb.Color
	cube  gpu.Mesh
	angle float32
}

func (rotating *RotatingCube) Load() (err error) {
	rotating.cube, err = gpu.NewMesh(xyz.Cube())
	rotating.color = rgb.Bytes(255, 255, 0, 255)
	return err
}

func (*RotatingCube) Vertex(core dsl.Core) {
	position := core.In.Vec3("position")
	normal := core.In.Vec3("normal")

	camera := core.Uniform.Mat4(&gpu.Camera)
	transform := core.Uniform.Mat4(&gpu.Transform)

	core.Set.Vec3(core.Position, camera.Times(transform).Transform(position))
	core.Set.Vec3(core.Out.Vec3("normal"), transform.TransformNormal(normal))
}

func (rotating *RotatingCube) Fragment(core dsl.Core) {
	normal := core.In.Vec3("normal")

	f := core.Float
	core.Set.RGB(core.Fragment, core.RGB(f(1), f(0), f(0), f(1)))
	core.Set.Vec3(core.Normal, normal)
}

func (rotating *RotatingCube) Update() {
	gpu.Camera = xyz.NewProjection(45, float32(app.Width)/float32(app.Height), 0.1, 100).Mul(
		xyz.LookAt(
			xyz.Vector{0, 2.5, -5},
			xyz.Vector{0, 0, 0},
			xyz.Vector{0, 1, 0},
		).Inverse(),
	)

	gpu.NewFrame(0)

	gpu.Transform = xyz.Rotate(rotating.angle, 0, 1, 0)
	rotating.angle += 0.01

	rotating.Draw(rotating.cube)
}

func main() {
	if err := app.Open("Rotating Cube",
		new(RotatingCube),
	); err != nil {
		log.Fatal(err)
	}
}
