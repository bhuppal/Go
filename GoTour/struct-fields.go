package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{X: 320, Y:872}
	fmt.Println(v.X,v.Y)
}
