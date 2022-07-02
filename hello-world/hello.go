package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Welcome to Go💙 land")
	quote := getAGoQuote()
	fmt.Println("💡", quote)
}

func getAGoQuote() string {
	return quote.Go()
}
