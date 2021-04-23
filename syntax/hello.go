package main

import (
	"fmt"
	"os"
	"strings"
)

/*
Packages
********
fmt - provides functions for formatting text and for reading formatted text
os - provides platform-independent operating system variables and functions
strings - provides functions for manipulating strings
*/

func main() {

	var argValues string  //defining an argValues
	if len(os.Args) > 1 { //checking the argument values for ex: go run hello.go hello bhuppal kumar
		argValues = strings.Join(os.Args[1:], " ")
	}
	fmt.Println(argValues)
}
