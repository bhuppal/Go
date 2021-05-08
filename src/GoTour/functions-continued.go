package main

import "fmt"

func multiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Printf("The product of 2 numbers: %d", multiply(5, 5))
}
