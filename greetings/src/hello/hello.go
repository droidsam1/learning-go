package main

import (
	"fmt"

	"example.com/greeting"
)

func main() {
	message := greeting.Hello("Sam")
	fmt.Println(message)
}
