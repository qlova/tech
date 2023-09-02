package sdl

import (
	"qlova.tech/abi"
	"qlova.tech/ffi"
)

func Link() error {
	return ffi.Link(
		&Windows,
		&Draw,
		&Timer,
		&System,
	)
}

type Window abi.Pointer

type WindowFlags abi.Uint32

const (
	WindowOpenGL = 0x00000002
	WindowShown  = 0x00000004
)

const (
	WindowCentered = 0x2FFF0000
)

var Windows struct {
	ffi.Header `linux:"libSDL2-2.0.so.0"`

	Error func() string `ffi:"SDL_GetError"`

	Create func(title string, x, y, w, h abi.Int, flags WindowFlags) (Window, error) `ffi:"SDL_CreateWindow"`

	GetSurface    func(Window) (Surface, error) `ffi:"SDL_GetWindowSurface"`
	UpdateSurface func(Window) abi.Error        `ffi:"SDL_UpdateWindowSurface"`
	Destroy       func(Window)                  `ffi:"SDL_DestroyWindow"`
}

type Surface abi.Pointer

type Color abi.Uint32

type Rect struct {
	X, Y, W, H abi.Int
}

var Draw struct {
	ffi.Header `linux:"libSDL2-2.0.so.0"`

	FilledRect func(Surface, *Rect, Color) `ffi:"SDL_FillRect"`
}

var Timer struct {
	ffi.Header `linux:"libSDL2-2.0.so.0"`

	Delay func(ms abi.Uint32) `ffi:"SDL_Delay"`
}

var System struct {
	ffi.Header `linux:"libSDL2-2.0.so.0"`

	Quit func() `ffi:"SDL_Quit"`
}
