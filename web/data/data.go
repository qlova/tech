package data

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"qlova.tech/rgb"
	"qlova.tech/use/html"
	"qlova.tech/use/html/attributes"
	"qlova.tech/use/html/division"
	"qlova.tech/web/tree"
)

type specialPath string

const (
	Value  specialPath = "..v"
	Index  specialPath = "..i"
	Offset specialPath = "..o"
)

type Patch struct{}

var global = make(map[reflect.Type]interface{})
var mutex sync.Mutex

func Import[T any](seed tree.Seed) *T {
	var data T

	mutex.Lock()
	defer mutex.Unlock()

	rtype := reflect.TypeOf(data)

	defer func() {
		seed.TreeRenderers[rtype] = global[rtype]
	}()

	if t, ok := global[rtype]; ok {
		return t.(*T)
	}

	global[rtype] = &data
	Register(&data)

	return &data
}

func Use(data any, seed tree.Seed) any {
	mutex.Lock()
	defer mutex.Unlock()

	rtype := reflect.TypeOf(data)
	rvalue := reflect.ValueOf(data)

	for rtype.Kind() == reflect.Ptr && !rvalue.IsNil() {
		rtype = rtype.Elem()
		rvalue = rvalue.Elem()
	}
	iface := rvalue.Addr().Interface()

	defer func() {
		seed.TreeRenderers[rtype] = global[rtype]
	}()

	if t, ok := global[rtype]; ok {
		return t
	}

	global[rtype] = iface
	Register(iface)

	return iface
}

var paths = map[any]string{
	Index: "..i",
	Value: "..v",
}

func register(value reflect.Value, path string) {
	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			register(value.Field(i), path+"."+value.Type().Field(i).Name)
		}
	}
	key := value.Addr().Interface()
	paths[key] = path
}

// Register the value, so that field
// paths can be retrieved (with PathOf)
// for any field in that value. The
// value acts as a reflection mirror.
func Register(value any) {
	rtype := reflect.TypeOf(value)
	rvalue := reflect.ValueOf(value)

	if rtype.Kind() != reflect.Ptr {
		panic("data.Register: value must be a pointer")
	}

	for rtype.Kind() == reflect.Ptr && !rvalue.IsNil() {
		rtype = rtype.Elem()
		rvalue = rvalue.Elem()
	}

	register(rvalue, rtype.Name())
}

func PathOf(value any) string {
	return paths[value]
}

func View(data any) html.Attribute {
	if data == nil {
		return html.Attr("data-view", "..")
	}
	return html.Attr("data-view", PathOf(data))
}

func Scan(data any) html.Attribute {
	return html.Attr("oninput", fmt.Sprintf("data.set('%s', this.value);", PathOf(data)))
}

func Sync(data any) []any {
	return []any{
		View(data),
		Scan(data),
	}
}

func when(condition Condition, args ...any) []any {
	for i, arg := range args {
		switch v := arg.(type) {
		case rgb.Color:
			args[i] = html.Attr("style", fmt.Sprintf("background-color: #%06x;", v))
		}
	}

	for i, arg := range args {
		if renderer, ok := arg.(attributes.Renderer); ok {
			attr := renderer.RenderAttr()
			name, value, ok := strings.Cut(string(attr), "=")

			a := html.Attribute(condition.attr + name)
			if ok {
				a += html.Attribute(fmt.Sprintf("=%v", value))
			}
			args[i] = a
		}
	}
	return args
}

type Condition struct {
	args []string
	attr string
}

func Zero(ptr any) Condition {
	path := PathOf(ptr)
	return Condition{
		args: []string{path},
		attr: fmt.Sprintf("data-when:%v:0:", path),
	}
}

func notZero(ptr any) Condition {
	path := PathOf(ptr)
	return Condition{
		args: []string{path},
		attr: fmt.Sprintf("data-when:%v:", path),
	}
}

func When(ptr any, args ...any) []any {

	cond, ok := ptr.(Condition)
	if !ok {
		cond = notZero(ptr)
	}

	return append(when(cond, args...), html.Attr("data-args", strings.Join(cond.args, " ")))
}

func Echo(format string, args ...any) []any {

	var paths []string
	for _, arg := range args {
		paths = append(paths, PathOf(arg))
	}

	for i := range args {
		format = strings.Replace(format, "%v", "%"+strconv.Itoa(i), 1)
	}

	return []any{
		html.Attr("data-echo", format),
		html.Attr("data-args", strings.Join(paths, " ")),
	}
}

func Feed(ptr any, args ...any) tree.Node {
	return division.New(
		append(args, html.Attr("data-feed", PathOf(ptr)))...,
	)
}
