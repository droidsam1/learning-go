package main

import (
	"fmt"
	"log"

	"example.com/greeting"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greeting.Hello("Sam")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
