package main

import (
	"fmt"
)

//go:generate go run version-generate/main.go

func main() {
	fmt.Println("This is version: " + vgVersion)
	fmt.Println("Commit         : " + vgHash)
}
