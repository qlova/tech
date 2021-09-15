// +build !js

//Package run provides a cross-platform application runner.
package run

import (
	"qlova.tech/gpu"
	"qlova.tech/win"

	_ "qlova.tech/gpu/driver/opengl46"
	_ "qlova.tech/win/driver/glfw"
)

//App runs the given app function in a main loop
//after opening a window and the GPU.
func App(name string, app func()) error {
	win.Name = name

	win.Open()
	if win.Error != nil {
		return win.Error
	}
	if err := gpu.Open(); err != nil {
		return err
	}

	gpu.Set("camera", gpu.NewTransform())

	for win.Open() && gpu.Frames() {
		app()
		if err := gpu.Sync(); err != nil {
			return err
		}
	}
	return nil
}

type Test struct {
	String string
}
