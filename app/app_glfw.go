//go:build !android && !js

package app

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
)

var window *glfw.Window

func open(name string) error {
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
		glfw.PollEvents()

		if window.ShouldClose() {
			window.Destroy()
			glfw.Terminate()
			return nil
		}

		Width, Height = window.GetSize()

		for _, system := range systems {
			system.Update()
		}

		window.SwapBuffers()
	}
}
