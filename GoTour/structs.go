package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {
	myVertex := Vertex{x: 100, y: 200}

	fmt.Println(myVertex.x, myVertex.y)
}