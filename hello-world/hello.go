package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Welcome to Goð land")
	quote := getAGoQuote()
	fmt.Println("ð¡", quote)
}

func getAGoQuote() string {
	return quote.Go()
}
