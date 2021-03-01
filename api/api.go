package api

import "reflect"

//Function is a function provided by an API.
type Function struct {
	Endpoint string

	reflect.StructField
	Value reflect.Value
}

//Interface implements API.
//Embed it inside of a struct to create a new API.
type Interface interface {
	ConnectAPI(host string, functions []Function) error
}

//Tags is a set of API tags.
type Tags map[string]string

//Tag is a field that represents an API tag.
type Tag struct{}

//Connect connects to, and enables the API so that it can be used.
func Connect(api Interface) {
	rtype := reflect.TypeOf(api).Elem()
	rvalue := reflect.ValueOf(api).Elem()

	var host string
	var functions []Function
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		if field.Name == "Interface" {
			host = field.Tag.Get("api")
		}

		if field.Type.Kind() == reflect.Func {
			functions = append(functions, Function{
				Endpoint:    field.Tag.Get("api"),
				StructField: field,
				Value:       rvalue.Field(i),
			})
		}
	}

	api.ConnectAPI(host, functions)
}
