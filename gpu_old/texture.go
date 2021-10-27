package gpu

import "image"

//TextureOptions type
type TextureOptions uint64

//TextureOptions
const (
	Nearest = 1 << iota
	NoMipmaps
)

//Texture is an image on the GPU.
type Texture struct {
	uint64
	uploadedAtVersion uint64
}

func (t Texture) Value() uint64 {
	return t.uint64
}

//NewTexture returns a new GPU texture from the given image.Image
func NewTexture(img image.Image) (Texture, error) {
	return context.NewTexture(img)
}

//NewTexture returns a new texture from the given image.
func (context *Context) NewTexture(img image.Image) (Texture, error) {
	if context.Load == nil {
		return Texture{0, context.version}, ErrNotOpen
	}

	buf, err := context.Load(img)
	if err != nil {
		return Texture{0, context.version}, err
	}

	return Texture{buf, context.version}, nil
}
