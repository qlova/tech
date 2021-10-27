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
	gpu.Open()

	gpu.Set("camera", gpu.NewTransform())

	for win.Open() && gpu.Open() {
		app()
	}
	return nil
}

type Test struct {
	String string
}
