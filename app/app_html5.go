//go:build js

package app

import (
	"syscall/js"

	"qlova.tech/gpu"
)

var window = js.Global()
var canvas js.Value

func open(name string) error {
	window.Get("document").Set("title", name)

	//create a full-window canvas.
	canvas = window.Get("document").Call("createElement", "canvas")
	canvas.Set("id", "gpu")
	canvas.Set("width", window.Get("innerWidth"))
	canvas.Set("height", window.Get("innerHeight"))
	window.Get("document").Get("body").Call("appendChild", canvas)

	window.Get("document").Get("body").Get("style").Set("margin", "0")
	window.Get("document").Get("body").Get("style").Set("overflow", "hidden")

	return nil
}

func launch(systems ...System) error {

	var close chan error

	var loop func(this js.Value, args []js.Value) interface{}
	var loopFunc js.Func

	loop = func(this js.Value, args []js.Value) interface{} {
		canvas.Set("width", window.Get("innerWidth"))
		canvas.Set("height", window.Get("innerHeight"))
		Width, Height = window.Get("innerWidth").Int(), window.Get("innerHeight").Int()

		for _, system := range systems {
			system.Update()
		}
		gpu.Sync()

		window.Call("requestAnimationFrame", loopFunc)
		return nil
	}
	loopFunc = js.FuncOf(loop)

	window.Call("requestAnimationFrame", loopFunc)

	return <-close
}
