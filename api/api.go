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

//Importer is an API that can be imported.
type Importer interface {
	Import(host string, protocol Protocol, functions []Function) error
}

//Exporter is an API that can be exported.
type Exporter interface {
	Export(port string, protocol Protocol, functions []Function) error
}

//Tags is a set of API tags.
type Tags map[string]string

//Tag is a field that represents an API tag.
type Tag struct{}

/* (When Go gets generics)

func Import[type T Importer](api T) T {
	if err := Connect(api); err != nil {
		//create wrappers that attempt to reconnect and/or
		//return an error.
	}

	return api
}

*/

//Export exports the API and serves it.
func Export(api Exporter) error {
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

	return api.Export(host, protocol, functions)
}

//Connect connects to, and enables the API so that it can be used.
func Connect(api Importer) error {
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

	return api.Import(host, protocol, functions)
}
