package rest_test

import (
	"fmt"
	"os"
	"testing"

	"qlova.org/should"
	"qlova.tech/api"
	"qlova.tech/api/rest"
)

var API struct {
	rest.API

	Print func() `api:"/print"`

	Echo func(string) string `api:"/echo/{message=%v}"`
}

func init() {
	API.Print = Print
	API.Echo = Echo
}

func Print() {
	fmt.Println("Hello World")
}

func Echo(message string) string {
	return message
}

func Test_Rest(t *testing.T) {

	//Host the API on localhost.
	go api.Export(&API)

	var local = API

	//Nil out the functions so that they can't be
	//called locally.
	local.Print = nil
	local.Echo = nil

	local.Host = "http://localhost" + os.Getenv("PORT")
	api.Connect(&local)

	local.Print()
	should.Be("Hello Server")(local.Echo("Hello Server")).Test(t)
}
