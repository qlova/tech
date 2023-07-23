package data

import (
	"reflect"
	"strings"
)

type Sync struct {
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
		pointers: pointers,
	}
}

func Path(sync Sync, ptr any) string {
	return sync.pointers[ptr]
}
