package main

import (
	"fmt"
	"runtime"

	"qlova.tech/lib/sdl/v2"
)

func init() {
	runtime.LockOSThread()
	if err := sdl.Link(); err != nil {
		panic(err)
	}
}

func main() {
	if err := sdl.System.Init(sdl.Modules); err != 0 {
		panic(sdl.Errors.Get())
	}

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

	fmt.Println(sdl.Audio.Driver())

	var event sdl.Event
	for ; true; sdl.Events.Poll(&event) {
		switch data := event.Data().(type) {
		case *sdl.Quit:
			fmt.Println(data.Timestamp)
			sdl.System.Quit()
			return
		}
	}
}
