package main

import (
	"fmt"
	"log"

	"qlova.tech/api/free/catfacts"
)

func main() {
	fact, err := catfacts.API.Random()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(fact.Text)
}
