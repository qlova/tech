package data

import (
	"reflect"
	"strings"
)

type Sync struct {
	parent   *Sync
	base     string
	pointers map[any]string
}

func mirror(value reflect.Value, pointers map[any]string, progress string) {
	switch value.Kind() {
	case reflect.Ptr:
		pointers[value.Interface()] = progress
		mirror(value.Elem(), pointers, progress)
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			if value.Type().Field(i).IsExported() {
				mirror(value.Field(i).Addr(), pointers, progress+"."+value.Type().Field(i).Name)
			}
		}
	}
}

func New(value any) Sync {
	var pointers = make(map[any]string)
	_, name, _ := strings.Cut(reflect.TypeOf(value).String(), ".")
	mirror(reflect.ValueOf(value), pointers, name)
	return Sync{
		base:     name,
		pointers: pointers,
	}
}

func Path(sync Sync, ptr any) string {
	path, ok := sync.pointers[ptr]
	if ok {
		return path
	}
	if sync.parent == nil {
		return ""
	}
	return Path(*sync.parent, ptr)
}

func (s Sync) With(ptr any, indexing any) Sync {

	var pointers = make(map[any]string)
	mirror(reflect.ValueOf(ptr), pointers, Path(s, indexing)+"[*]")
	child := Sync{
		base:     s.base + "[i]",
		pointers: pointers,
	}
	child.parent = &s
	return child
}
