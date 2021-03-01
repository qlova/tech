package main

import (
	"qlova.tech/api/cmd"
	"qlova.tech/api/free/dog"
)

func main() {
	cmd.Main(&dog.API, dog.API.Breeds)
}
