package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("One", "Two")
	fmt.Printf("The swap of 2 strings: %s %s", a, b)
}
