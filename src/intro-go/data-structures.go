package main

import "fmt"

func main() {
	/*
	Arrays
	Slice
	Maps
	Struct
	 */

	//Arrays
	// In Go, arrays are fixed size
	var nums [5] int // int array of 5 items
	fmt.Println(nums) // [0 0 0 0 0]

	var names [3] string // string array of 3 items
	fmt.Println(names, len(names), cap(names)) // [  ] 3 3


	var raining [3] bool // bool array of 3 items
	fmt.Println(raining) // [false false false]

	names[0] = "iOS"
	names[1] = "Android"
	names[2] = "Symbian"
	fmt.Println(names) //[iOS Android Symbian]

	//Slices
	// Slices as a view into an array and its a light-weight data structure that's more flexible than arrays
	x := make([] int, 5) // creates a slice of 5 elements, capacity = 5
	fmt.Println(x) // [0 0 0 0 0]

	// The make() function allocates and initializes an array of the specified type.
	// make([] int, 5) - is a slice of five elements
	// make([] int, 2, 3) - is a slice of two elements, but with a maximum capacity of three

	y := make([] int, 2, 3) // Creates a slice of 2 elements, capacity = 3 [0 0]
	fmt.Println(y)

	odds := [] int {1, 3, 5}
	fmt.Println(odds)

	//difference between array and slice - inside the bracket if nothing is mentioned then slice otherwise array
    var nums100  [5] int // nums is an array
    var nums200 [] int // nums is now a slice
    fmt.Println(nums100, nums200)

	//Understanding the Behavior slices
	original := [] int{1, 2, 3, 4}
	other := original
	fmt.Println(original)
	fmt.Println(other)
	other[2] = 8
	fmt.Println(original)
	fmt.Println(other)

	other = append(original, 5) // "Other" points to a new slice (as it has exceeded its capacity of four) now make changes to others, original won't change
	fmt.Println(original)
	fmt.Println(other)

	other[2] = 9
	fmt.Println(other)
	fmt.Println(original)

	a := make([] int, 2, 4)
	b := a
	fmt.Println(a)
	fmt.Println(b)

	b = append(a, 5) // x still points to the original two numbers, and y now points to the same numbers, plus the additional one appended to x. This is because y ( as well asx) has the capacity of four and has room for up to four items.
	fmt.Println(a)
	fmt.Println(b)
	a[1] = 99 // now when we modify the second item in y, both x and y are affected
	b[0] = 88
	fmt.Println(a)
	fmt.Println(b)

	// Slicing on Slice/Arrays
	// We can perform slicing (extracting a range of values) on arrays and slices
	var c[3] string
	c[0] ="iOS"
	c[1] = "Android"
	c[2] = "Windows"

	fmt.Println(c)

	d := c[0:2]
	fmt.Println(d, cap(d), len(d))

	e := c[0:2:2] // We can change the capacity of the slice b, by specifying the capacity as the third argument in the slicing
	fmt.Println(e, cap(e), len(e))

	// Maps
	// Map which implements a hash table.
    heights := map[string]int{}
	fmt.Println(heights)

	height1 := make(map [string]int)
	fmt.Println(height1)

	// declare and initialize the map variable with some values
	weights := map[string] float32 {
		"Peter": 45.9,
		"Joan": 56.8,
	}
	fmt.Println(weights)

	heights["Peter"] = 178

	// delete() function  - to delete key/value pair
	delete(heights, "Peter")

	//To check whether a key exists in the map
	if value, exists := heights["Peter"]; exists  {
		fmt.Println(value)
	} else {
		fmt.Println("Key does not exists")
	}

	//Looping through key/value pairs
	for k,v := range weights {
		fmt.Println(k,v)
	}

	// Structs
	// Go don't have classes, but it supports structs.
	type Point struct {
		X float64
		Y float64
	}
	//We can also define methods on Structs
	// A method is a function with a special receiver argument
	// To add a method to a struct, define a function with the struct passed in as an argument defined before the function name
	/*
	func (p point) Length() float64 {
		return math.Sqrt(math.Pow(p.X, 2.0) + math.Pow(p.Y, 2.0))
	}
	*/

	ptA := Point{5, 6}
	fmt.Println(ptA)

	// Create a reference to another struct, use the & character
	ptB := &ptA
	fmt.Println(ptB)
	ptB.X = 55
	fmt.Println(ptA)
	fmt.Println(ptB)

	// Creating a new instance of the Point struct
	pt1 := Point{X:100, Y:200}
	pt2 := pt1
	fmt.Println(pt1)
	fmt.Println(pt2)
}