package main

import (
	"fmt"
	"pkg/pkg"
)

func main() {
	// Demonstrating having a Global function that returns a internal map
	// can also be manipluated outside.
	// Welcome to PHP style.
	fmt.Print("Hello, World")
	fmt.Printf("Hello, global: %s", pkg.GetGlobal())

	g := pkg.GetGlobal()
	g["extern"] = "Richtig Prutschen"

	fmt.Printf("Hello, global: %s", pkg.GetGlobal())
}
