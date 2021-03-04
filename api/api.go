package api

import (
	"context"
	"io"
	"reflect"
)

//Function is a API-provided function.
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
	Setenv(key, value string) error

	Import(host string, protocol Protocol, functions []Function) error
}

//Exporter is an API that can be exported.
type Exporter interface {
	Export(port string, protocol Protocol, functions []Function) error
}

//Authenticator is a type that can Authenticate itself from a Request.
type Authenticator interface {
	Authenticate(Request) error
}

//Request is a raw api request, used for identity, verification and authentication.
//If Target == Origin then this is a local request from within the same Go process.
type Request interface {
	context.Context

	//Getenv returns a named environmental variable, option, context or cookie
	//that exists to authenticate, identify or verify a request.
	Getenv(key string) string

	//Target is the request's target endpoint. Either an IP address, a URL or
	//another string that represents the location that this request was sent to.
	Target() string

	//Origin is the origin of the request, either an IP address or another string
	//that represents the location that this request was sent from.
	Origin() string
}

/* (When Go gets generics)

type Handler[type T Authenticator] interface{
	Handle(T)
}

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
