package gpu

import (
	"image"

	"qlova.tech/gpu/internal/gputype"
)

//TextureOptions type
type TextureOptions uint64

//TextureOptions
const (
	Nearest = 1 << iota
	NoMipmaps
)

//Texture is an image on the GPU.
type Texture = gputype.Texture

//NewTexture returns a new GPU texture from the given image.Image
func NewTexture(img image.Image) (Texture, error) {
	return context.NewTexture(img)
}

//NewTexture returns a new texture from the given image.
func (context *Context) NewTexture(img image.Image) (Texture, error) {
	if context.Load == nil {
		return gputype.NewTexture(0), ErrNotOpen
	}

	buf, err := context.Load(img)
	if err != nil {
		return gputype.NewTexture(0), err
	}

	return gputype.NewTexture(buf), nil
}
