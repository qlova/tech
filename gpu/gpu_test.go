package gpu_test

import (
	"image"
	"log"
	"os"

	"qlova.tech/gpu"
	"qlova.tech/gpu/internal/shaders"
	"qlova.tech/win"

	_ "image/png"

	_ "qlova.tech/gpu/driver/opengl46"
	_ "qlova.tech/win/driver/glfw"
)

var TextureShader shaders.TexturedMeshRenderer

type Vertex struct {
	Position gpu.Vec3
	UV       gpu.Vec2
}

func main() {
	win.Name = "HelloCube"

	mesh, _ := gpu.NewMesh([]Vertex{
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

	TextureShader.Camera = gpu.NewTransform()

	if err := gpu.Upload(); err != nil {
		log.Fatalln(err)
	}

	var t gpu.Transform

	for win.Open() && gpu.Open() {

		TextureShader.Render(mesh, t, tex)

		if err := gpu.Sync(); err != nil {
			log.Fatalln(err)
		}
	}
}
