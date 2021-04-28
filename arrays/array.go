package main

import "fmt"

func main() {
	fmt.Println("GO - Arrary concepts")

	array := [5]int{10, 20, 30, 40, 50}
	fmt.Println(array)

	//Declaring an array with Go calculating size
	arraysize := [...]int{10, 20, 30, 40, 50, 60}
	fmt.Println(arraysize)

	//Accessing array pointer elements
	arraypointers := [5] *int{0: new(int), 1: new(int)}
	*arraypointers[0] = 100
	*arraypointers[1] = 200

	fmt.Println(*arraypointers[0])

	//Assigning one array to another of the same type
	var array1 [5]string

	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

	array1 = array2
	array2[1] = "White"
	fmt.Println(array1)
	fmt.Println(array2)

	//Assigning one array of three elements
	var arraystr1  [3]*string

	arraystr2 := [3] *string {new(string), new(string), new(string)}

	//Add colors to each element
    *arraystr2[0] = "Red"
    *arraystr2[1] = "Blue"
    *arraystr2[2] = "Green"

    //Copy the values from array2 into array1
   // arraystr1 = arraystr2
    fmt.Println(arraystr1)
    fmt.Println(arraystr2)

    //fmt.Println(&arraystr1)
    //fmt.Println(&arraystr2)

    //Declaring two-dimensional integer array of four elements by two elements
    var arraymultidimensional [4][2]int

    //Use an array literal to declare and initialize a two dimensional integer array
	arraymultidimensional = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41} }

	// Declare and initialize index 1 and 3 of the outer array.
	arraymultidimensional = [4][2]int{1:{20, 21}, 3:{40, 41}}

	//Declare and initialize individual elements of the outer and inner array
	arraymultidimensional = [4][2]int{1: {0:20}, 3:{1:41}}

	fmt.Println(arraymultidimensional)

	//Accessing elements of a two-dimensional array
	var arraymd1 [2][2]int
	arraymd1[0][0] = 100
	arraymd1[0][1] = 200
	arraymd1[1][0] = 300
	arraymd1[1][1] = 400

	fmt.Println(arraymd1)

	var arraymd2 [2]int = arraymd1[1]
	fmt.Println(arraymd2)

	var valuearr int = arraymd1[0][1]
	fmt.Println(valuearr)

	var arraymillion [1e6]int
	//Pass a pointer to the array and only copy eight bytes, instead of eight megabytes of memory on the stack
	//foo function takes a pointer to an array of one million elements of type integer. passes the address of the array
	// which requires eight bytes of memory to be allocated on the stack for the pointer variable
	foo(&arraymillion)
}
func foo(array *[1e6]int) {
	fmt.Println("Calling foo")
}