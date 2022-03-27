package page

import (
	"reflect"

	"qlova.tech/use/html"
	"qlova.tech/use/js"
	"qlova.tech/web/data"
	"qlova.tech/web/tree"
)

type renderer struct {
	index tree.Renderer
}

func (r renderer) RenderTree(seed tree.Seed) tree.Node {
	renderer := data.Use(r.index, seed).(tree.Renderer)
	return html.Main(
		renderer.RenderTree(seed)...,
	)
}

func New(args ...any) tree.Renderer {
	var page renderer
	for i, arg := range args {
		if index, ok := arg.(tree.Renderer); ok {
			page.index = index
			args[i] = nil
		}
	}
	return page
}

func Goto(page tree.Renderer) js.String {
	return js.String(`page.goto('` + reflect.TypeOf(page).Elem().Name() + `');`)
}
