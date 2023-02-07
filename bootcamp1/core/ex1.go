package main

import (
	"bufio"
	f "fmt"
	"os"
	"strconv"
	"strings"
)

var pl = f.Println

func main() {

	// ask user to input age
	pl("what is your age ?")
	reader := bufio.NewReader(os.Stdin)
	age, _ := reader.ReadString('\n')

	//trim the white space
	age = strings.TrimSpace(age)
	//convert age to int
	iAge, _ := strconv.Atoi(age)

	// age < 5 too young
	if iAge < 5 {
		pl("Too young to go to school")
	} else if (iAge > 5) && (iAge <= 18) {
		if (iAge > 5) && (iAge <= 10) {
			pl("You are in primary school")
		} else if (iAge > 10) && (iAge <= 14) {
			pl("You are in secondary school")
		} else {
			pl("You are in high school")
		}
	} else {
		pl("You are in college")
	}

}
