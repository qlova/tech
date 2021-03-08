//Package rest provides a REST API for use with qlova.tech/api
package rest

import "reflect"

//API is a REST api.Interface
type API struct {
	Host string
}

//Tag is a fallback for the struct field tag.
func (api API) Tag() reflect.StructTag {
	return reflect.StructTag(`rest:"` + api.Host + `"`)
}
