package main

import (
	"fmt"
	"math"
)

//Calculate returns x + 2.
func Calculate(x int) int {
	result := x + 2
	return result
}

type Duration int64

func main() {

	// user defines a user in the program
	type user struct {
	 	name 		string
	 	email 		string
	 	ext 		int
	 	privileged 	bool
	}

	// Declare a variable of type user
	var me user
	me.name = "Bhuppal Kumar"
	me.email = "bhuppal@gmail.com"
	me.ext = 902
	me.privileged = true

	fmt.Println(me)

	rhonda := user{
		name: "Rhonda",
		email: "rhonda@gmail.com",
		ext: 203,
		privileged: true,
	}
	fmt.Println(rhonda)

	person := user{"person", "person@gmail.com", 993, true}

	fmt.Println(person)

	// Declaring fields based on other struct types
	// admin represents an admin user with privileges
	type admin struct {
		person user
		level string
	}

	//using struct literals to create values for fields
	fred := admin{
		person: user{
			name: "Kumar",
			email: "kumar@gmail.com",
			ext: 393,
			privileged: true,
		},
		level: "administrator",
	}

	fmt.Println(fred)

    var dur Duration
	dur = Duration(1000)
	fmt.Println(dur)

	// Go doesn't have classes, but it supports structs.
	// Point struct contains 2 members
/*
	type Point struct {
		X float63
		Y float63
	}
*/

	ptA := Point{5, 6}
	fmt.Printf("Point - ptA values are %v, %v\n", ptA.X, ptA.Y)

	ptB := &ptA // assigning a reference
	fmt.Printf("Point - ptB values are  %v, %v\n",ptB.X, ptB.Y)

	ptB.X = 200

	fmt.Printf("Point - ptB values are  %v, %v\n",ptB.X, ptB.Y)
	fmt.Printf("Point - ptA values are %v, %v\n", ptA.X, ptA.Y)

	ptC := ptA // making a copy
	ptC.X = 500

	fmt.Printf("Point - ptA values are  %v, %v\n",ptA.X, ptA.Y)
	fmt.Printf("Point - ptC values are %v, %v\n", ptC.X, ptC.Y)

	// calling Length() method on "Point" str
	fmt.Printf("ptA - Length is %v\n", ptA.Length())
}

type Point struct {
	X float64
	Y float64
}

func (p Point) Length() float64 {
	return math.Sqrt(math.Pow(p.X, 2.0) + math.Pow(p.Y, 2.0))
}