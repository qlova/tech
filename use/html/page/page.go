package page

import (
	"reflect"

	"qlova.tech/use/html/data"
	"qlova.tech/use/html/division"
	"qlova.tech/use/js"
	"qlova.tech/web/tree"
)

// Renderer is any value that
// can render itself as an HTML
// page.
type Renderer interface {
	RenderPage() tree.Node
}

func New(args ...any) tree.Node {
	return division.New(args...)
}

func Goto(page Renderer) js.String {
	return js.String(`page.goto('` + data.PathOf(page) + `');`)
}

func pagesOf(rvalue reflect.Value, components map[reflect.Type]any) {
	if rvalue.Kind() != reflect.Ptr {
		return
	}

	rtype := rvalue.Type().Elem()

	if rvalue.IsNil() {

		if ptr, ok := components[rtype]; ok {
			rvalue.Set(reflect.ValueOf(ptr))
			return
		}

		rvalue.Set(reflect.New(rtype))
		components[rtype] = rvalue.Interface()
	}

	if rtype.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < rtype.NumField(); i++ {
		pagesOf(rvalue.Elem().Field(i), components)
	}
}

// PagesOf crawls an index page and returns
// all of the reachable pages. The pages
// are setup so that they all reference
// each-other.
func PagesOf(index Renderer) []Renderer {
	var components = make(map[reflect.Type]any)

	rvalue := reflect.ValueOf(index)
	rtype := rvalue.Type().Elem()
	components[rtype] = rvalue.Interface()

	pagesOf(rvalue, components)

	var slice = make([]Renderer, 0)
	for _, component := range components {
		data.Register(component)
		if page, ok := component.(Renderer); ok {
			slice = append(slice, page)
		}
	}

	return slice
}
