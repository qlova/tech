//Package bin provides useful types for initialising with go:embed
package bin

import (
	"bytes"
	"image"
	"unsafe"

	"qlova.tech/gpu"

	_ "image/jpeg"
	_ "image/png"
)

var square gpu.Mesh
var textured = &gpu.Textured{}

//Image is image embedded inside of the binary. It will be decoded
//and uploaded to the GPU the first time is is used in a draw call.
type Image []byte

func (img *Image) Draw(ops gpu.DrawOptions, t *gpu.Transform) {

	//TODO generate a pink texture when there is a loading error.

	switch len(*img) {
	case 0:
		return

	case cap(*img):
		//Hacky way to determine whether or not the data has been
		//loaded.
		pix, _, err := image.Decode(bytes.NewReader(*img))
		if err != nil {
			return
		}

		texture, err := gpu.NewTexture(pix)
		if err != nil {
			return
		}

		var abomination = make([]byte, unsafe.Sizeof(texture), unsafe.Sizeof(texture)+1)
		*(*gpu.Texture)(unsafe.Pointer(&abomination[0])) = texture
		*img = abomination

		//Texture is now loaded.

		if square.Nil() {
			square, err = gpu.NewMesh(gpu.Vertices{
				{0, 0, 0},
				{-1, 0, 0},
				{-1, -1, 0},

				{-1, -1, 0},
				{0, -1, 0},
				{0, 0, 0},
			}, gpu.UVs{
				{0, 0},
				{-1, 0},
				{-1, 1},

				{-1, 1},
				{0, 1},
				{0, 0},
			})
			square.SetShader(textured)

			gpu.Upload() //TODO this should happen auto.
		}

	default:
		textured.Texture = *(*gpu.Texture)(unsafe.Pointer(&(*img)[0]))
		square.Draw(ops, t)
	}
}
