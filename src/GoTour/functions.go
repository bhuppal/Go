package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Printf("The addition of 2 numbers: %d", add(29, 35))
}
