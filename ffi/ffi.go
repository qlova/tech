// Package ffi provides an interface for loading C shared libraries dynamically.
package ffi

import (
	"errors"
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"qlova.tech/abi"
	"qlova.tech/ffi/internal/dyncall"
)

// #include <math.h>
// #include <internal/dyncall/dyncall.h>
import "C"

func Sqrt(a float64) float64 {
	return float64(C.sqrt(C.double(a)))
}

// Library can be embedded inside of a struct to
// mark it as a library interface structure. Each
// other field in the struct must be a func.
type Library interface {
	library()
}

var vm4096 sync.Pool
var vm8 sync.Pool

func init() {
	vm8.New = func() any {
		return dyncall.NewVM(8)
	}
	vm4096.New = func() any {
		return dyncall.NewVM(4096)
	}
}

// Link dynamically links the given libraries, based on
// the platform struct tags of the embedded [Library] field.
func Link(libraries ...Library) error {
	for _, library := range libraries {
		var header = reflect.TypeOf(library).Elem().Field(0)
		for header.Type.Kind() == reflect.Struct {
			header = header.Type.Field(0)
		}
		if err := Set(library, header.Tag.Get(runtime.GOOS)); err != nil {
			return err
		}
	}
	return nil
}

func sigRune(t reflect.Type) rune {
	switch t.Kind() {
	case reflect.TypeOf(abi.Bool(false)).Kind():
		return dyncall.Bool
	case reflect.TypeOf(abi.Char(0)).Kind():
		return dyncall.Char
	case reflect.TypeOf(abi.CharUnsigned(0)).Kind():
		return dyncall.UnsignedChar
	case reflect.TypeOf(abi.Short(0)).Kind():
		return dyncall.Short
	case reflect.TypeOf(abi.ShortUnsigned(0)).Kind():
		return dyncall.UnsignedShort
	case reflect.TypeOf(abi.Int(0)).Kind():
		return dyncall.Int
	case reflect.TypeOf(abi.IntUnsigned(0)).Kind():
		return dyncall.Uint
	case reflect.TypeOf(abi.Long(0)).Kind():
		return dyncall.Long
	case reflect.TypeOf(abi.LongUnsigned(0)).Kind():
		return dyncall.UnsignedLong
	case reflect.TypeOf(abi.LongLong(0)).Kind():
		return dyncall.LongLong
	case reflect.TypeOf(abi.LongLongUnsigned(0)).Kind():
		return dyncall.UnsignedLongLong
	case reflect.TypeOf(abi.Float(0)).Kind():
		return dyncall.Float
	case reflect.TypeOf(abi.Double(0)).Kind():
		return dyncall.Double
	case reflect.String:
		return dyncall.String
	case reflect.Pointer:
		return dyncall.Pointer
	case reflect.Struct:
		if t.Implements(reflect.TypeOf([0]abi.IsPointer{}).Elem()) {
			return dyncall.Pointer
		} else {
			panic("unsupported struct " + t.String())
		}
	default:
		panic("unsupported type " + t.String())
	}
}

func newSignature(ftype reflect.Type) dyncall.Signature {
	var sig dyncall.Signature
	for i := 0; i < ftype.NumIn(); i++ {
		sig.Args = append(sig.Args, sigRune(ftype.In(i)))
	}
	if ftype.NumOut() > 1 {
		sig.Returns = sigRune(ftype.Out(0))
	} else {
		sig.Returns = dyncall.Void
	}
	return sig
}

