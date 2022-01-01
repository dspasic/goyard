package main

import (
    "fmt"
    "log"

    "github.com/dspasic/goyard/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    names := []string{"Go", "Java", "Python"}

    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(messages)
}
