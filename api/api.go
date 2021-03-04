package api

import (
	"context"
	"io"
	"reflect"
)

//Key is an API key, embed this in an API definition when some kind
//of key is required in order to authenticate with the API.
type Key string

//Tagger types can provide a fallback struct-tag. For packages that
//use reflection and check for a Tagger.
type Tagger interface {
	Tag() string
}

//Definition is a definition of an API.
type Definition struct {
	//Tag is the tag of the API field.
	Tag string

	//Key is the API Key for this API.
	Key struct {
		Pointer *Key
		Tag     string
	}

	//Protocol is the protocol that the API encodes information in.
	Protocol Protocol

	//Functions is a list of functions this API provides.
	Functions []Function
}

//Function is a API-provided function.
type Function struct {
	Name, Tag string

	Type  reflect.Type
	Value reflect.Value
}

//Protocol determines the protocol of an API.
type Protocol interface {
	DecodeValue(io.Reader, interface{}) error
	EncodeValue(io.Writer, interface{}) error
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

		if field.Name == "Protocol" && field.Type.Implements(reflect.TypeOf([0]Protocol{}).Elem()) {
			def.Protocol = rvalue.Field(i).Interface().(Protocol)
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
