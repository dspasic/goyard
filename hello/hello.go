package main

import (
	"fmt"
	"log"

	"rsc.io/quote/v4"
	"github.com/dspasic/goyard/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message := greetings.Hello("Zig")

	names := []string{"Go", "Java", "Python"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	fmt.Println(quote.Go());
}
