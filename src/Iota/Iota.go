package main

import "fmt"

const (
	a = iota
	b
	c
)

const (
	d = iota
	e
	f
)

func main() {

	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %v\n", c, c)

	fmt.Printf("%T %v\n", d, d)
	fmt.Printf("%T %v\n", e, e)
	fmt.Printf("%T %v\n", f, f)
}
