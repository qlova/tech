package opengl2

import (
	"github.com/go-gl/gl/v2.1/gl"
	"qlova.tech/gpu"
)

func init() {
	gpu.Register("opengl2", func() (gpu.Driver, error) {
		if err := open(); err != nil {
			return gpu.Driver{}, err
		}

		return gpu.Driver{}, nil
	})
}

func open() error {
	if err := gl.Init(); err != nil {
		return err
	}

	return nil
}
