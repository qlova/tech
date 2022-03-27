package data

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"qlova.tech/use/js"
)

/*
	Interfaces based on HTTP methods
	can be implemented on a Go value
	to enable them to be handled
	like an HTTP resource.
*/
type (

	// ResourceWithGet handles a HTTP GET
	// operation. The reciever
	// value will be returned to the
	// caller after this function
	// returns a nil error.
	ResourceWithGet interface {
		Get(context.Context) error
	}

	// ResourceWithSearch handles a HTTP SEARCH
	// operation. Semantically, this should
	// be used instead of MethodGet or MethodQuery
	// when multiple 'result' resources are
	// returned. The reciever
	// value will be returned to the
	// caller after this function
	// returns a nil error.
	ResourceWithSearch interface {
		Search(context.Context) error
	}

	// ResourceWithPost handles a HTTP POST
	// operation. The patch returned
	// will be applied to the caller's
	// state.
	ResourceWithPost interface {
		Post(context.Context) error
	}

	// ResourceWithPut handles a HTTP PUT
	// operation. No data is returned.
	ResourceWithPut interface {
		Put(context.Context) error
	}

	// ResourceWithDelete handles a HTTP DELETE
	// operation. No data is returned.
	ResourceWithDelete interface {
		Delete(context.Context) error
	}

	// ResourceWithPatch handles a HTTP PATCH
	// operation. The patch must be applied
	// atomically, otherwise an error is returned
	// and the patch is not applied.
	ResourceWithPatch interface {
		Patch(context.Context, Patch) error
	}
)

// Get returns javacript that will get the requested
// resource and apply it to the client's state.
func Get(resource ResourceWithGet) js.String {
	return js.String(fmt.Sprintf("ajax.get('%v')", PathOf(resource)))
}

// Search returns javacript that will query the
// server for the resource and apply it to the
// client's state.
func Search(resource ResourceWithSearch) js.String {
	return js.String(fmt.Sprintf("ajax.search('%v')", PathOf(resource)))
}

// Post returns javascript that will post the given
// resource and apply any of the resulting patches.
func Post(resource ResourceWithPost) js.String {
	return js.String(fmt.Sprintf("ajax.post('%v')", PathOf(any(resource))))
}

// Put returns javascript that will write the given
// resource idempotently.
func Put(resource ResourceWithPut) js.String {
	return js.String(fmt.Sprintf("ajax.put('%v')", PathOf(resource)))
}

// Delete returns javascript that will delete the given
// resource idempotently.
func Delete(resource ResourceWithDelete) js.String {
	return js.String(fmt.Sprintf("ajax.delete('%v')", PathOf(resource)))
}

type resource struct {
	rtype reflect.Type
}

func (res resource) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handle := func(err error) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	MethodNotAllowed := func() {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	BadRequest := func(err error) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	iface := reflect.New(res.rtype).Interface()
	switch r.Method {
	case http.MethodGet:
		handler, ok := iface.(ResourceWithGet)
		if !ok {
			MethodNotAllowed()
			return
		}

		//TODO decode query.

		if err := handler.Get(ctx); err != nil {
			handle(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(iface); err != nil {
			handle(err)
			return
		}
	case "SEARCH":
		handler, ok := iface.(ResourceWithSearch)
		if !ok {
			MethodNotAllowed()
			return
		}
		if err := json.NewDecoder(r.Body).Decode(iface); err != nil {
			BadRequest(err)
			return
		}
		if err := handler.Search(ctx); err != nil {
			handle(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(iface); err != nil {
			handle(err)
			return
		}
	case http.MethodPost:
		handler, ok := iface.(ResourceWithPost)
		if !ok {
			MethodNotAllowed()
			return
		}
		if err := json.NewDecoder(r.Body).Decode(iface); err != nil {
			BadRequest(err)
			return
		}
		if err := handler.Post(ctx); err != nil {
			handle(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(iface); err != nil {
			handle(err)
			return
		}
	case http.MethodPut:
		handler, ok := iface.(ResourceWithPut)
		if !ok {
			MethodNotAllowed()
			return
		}
		if err := json.NewDecoder(r.Body).Decode(iface); err != nil {
			BadRequest(err)
			return
		}
		if err := handler.Put(ctx); err != nil {
			handle(err)
			return
		}
	case http.MethodDelete:
		handler, ok := iface.(ResourceWithDelete)
		if !ok {
			MethodNotAllowed()
			return
		}
		if err := handler.Delete(ctx); err != nil {
			handle(err)
			return
		}
	case http.MethodPatch:
		handler, ok := iface.(ResourceWithPatch)
		if !ok {
			MethodNotAllowed()
			return
		}
		var patch Patch
		if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
			handle(err)
			return
		}
		if err := handler.Patch(ctx, patch); err != nil {
			handle(err)
			return
		}
	}
}

func pathsOf(rvalue reflect.Value, into map[string]http.Handler) {
	if rvalue.Kind() == reflect.Ptr {
		return
	}

	path := PathOf(rvalue.Addr().Interface())
	path = "/data/" + strings.Replace(path, ".", "/", -1)

	into[path] = resource{rvalue.Type()}

	if rvalue.Kind() == reflect.Struct {
		for i := 0; i < rvalue.NumField(); i++ {
			var field = rvalue.Field(i)
			pathsOf(field, into)
		}
	}
}

// PathsOf returns a map of all the ajax handlers embedded
// in the given value.
func HandlersOf(states map[reflect.Type]any) map[string]http.Handler {
	var paths = make(map[string]http.Handler)
	for _, state := range states {
		rvalue := reflect.ValueOf(state)
		if rvalue.Kind() == reflect.Ptr {
			rvalue = rvalue.Elem()
		}
		pathsOf(rvalue, paths)
	}
	return paths
}
