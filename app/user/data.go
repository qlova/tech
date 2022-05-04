package user

import (
	"context"
	"fmt"
	"reflect"
	"strings"
)

/*
	Interfaces based on HTTP methods
	can be implemented on a Go value
	to enable them to be handled
	like an HTTP resource.
*/
type (

	// DataWithGet handles a HTTP GET
	// operation. The reciever
	// value will be returned to the
	// caller after this function
	// returns a nil error.
	DataWithGet interface {
		Get(context.Context) error
	}

	// DataWithSearch handles a HTTP SEARCH
	// operation. Semantically, this should
	// be used instead of MethodGet or MethodQuery
	// when multiple 'result' resources are
	// returned. The reciever
	// value will be returned to the
	// caller after this function
	// returns a nil error.
	DataWithSearch interface {
		Search(context.Context) error
	}

	// DataWithPost handles a HTTP POST
	// operation. The patch returned
	// will be applied to the caller's
	// state.
	DataWithPost interface {
		Post(context.Context) error
	}

	// DataWithPut handles a HTTP PUT
	// operation. No data is returned.
	DataWithPut interface {
		Put(context.Context) error
	}

	// DataWithDelete handles a HTTP DELETE
	// operation. No data is returned.
	DataWithDelete interface {
		Delete(context.Context) error
	}

	// DataWithPatch handles a HTTP PATCH
	// operation. The patch must be applied
	// atomically, otherwise an error is returned
	// and the patch is not applied.
	DataWithPatch interface {
		Patch(context.Context, DataPatch) error
	}
)

type DataPatch struct{}

// Data refers to the root application context which can be used
// to refer to and retrieve different data models.
type Data struct {
	paths map[any][]string
	types map[reflect.Type]any
	args  map[any][]string
}

func DataOf(ui InterfaceRenderer) Data {
	data := Data{
		paths: make(map[any][]string),
		types: make(map[reflect.Type]any),
		args:  make(map[any][]string),
	}
	rtype := reflect.TypeOf(ui)
	if rtype.Kind() == reflect.Ptr {
		rtype = rtype.Elem()
	}
	data.types[rtype] = ui
	data.register(ui)
	return data
}

func (data Data) Types() []reflect.Type {
	types := make([]reflect.Type, 0, len(data.types))
	for t := range data.types {
		types = append(types, t)
	}
	return types
}

func (data Data) ArgsOf(v any) []string {
	return data.args[v]
}

func (data Data) Format(format string, args ...any) string {
	for i, arg := range args {
		if path := data.PathOf(arg); path != "" {
			args[i] = path
		}
	}
	return fmt.Sprintf(format, args...)
}

func (data Data) NameOf(value any) string {
	slice := data.paths[value]
	if len(slice) == 0 {
		return ""
	}
	return data.paths[value][len(data.paths[value])-1]
}

func (data Data) PathOf(value any) string {
	return strings.Join(data.paths[value], ".")
}

func (data Data) registerPath(value reflect.Value, path []string) {
	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			data.registerPath(value.Field(i), append(path, value.Type().Field(i).Name))
		}
	}
	key := value.Addr().Interface()
	data.paths[key] = path
}

// register the given value.
func (data Data) register(value any) {
	rtype := reflect.TypeOf(value)
	rvalue := reflect.ValueOf(value)

	if rtype.Kind() != reflect.Ptr {
		return
	}

	for rtype.Kind() == reflect.Ptr && !rvalue.IsNil() {
		rtype = rtype.Elem()
		rvalue = rvalue.Elem()
	}

	data.registerPath(rvalue, []string{rtype.Name()})
}
