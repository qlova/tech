package cgo

import (
	"errors"
	"reflect"

	"qlova.tech/cgo/internal/dyncall"
)

// #include <internal/dyncall/dyncall.h>
import "C"

type Header interface {
	header()
}

func Set(header Header, library string) error {
	lib := dlopen(library)
	if lib == nil {
		return errors.New(dlerror())
	}

	rtype := reflect.TypeOf(header).Elem()
	rvalue := reflect.ValueOf(header).Elem()

	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		value := rvalue.Field(i)

		if field.Type.Kind() != reflect.Func {
			continue
		}

		name := field.Tag.Get("cgo")
		if name == "" {
			name = field.Name
		}

		symbol := dlsym(lib, name)
		if symbol == nil {
			return errors.New(dlerror())
		}

		switch fn := value.Addr().Interface().(type) {
		case *func(float64) float64:
			*fn = func(a float64) float64 {
				vm := dyncall.NewVM(8)
				vm.PushFloat64(a)
				return vm.CallFloat64(symbol)
			}
		default:
			value.Set(reflect.MakeFunc(field.Type, func(args []reflect.Value) []reflect.Value {
				var vm = dyncall.NewVM(4096)
				defer vm.Free()
				push := func(value reflect.Value) {
					switch value.Kind() {
					case reflect.Bool:
						vm.PushBool(value.Bool())
					case reflect.Int8:
						vm.PushByte(byte(value.Int()))
					case reflect.Int16:
						vm.PushInt16(int16(value.Int()))
					case reflect.Int32:
						vm.PushInt32(int32(value.Int()))
					case reflect.Int64:
						vm.PushInt64(value.Int())
					case reflect.Float32:
						vm.PushFloat32(float32(value.Float()))
					case reflect.Float64:
						vm.PushFloat64(value.Float())
					case reflect.Pointer:
						vm.PushPointer(value.UnsafePointer())
					default:
						panic("unsupported type " + value.Type().String())
					}
				}
				for _, arg := range args {
					push(arg)
				}
				var results = make([]reflect.Value, field.Type.NumOut())
				for i := 0; i < field.Type.NumOut(); i++ {
					results[i] = reflect.New(field.Type.Out(i)).Elem()
				}
				switch field.Type.NumOut() {
				default:
					if field.Type.NumOut() > 1 {
						for i := 1; i < field.Type.NumOut(); i++ {
							push(results[i].Addr())
						}
					}
					fallthrough
				case 1:
					switch field.Type.Out(0).Kind() {
					case reflect.Bool:
						results[0].SetBool(vm.CallBool(symbol))
					case reflect.Int8:
						results[0].SetInt(int64(vm.CallInt8(symbol)))
					case reflect.Int16:
						results[0].SetInt(int64(vm.CallInt16(symbol)))
					case reflect.Int32:
						results[0].SetInt(int64(vm.CallInt32(symbol)))
					case reflect.Int64:
						results[0].SetInt(int64(vm.CallInt64(symbol)))
					case reflect.Float32:
						results[0].SetFloat(float64(vm.CallFloat32(symbol)))
					case reflect.Float64:
						results[0].SetFloat(float64(vm.CallFloat64(symbol)))
					default:
						panic("unsupported type " + field.Type.Out(0).String())
					}
				}
				return results
			}))
		}
	}

	return nil
}
