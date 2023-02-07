package main

import (
	"bufio"
	f "fmt"
	"os"
)

var pl = f.Println

func main() {
	pl("Hello World")
	pl("whats is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		pl(err)
	} else {
		pl("Hello ", name)
	}
}
