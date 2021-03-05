//Package rest provides a REST API for use with qlova.tech/api
package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
	"qlova.tech/api"
	"qlova.tech/enc"
)

type argLocation byte

const (
	inBody argLocation = iota
	inQuery
	inPath
)

type argument struct {
	location argLocation
	name     string
}

var inPathRegex = regexp.MustCompile("{(.*?)=%v}")
var inQueryRegex = regexp.MustCompile(`[?&](.*?)=%v`)

//Export implements a naive REST api.Exporter
func (*API) Export(def api.Definition) error {
	if def.Protocol == nil {
		def.Protocol = enc.JSON
	}

	var router = mux.NewRouter()

	for i := range def.Functions {
		fn := def.Functions[i]

		var pattern string = fn.Tag
		var locations = make([]argument, fn.Type.NumIn())

		if pattern == "" {
			pattern = fn.Name
		}

		//Does the endpoint contains parameters? this
		//means that the Go function's arguments are
		//found inside the path or query of the request.
		if strings.Contains(fn.Tag, "%") {

			//Match path parameters.
			matches := inPathRegex.FindAllStringSubmatch(pattern, -1)

			var i int
			for _, match := range matches {
				for _, submatch := range match {
					locations[i] = argument{
						name:     submatch,
						location: inPath,
					}
				}
				i++
			}

			//Cleanup the pattern.
			pattern = strings.Replace(pattern, "=%v}", "}", -1)

			split := strings.Split(pattern, "?")

			pattern = split[0]

			//Process the query.
			if len(split) > 1 {
				query := "?" + split[1]

				matches = inQueryRegex.FindAllStringSubmatch(query, -1)

				for _, match := range matches {
					for _, submatch := range match {
						locations[i] = argument{
							name:     submatch,
							location: inQuery,
						}
					}
					i++
				}
			}
		}

		router.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			var in = make([]reflect.Value, fn.Type.NumIn())

			for i, arg := range locations {
				switch arg.location {
				case inPath:
					in[i] = reflect.ValueOf(vars[arg.name])
				case inQuery:
					in[i] = reflect.ValueOf(r.URL.Query().Get(arg.name))
				default:
					panic("not implemented")
				}
			}

			results := fn.Value.Call(in)

			if len(results) == 1 {
				if err := json.NewEncoder(w).Encode(results[0].Interface()); err != nil {
					log.Println(err)
				}
				return
			}

			var converted = make([]interface{}, 0, len(results))

			for _, v := range results {
				converted = append(converted, v.Interface())
			}

			if err := json.NewEncoder(w).Encode(converted); err != nil {
				log.Println(err)
			}
		})
	}

	return http.ListenAndServe(os.Getenv("PORT"), router)
}
