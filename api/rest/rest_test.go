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

	Print func() `rest:"/print"`

	Echo func(string) string `rest:"/echo/{message=%v}"`

	Concat func(a, b string) string `rest:"/concat?a=%v&b=%v"`
}

func init() {
	API.Print = Print
	API.Echo = Echo
	API.Concat = Concat
}

func Print() {
	fmt.Println("Hello World")
}

func Echo(message string) string {
	return message
}

func Concat(a, b string) string {
	return a + b
}

func Test_Rest(t *testing.T) {
	//Host the API on localhost.
	go api.Export(&API)

	var local = API

	//Nil out the functions so that they can't be
	//called locally.
	local.Print = nil
	local.Echo = nil
	local.Concat = nil

	local.Host = "http://localhost:" + os.Getenv("PORT")
	api.Connect(&local)

	local.Print()
	should.Be("HelloServer")(local.Echo("HelloServer")).Test(t)

	should.Be("This string is concatenated")(local.Concat("This string", " is concatenated")).Test(t)
}
