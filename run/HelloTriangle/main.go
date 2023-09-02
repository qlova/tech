package main

import (
	"qlova.tech/lib/sdl"
)

func init() {
	sdl.Link()
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
	sdl.Timer.Delay(2000)
	sdl.System.Quit()
}
