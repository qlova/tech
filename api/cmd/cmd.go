//Package cmd provides an easy way to create command line apps that wrap an API.
package cmd

import (
	"fmt"
	"image"
	"os"
	"reflect"
	"strings"

	"github.com/sanity-io/litter"
	"qlova.tech/api"
)

func dump(values ...interface{}) {
	for _, value := range values {
		switch v := value.(type) {
		case fmt.Stringer:
			fmt.Println(v.String())
		case image.Image:
			display(v)
		case string:
			fmt.Println(v)
		default:
			litter.Dump(value)
		}
	}
}

func call(fn interface{}) {
	results := reflect.ValueOf(fn).Call(nil)
	var converted []interface{}
	for _, result := range results {
		if result.Type() != reflect.TypeOf([0]error{}).Elem() {
			converted = append(converted, result.Interface())
		}
	}
	dump(converted...)
}

//Main runs a main function for the given command, if
//no arguments are provided, it runs the entrypoint.
func Main(API api.Importer, entrypoint interface{}) {
	if len(os.Args) < 2 {
		call(entrypoint)
		return
	}

	if len(os.Args) > 1 {
		arg := os.Args[1]
		if strings.HasPrefix(arg, "-") {
			arg = strings.TrimPrefix(arg, "-")

			rtype := reflect.TypeOf(API).Elem()

			for i := 0; i < rtype.NumField(); i++ {
				field := rtype.Field(i)

				if strings.ToLower(field.Name) == strings.ToLower(arg) {
					call(reflect.ValueOf(API).Elem().Field(i).Interface())
					return
				}
			}
		}
	}

	fmt.Println("command not recognised")
	os.Exit(1)
}
