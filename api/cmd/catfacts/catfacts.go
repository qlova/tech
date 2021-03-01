package main

import (
	"qlova.tech/api/cmd"
	"qlova.tech/api/free/catfacts"
)

func main() {
	cmd.Main(&catfacts.API, catfacts.API.Random)
}