func newCallback(signature dyncall.Signature, function reflect.Value) dyncall.CallbackHandler {
	return func(cb *dyncall.Callback, args *dyncall.Args, result unsafe.Pointer) rune {
		var values = make([]reflect.Value, len(signature.Args))
		for i := range values {
			values[i] = reflect.New(function.Type().In(i)).Elem()
		}
		for i := 0; i < len(signature.Args); i++ {
			switch signature.Args[i] {
			case dyncall.Bool:
				switch args.Bool() {
				case 0:
					values[i].SetBool(false)
				case 1:
					values[i].SetBool(true)
				}
			case dyncall.Char:
				values[i].SetInt(int64(args.Char()))
			case dyncall.UnsignedChar:
				values[i].SetUint(uint64(args.UnsignedChar()))
			case dyncall.Short:
				values[i].SetInt(int64(args.Short()))
			case dyncall.UnsignedShort:
				values[i].SetUint(uint64(args.UnsignedShort()))
			case dyncall.Int:
				values[i].SetInt(int64(args.Int()))
			case dyncall.Uint:
				values[i].SetUint(uint64(args.UnsignedInt()))
			case dyncall.Long:
				values[i].SetInt(int64(args.Long()))
			case dyncall.UnsignedLong:
				values[i].SetUint(uint64(args.UnsignedLong()))
			case dyncall.LongLong:
				values[i].SetInt(int64(args.LongLong()))
			case dyncall.UnsignedLongLong:
				values[i].SetUint(uint64(args.UnsignedLongLong()))
			case dyncall.Float:
				values[i].SetFloat(float64(args.Float()))
			case dyncall.Double:
				values[i].SetFloat(float64(args.Double()))
			case dyncall.String:
				ptr := args.Pointer()
				switch values[i].Kind() {
				case reflect.String:
					values[i].SetString(C.GoString((*C.char)(ptr)))
				case reflect.Struct:
					values[i].Set(reflect.ValueOf(*(*abi.String)(unsafe.Pointer(&ptr))))
				default:
					panic("unsupported type " + values[i].Type().String())
				}
			case dyncall.Pointer:
				switch values[i].Kind() {
				case reflect.UnsafePointer:
					values[i].SetPointer(unsafe.Pointer(args.Pointer()))
				default:
					settable, ok := values[i].Addr().Interface().(interface {
						SetPointer(unsafe.Pointer)
					})
					if !ok {
						panic("unsupported type " + values[i].Type().String())
					}
					settable.SetPointer(unsafe.Pointer(args.Pointer()))
				}
			default:
				panic("unsupported type " + string(signature.Args[i]))
			}
		}
		results := function.Call(values)
		switch signature.Returns {
		case dyncall.Void:
		case dyncall.Bool:
			*(*abi.Bool)(result) = abi.Bool(results[0].Bool())
		case dyncall.Char:
			*(*abi.Char)(result) = abi.Char(results[0].Int())
		case dyncall.UnsignedChar:
			*(*abi.CharUnsigned)(result) = abi.CharUnsigned(results[0].Uint())
		case dyncall.Short:
			*(*abi.Short)(result) = abi.Short(results[0].Int())
		case dyncall.UnsignedShort:
			*(*abi.ShortUnsigned)(result) = abi.ShortUnsigned(results[0].Uint())
		case dyncall.Int:
			*(*abi.Int)(result) = abi.Int(results[0].Int())
		case dyncall.Uint:
			*(*abi.IntUnsigned)(result) = abi.IntUnsigned(results[0].Uint())
		case dyncall.Long:
			*(*abi.Long)(result) = abi.Long(results[0].Int())
		case dyncall.UnsignedLong:
			*(*abi.LongUnsigned)(result) = abi.LongUnsigned(results[0].Uint())
		case dyncall.LongLong:
			*(*abi.LongLong)(result) = abi.LongLong(results[0].Int())
		case dyncall.UnsignedLongLong:
			*(*abi.LongLongUnsigned)(result) = abi.LongLongUnsigned(results[0].Uint())
		case dyncall.Float:
			*(*abi.Float)(result) = abi.Float(results[0].Float())
		case dyncall.Double:
			*(*abi.Double)(result) = abi.Double(results[0].Float())
		case dyncall.String:
			*(*abi.String)(result) = abi.NewString(results[0].String()) // FIXME allocate in C memory?
		case dyncall.Pointer:
			*(*unsafe.Pointer)(result) = results[0].UnsafePointer()
		default:
			panic("unsupported type " + results[0].Type().String())
		}
		return signature.Returns
	}
}

