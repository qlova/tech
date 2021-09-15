package glfw

import (
	"fmt"
	"sync"

	"github.com/go-gl/glfw/v3.3/glfw"

	"qlova.tech/win"
)

var inited bool
var windows = []*glfw.Window{nil}
var mutex sync.Mutex

func init() {
	win.Driver = Open
}

func Open(w *win.Window) bool {
	mutex.Lock()
	defer mutex.Unlock()

	var err error

	if !inited {
		if err = glfw.Init(); err != nil {
			w.Error = fmt.Errorf("could not initialise glfw: %w", err)
			return false
		}
	}

	var window = windows[w.ID]
	if w.ID == 0 {
		if w.Width == 0 || w.Height == 0 {
			w.Width, w.Height = 640, 480
		}

		window, err = glfw.CreateWindow(w.Width, w.Height, w.Name, nil, nil)
		if err != nil {
			w.Error = fmt.Errorf("could not open window: %w", err)
			return false
		}

		windows = append(windows, window)
		w.ID = len(windows) - 1
	}

	window.MakeContextCurrent()

	glfw.PollEvents()

	if w.Closed || window.ShouldClose() {
		w.Closed = true
		window.Destroy()

		if len(windows) == 0 {
			glfw.Terminate()
			inited = false
		}

		return false
	}

	width, height := window.GetSize()

	w.Width = width
	w.Height = height

	window.SwapBuffers()

	return true
}
