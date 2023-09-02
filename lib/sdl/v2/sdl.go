package sdl

import (
	"unsafe"

	"qlova.tech/abi"
	"qlova.tech/ffi"
)

func Link() error {
	return ffi.Link(
		&Windows,
		&Draw,
		&Timer,
		&System,
		&Events,
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
	ffi.Header `linux:"libSDL2-2.0.so.0" darwin:"libSDL2.dylib" windows:"SDL2.dll"`

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
	ffi.Header `linux:"libSDL2-2.0.so.0" darwin:"libSDL2.dylib"`

	FilledRect func(Surface, *Rect, Color) `ffi:"SDL_FillRect"`
}

var Timer struct {
	ffi.Header `linux:"libSDL2-2.0.so.0" darwin:"libSDL2.dylib"`

	Delay func(ms abi.Uint32) `ffi:"SDL_Delay"`
}

var System struct {
	ffi.Header `linux:"libSDL2-2.0.so.0" darwin:"libSDL2.dylib"`

	Quit func() `ffi:"SDL_Quit"`
}

type eventType abi.Uint32

const (
	eventQuit eventType = 0x100
)

type Event struct {
	etype eventType
	data  [max(
		unsafe.Sizeof(Quit{}),
	) - unsafe.Sizeof(abi.Uint32(0))]byte
}

func (ev *Event) Data() any {
	switch ev.etype {
	case eventQuit:
		return (*Quit)(unsafe.Pointer(ev))
	default:
		return nil
	}
}

type Quit struct {
	_         eventType
	Timestamp abi.Uint32
}

var Events struct {
	ffi.Header `linux:"libSDL2-2.0.so.0" darwin:"libSDL2.dylib"`

	Poll func(*Event) abi.Int `ffi:"SDL_PollEvent"`
}
