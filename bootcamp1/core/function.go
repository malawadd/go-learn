package main

import (
	"fmt"
	f "fmt"
)

var pl = f.Println

func main() {

	// iAge := 9
	// if (iAge > 1) && (iAge <= 18) {
	// 	pl("You are a kid")
	// } else if (iAge > 18) && (iAge <= 60) {
	// 	pl("You are an adult")
	// } else {
	// 	pl("You are a senior citizen")
	// }

	//formated output
	fmt.Printf("Hello %s, you are %d years old , and you have %.2f in bank, yes that's %t", "John", 20, 3.015, true)
}