// Set links the given library using the specified shared
// library file name. The system linker will look for this
// file in the system library paths.
func Set(library Library, file string) error {
	lib := dlopen(file)
	if lib == nil {
		return errors.New(dlerror())
	}

	rtype := reflect.TypeOf(library).Elem()
	rvalue := reflect.ValueOf(library).Elem()

	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		value := rvalue.Field(i)

		if field.Type.Kind() != reflect.Func {
			continue
		}

		name := field.Tag.Get("ffi")
		if name == "" {
			name = field.Name
		}

		symbol := dlsym(lib, name)
		if symbol == nil {
			continue
		}

		getErr := rvalue.FieldByName("Error")

		switch fn := value.Addr().Interface().(type) {
		case *func(abi.Double) abi.Double:
			*fn = func(a abi.Double) abi.Double {
				vm := vm8.Get().(*dyncall.VM)
				vm.Reset()
				defer vm8.Put(vm)
				vm.PushFloat64(float64(a))
				return abi.Double(vm.CallFloat64(symbol))
			}
		case *func():
			*fn = func() {
				vm := dyncall.NewVM(0)
				defer vm.Free()
				vm.Call(symbol)
			}
		default:
			_ = fn
			value.Set(reflect.MakeFunc(field.Type, func(args []reflect.Value) []reflect.Value {
				var vm = dyncall.NewVM(4096)
				defer vm.Free()
				push := func(value reflect.Value) {
					switch value.Kind() {
					case reflect.Bool:
						vm.PushBool(value.Bool())
					case reflect.Int8:
						vm.PushInt8(int8(value.Int()))
					case reflect.Int16:
						vm.PushInt16(int16(value.Int()))
					case reflect.Int32:
						vm.PushInt32(int32(value.Int()))
					case reflect.Int64:
						vm.PushInt64(value.Int())
					case reflect.Uint8:
						u8 := uint8(value.Uint())
						vm.PushInt8(*(*int8)(unsafe.Pointer(&u8)))
					case reflect.Uint16:
						u16 := uint16(value.Uint())
						vm.PushInt16(*(*int16)(unsafe.Pointer(&u16)))
					case reflect.Uint32:
						u32 := uint32(value.Uint())
						vm.PushInt32(*(*int32)(unsafe.Pointer(&u32)))
					case reflect.Uint64:
						u64 := uint64(value.Uint())
						vm.PushInt64(*(*int64)(unsafe.Pointer(&u64)))
					case reflect.Float32:
						vm.PushFloat32(float32(value.Float()))
					case reflect.Float64:
						vm.PushFloat64(value.Float())
					case reflect.Pointer, reflect.UnsafePointer:
						vm.PushPointer(value.UnsafePointer())
					case reflect.String:
						s := abi.NewString(value.String())
						vm.PushPointer(unsafe.Pointer(s.Pointer()))
					case reflect.Struct:
						if value.Type().Implements(reflect.TypeOf([0]abi.IsPointer{}).Elem()) {
							vm.PushPointer(unsafe.Pointer(value.Interface().(abi.IsPointer).Pointer()))
						} else {
							panic("unsupported struct " + value.Type().String())
						}
					case reflect.Func:
						signature := newSignature(value.Type())
						ptr := dyncall.NewCallback(signature, newCallback(signature, value))
						vm.PushPointer(unsafe.Pointer(ptr))
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
				var returnsError bool
				switch field.Type.NumOut() {
				default:
					if field.Type.NumOut() > 1 {
						length := field.Type.NumOut()
						if field.Type.Out(length-1) == reflect.TypeOf([0]error{}).Elem() {
							returnsError = true
							length--
						}
						for i := 1; i < length; i++ {
							push(results[i].Addr())
						}
					}
					fallthrough
				case 1:
					rtype := field.Type.Out(0)
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
					case reflect.Uint8:
						u8 := vm.CallInt8(symbol)
						results[0].SetUint(uint64(*(*uint8)(unsafe.Pointer(&u8))))
					case reflect.Uint16:
						u16 := vm.CallInt16(symbol)
						results[0].SetUint(uint64(*(*uint16)(unsafe.Pointer(&u16))))
					case reflect.Uint32:
						u32 := vm.CallInt32(symbol)
						results[0].SetUint(uint64(*(*uint32)(unsafe.Pointer(&u32))))
					case reflect.Uint64:
						u64 := vm.CallInt64(symbol)
						results[0].SetUint(uint64(*(*uint64)(unsafe.Pointer(&u64))))
					case reflect.Float32:
						results[0].SetFloat(float64(vm.CallFloat32(symbol)))
					case reflect.Float64:
						results[0].SetFloat(float64(vm.CallFloat64(symbol)))
					case reflect.String:
						results[0].SetString(C.GoString((*C.char)(vm.CallPointer(symbol))))
					case reflect.UnsafePointer:
						results[0].SetPointer(vm.CallPointer(symbol))
					case reflect.Pointer:
						results[0] = reflect.NewAt(field.Type.Out(0).Elem(), unsafe.Pointer(vm.CallPointer(symbol)))
					case reflect.Struct:
						if rtype.Implements(reflect.TypeOf([0]abi.IsPointer{}).Elem()) {
							*(*unsafe.Pointer)(results[0].Addr().UnsafePointer()) = vm.CallPointer(symbol)
						} else {
							panic("unsupported struct " + field.Type.Out(0).String())
						}
					default:
						panic("unsupported type " + field.Type.Out(0).String())
					}
				case 0:
					vm.Call(symbol)
				}
				if returnsError {
					if results[0].IsZero() {
						if !getErr.IsValid() {
							panic("an error occured")
						}
						switch fn := getErr.Interface().(type) {
						case func() string:
							results[len(results)-1] = reflect.ValueOf(errors.New(fn()))
						}
					}
				}
				return results
			}))
		}
	}

	return nil
}
