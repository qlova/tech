package glfw

import (
	"fmt"
	"runtime"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"

	"qlova.tech/win"
)

var MainWindow *glfw.Window

func init() {
	win.CurrentDriver = Driver{}

	runtime.LockOSThread()
}

type Driver struct{}

var previousTime = time.Now()
var lastFrame = time.Now()
var framecount int

func (Driver) Button(id string) bool {
	return Button(id)
}

func Button(id string) bool {
	switch id {
	case "w":
		return MainWindow.GetKey(glfw.KeyW) == glfw.Press
	case "a":
		return MainWindow.GetKey(glfw.KeyA) == glfw.Press
	case "s":
		return MainWindow.GetKey(glfw.KeyS) == glfw.Press
	case "d":
		return MainWindow.GetKey(glfw.KeyD) == glfw.Press
	default:
		return false
	}
}

func (Driver) DeltaTime() float32 {
	return DeltaTime()
}

func DeltaTime() float32 {
	return float32(time.Since(lastFrame).Seconds())
}

func (Driver) Open(name string) error {
	if err := glfw.Init(); err != nil {
		return fmt.Errorf("could not initialise glfw: %w", err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Samples, 4)

	var err error
	MainWindow, err = glfw.CreateWindow(640, 480, name, nil, nil)
	if err != nil {
		return fmt.Errorf("could not open window: %w", err)
	}

	MainWindow.MakeContextCurrent()
	glfw.SwapInterval(1)

	fmt.Println(MainWindow.GetSize())

	return nil
}

func (Driver) Update() bool {

	lastFrame = time.Now()

	MainWindow.SwapBuffers()
	glfw.PollEvents()

	currentTime := time.Now()
	framecount++
	if currentTime.Sub(previousTime) > time.Second {
		fmt.Printf("FPS %v\r", framecount)

		framecount = 0
		previousTime = currentTime
	}

	return !MainWindow.ShouldClose()
}

func (Driver) Close() {
	glfw.Terminate()
}
