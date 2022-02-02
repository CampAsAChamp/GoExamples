package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	// Get a greeting and print it
	message := greetings.Hello("Nick")
	fmt.Println(message)
}
