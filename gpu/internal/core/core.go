package core

import "qlova.tech/gpu/texture"

type Pointer [3]uint64

type Texture struct {
	reader  texture.Reader
	pointer Pointer
}

func NewTexture(reader texture.Reader, pointer Pointer) Texture {
	return Texture{reader, pointer}
}
