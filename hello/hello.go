package main

import (
	"fmt"
	"log"

	"github.com/dspasic/goyard/greetings"
	"rsc.io/quote/v4"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Go", "Java", "Python", "Rust", "Zig"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	fmt.Println(quote.Go())
}
