package main

import (
	"fmt"

	"rsc.io/quote/v4"
)

func main() {
	fmt.Println("quote.go: ", quote.Go())
}

func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
