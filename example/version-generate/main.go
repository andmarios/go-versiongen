package main

import (
	"log"

	"github.com/andmarios/go-versiongen"
)

func main() {
	err := versiongen.Create()
	if err != nil {
		log.Fatalln(err)
	}
}
