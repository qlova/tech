package api

import (
	"io"
	"reflect"
)

//Function is a function provided by an API.
type Function struct {
	Endpoint string

	reflect.StructField
	Value reflect.Value
}

//Protocol determines the protocol of an API.
type Protocol interface {
	DecodeValue(io.Reader, interface{}) error
	EncodeValue(io.Writer, interface{}) error
}

//Interface implements API.
//Embed it inside of a struct to create a new API.
type Interface interface {
	Import(host string, protocol Protocol, functions []Function) error
}

//Tags is a set of API tags.
type Tags map[string]string

//Tag is a field that represents an API tag.
type Tag struct{}

/* (When Go gets generics)

func Import[type T Interface](api T) T {
	Connect(api)
	return api
}

*/

//Connect connects to, and enables the API so that it can be used.
func Connect(api Interface) {
	rtype := reflect.TypeOf(api).Elem()
	rvalue := reflect.ValueOf(api).Elem()

	var host string
	var functions []Function

	var protocol Protocol

	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		if field.Name == "API" {
			host = field.Tag.Get("api")
		}

		if field.Name == "Protocol" {
			protocol = rvalue.Field(i).Interface().(Protocol)
		}

		if field.Type.Kind() == reflect.Func {
			functions = append(functions, Function{
				Endpoint:    field.Tag.Get("api"),
				StructField: field,
				Value:       rvalue.Field(i),
			})
		}
	}

	api.Import(host, protocol, functions)
}
