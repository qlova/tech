//go:build !android && !js

package app

import (
	"fmt"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	"qlova.tech/gpu"
)

var window *glfw.Window

func open(name string) error {
	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {
		return fmt.Errorf("could not initialise glfw: %w", err)
	}

	window, err = glfw.CreateWindow(Width, Height, name, nil, nil)
	if err != nil {
		return fmt.Errorf("could not open window: %w", err)
	}

	window.MakeContextCurrent()

	return nil
}

func launch(systems ...System) error {
	for {
		Width, Height = window.GetSize()
		glfw.PollEvents()

		if window.ShouldClose() {
			window.Destroy()
			glfw.Terminate()
			return nil
		}

		for _, system := range systems {
			system.Update()
		}

		gpu.Sync()
		window.SwapBuffers()
		runtime.GC()
	}
}
