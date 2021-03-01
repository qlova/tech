package main

import (
	"qlova.tech/api/cmd"
	"qlova.tech/api/free/httpcat"
)

func main() {
	cmd.Main(&httpcat.API, func() (httpcat.Image, error) {
		return httpcat.API.Image(200)
	})
}
