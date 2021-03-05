/*
Package api contains a standard interface for an API.
This allows Go code to easily import an API and use it
without a bloated wrapper and without the need to care
about the transport or protocol layer. These API functions
are called like native funtions.

By convention, API's are global anonymous struct variables:

	var API struct {
		rest.API `https://example.com`

		Echo func(message string) string `/echo?message=%v`
	}

This rest API can be imported by calling api.Connect(&API),
and then API.Echo("Hello World") will be called over
HTTPS/REST to 'https://example.com/echo' with the given string
and return the result (or panic if there is an error).

Documentation on how different tags are handled can be found
in the included rest package.

The package can also export the API instead of importing it,
By calling api.Export(&API), a local HTTP/REST server will
listen on PORT and return the response of the Echo call to
any clients that call the '/echo' endpoint with a message
string (you do need to provide an implementation of course).

	API.Echo = func(message string) string { return message }

Exactly how things are imported/exported is implementation
defined, if the implementation does not support an argument
or tag, it should return an error, either on Connect or Export.
*/
package api

import (
	"context"
	"reflect"

	"qlova.tech/enc"
)

//Key is an API key, embed this in an API definition when some kind
//of authentication key is required in order to authenticate with the
//API, the value of this key will be used to authenticate with an API.
//The exact method and supported tags are implementation specific.
type Key string

//Tagger types can provide a fallback to a missing api tag for a
//field of the implementing type.
type Tagger interface {
	Tag() string
}

//Definition is a definition of an API. It can be constructed manually but
//should normally be handled by this package.
type Definition struct {
	//Tag is the tag of the API field.
	Tag string

	//Key is the API Key for this API.
	Key struct {
		Pointer *Key
		Tag     string
	}

	//Protocol is the protocol that the API should encodes information in.
	Protocol enc.Format

	//Functions is a list of functions this API provides.
	Functions []Function
}

//Function is a API-provided function.
type Function struct {
	Name, Tag string

	Type  reflect.Type
	Value reflect.Value
}

//Importer is an API that can be imported.
type Importer interface {
	Import(Definition) error
}

//Exporter is an API that can be exported.
type Exporter interface {
	Export(Definition) error
}

//Authenticator is a type that can Authenticate itself from a Request.
type Authenticator interface {
	Authenticate(Request) error
}

//Request is a raw api request, used for identity, verification and authentication.
//If Target == Origin then this is a local request from within the same Go process.
type Request interface {
	context.Context

	//Key returns the value of the environmental variable, option, context
	//or cookie Key that exists to authenticate, identify or verify the request.
	Key(key Key) string

	//Target is the request's target endpoint. Either an IP address, a URL or
	//another string that represents the location that this request was sent to.
	Target() string

	//Origin is the origin of the request, either an IP address or another string
	//that represents the location that this request was sent from.
	Origin() string

	//Body reads the request data so that it can be signature verified.
	Data() []byte
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

func tagOf(field reflect.StructField, value reflect.Value) string {
	tag := field.Tag.Get("api")

	if tag == "" {
		tag = string(field.Tag)
	}

	if tag == "" && field.Type.Implements(reflect.TypeOf([0]Tagger{}).Elem()) {
		tag = value.MethodByName("Tag").Interface().(func() string)()
	}

	return tag
}

func definitionOf(api interface{}) Definition {
	rtype := reflect.TypeOf(api).Elem()
	rvalue := reflect.ValueOf(api).Elem()

	var def Definition

	//TODO support embedded APIs.

	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		tag := tagOf(field, rvalue.Field(i))

		if field.Name == "API" {
			def.Tag = tag
		}

		if field.Name == "Protocol" && field.Type.Implements(reflect.TypeOf([0]enc.Format{}).Elem()) {
			def.Protocol = rvalue.Field(i).Interface().(enc.Format)
		}

		if field.Name == "Key" && field.Type == reflect.TypeOf(Key("")) {
			def.Key.Pointer = rvalue.Field(i).Addr().Interface().(*Key)
			def.Key.Tag = tag
		}

		if field.Type.Kind() == reflect.Func {
			def.Functions = append(def.Functions, Function{
				Name:  field.Name,
				Tag:   tag,
				Type:  field.Type,
				Value: rvalue.Field(i),
			})
		}
	}

	return def
}

//Export exports the API and serves it.
func Export(api Exporter) error {
	return api.Export(definitionOf(api))
}

//Connect connects to, and enables the API so that it can be used.
func Connect(api Importer) error {
	return api.Import(definitionOf(api))
}
