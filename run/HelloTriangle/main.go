package main

import (
	"runtime"

	"qlova.tech/lib/sdl"
)

func init() {
	runtime.LockOSThread()
	if err := sdl.Link(); err != nil {
		panic(err)
	}
}

func main() {
	window, err := sdl.Windows.Create("Hello Triangle", sdl.WindowCentered,
		sdl.WindowCentered, 640, 480, sdl.WindowOpenGL|sdl.WindowShown)
	if err != nil {
		panic(err)
	}
	defer sdl.Windows.Destroy(window)

	surface, err := sdl.Windows.GetSurface(window)
	if err != nil {
		panic(err)
	}
	sdl.Draw.FilledRect(surface, nil, 0xFFFFFF)
	sdl.Windows.UpdateSurface(window)

	var event sdl.Event
	for ; true; sdl.Events.Poll(&event) {
		if event.Type == sdl.Quit {
			break
		}
	}
	sdl.System.Quit()
}
