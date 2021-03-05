package rest

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"regexp"

	"qlova.tech/api"
	"qlova.tech/enc"
)

var pathReplacer = regexp.MustCompile(`{.*?=%v}`)

//Import implements a naive REST api.Importer
func (*API) Import(def api.Definition) error {
	if def.Protocol == nil {
		def.Protocol = enc.JSON
	}

	for i := range def.Functions {
		fn := def.Functions[i]

		pattern := fn.Tag

		pattern = pathReplacer.ReplaceAllLiteralString(pattern, "%v")

		//TODO prebuild with regex like in Export.

		fn.Value.Set(reflect.MakeFunc(fn.Type, func(args []reflect.Value) (results []reflect.Value) {
			results = make([]reflect.Value, fn.Type.NumOut())

			handle := func(err error) {
				var returned bool
				for i := 0; i < fn.Type.NumOut(); i++ {
					if fn.Type.Out(i) == reflect.TypeOf([0]error{}).Elem() {
						results[i] = reflect.ValueOf(err)
						returned = true
						continue
					}
					results[i] = reflect.Zero(fn.Type.Out(i))
				}
				if !returned {
					panic(err)
				}
			}

			var converted = make([]interface{}, len(args))
			for i := range args {
				converted[i] = url.QueryEscape(fmt.Sprint(args[i].Interface()))
			}

			resp, err := http.Get(def.Tag + fmt.Sprintf(pattern, converted...))
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
					if err := def.Protocol.Decode(val, resp.Body); err != nil {
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
