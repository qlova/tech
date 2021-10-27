package gui_test

import (
	"testing"

	"qlova.tech/gpu"
	"qlova.tech/gui"
	"qlova.tech/win"

	_ "qlova.tech/gpu/driver/opengl46"
	_ "qlova.tech/win/driver/glfw"
)

func TestMain(t *testing.T) {
	for win.Open() && gpu.Open() && gui.Open() {
		gui.Text("Hello World")
	}
}
