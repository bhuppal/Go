package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return //A return statement without arguments returns the named return values.
	// This is known as a "naked" return
}

func main() {
	a, b := split(45)
	fmt.Println(a, b)
}
