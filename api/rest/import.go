//Package rest provides a REST API for use with qlova.tech/api
package rest

import (
	"fmt"
	"io"
	"net/http"
	"reflect"

	"qlova.tech/api"
)

//API is a REST api.Interface
type API struct{}

//Import implements api.Interface.Import
func (*API) Import(host string, protocol api.Protocol, functions []api.Function) error {
	if protocol == nil {
		protocol = Protocol{}
	}

	for i := range functions {
		fn := functions[i]
		fn.Value.Set(reflect.MakeFunc(fn.Type, func(args []reflect.Value) (results []reflect.Value) {
			results = make([]reflect.Value, fn.Type.NumOut())

			handle := func(err error) {
				for i := 0; i < fn.Type.NumOut(); i++ {
					if fn.Type.Out(i) == reflect.TypeOf([0]error{}).Elem() {
						results[i] = reflect.ValueOf(err)
						continue
					}
					results[i] = reflect.Zero(fn.Type.Out(i))
				}
			}

			var converted = make([]interface{}, len(args))
			for i := range args {
				converted[i] = args[i].Interface()
			}

			resp, err := http.Get(host + fmt.Sprintf(fn.Endpoint, converted...))
			if err != nil {
				handle(err)
				return
			}

			for i := 0; i < fn.Type.NumOut(); i++ {
				results[i] = reflect.Zero(fn.Type.Out(i))
			}

			if fn.Type.NumOut() > 0 {
				T := results[0].Type()

				var ptr reflect.Value

				if T.Kind() == reflect.Ptr {
					results[0] = reflect.New(T.Elem())
					ptr = results[0]
				} else {
					ptr = reflect.New(T)
				}

				val := ptr.Interface()

				if decoder, ok := val.(interface{ Decode(io.Reader) error }); ok {
					decoder.Decode(resp.Body)
				} else {
					if err := protocol.DecodeValue(resp.Body, val); err != nil {
						handle(err)
						return
					}
				}

				if T.Kind() != reflect.Ptr {
					results[0] = ptr.Elem()
				}
			}

			return
		}))
	}

	return nil
}
