package main

import (
	"image"
	"log"
	"os"

	"qlova.tech/gpu"
	"qlova.tech/gpu/shader"
	"qlova.tech/win"

	_ "image/png"

	_ "qlova.tech/gpu/driver/opengl46"
	_ "qlova.tech/win/driver/glfw"
)

type Vertex struct {
	Position gpu.Vec3
	UV       gpu.Vec2
}

func main() {
	win.Name = "HelloCube"
	win.Open()

	if err := gpu.Load(); err != nil {
		log.Fatalln(err)
	}

	mesh1, _ := gpu.NewMesh([]Vertex{
		{gpu.Vec3{0, 0, 0}, gpu.Vec2{0, 0}},
		{gpu.Vec3{-1, 0, 0}, gpu.Vec2{1, 0}},
		{gpu.Vec3{-1, -1, 0}, gpu.Vec2{1, 1}},
	},
		gpu.Indicies{
			0,
			1,
			2,
		},
	)

	f, _ := os.Open("texture.png")
	img, _, _ := image.Decode(f)

	tex, _ := gpu.NewTexture(img)

	mesh1.SetShader(&shader.Textured{
		Texture: tex,
	})

	mesh1.SetPosition(-1, -1, 0)

	mesh2, _ := gpu.NewMesh(gpu.Vertices{
		{0, 0, 0},
		{-1, 0, 0},
		{-1, -1, 0},
	}, gpu.UVs{
		{0, 0},
		{-1, 0},
		{-1, -1},
	})
	mesh2.SetShader(&shader.Colored{
		Color: gpu.Vec3{1, 0, 0},
	})

	var model = gpu.NewModel(mesh1, mesh2)

	var t = gpu.NewTransform()

	t.SetPosition(0.5, 0.5, 0)
	t.SetScale(0.25, 0.25, 0.25)

	gpu.Set("camera", gpu.NewTransform())

	if err := gpu.Upload(); err != nil {
		log.Fatalln(err)
	}

	for win.Open() && gpu.Frames() {

		if err := model.Draw(0, &t); err != nil {
			log.Println(err)
		}

		if err := gpu.Sync(); err != nil {
			log.Fatalln(err)
		}
	}
}
