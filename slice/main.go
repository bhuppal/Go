package main

import "fmt"

func main() {
	fmt.Println("Slice concepts")

	//Create an array of three integers
	array100 := [3]int{100, 200, 300}
	fmt.Println(array100)

	//Create a slice of integers with a length and capacity of three
	slice100 := []int{10, 20, 30}
	fmt.Println(slice100)

	// Create a slice of strings
	// If we specify value inside [], then its array otherwise slice
	// Contains a length and capacity of 5 elements
	slice := make([]string, 3,5)
	slice[1] ="Bhuppal"
	fmt.Printf("The length of the slice is %d\n", len(slice))
	fmt.Printf("The capacity of the slice is %d\n", cap(slice))

	//slice1 := make([]int, 5, 3)
	//Compiler error: len larger than cap in make([]int)

	//Declaring a slice with a slice literal
	//Create a slice of strings
	//contains a length and capacity of 5 elemets
	slice1 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println(slice1)

	// Create a slice of integers
	// Contains a length and capacity of 3 elements
	slice2 := []int{10, 20, 30}
	fmt.Println(slice2)

	//Declaring a slice with index positions
	// Create a slice of strings
	// Initialize the 100th element with an empty string
	slice3 := []string{99:"bhuppal"}
	fmt.Println(slice3[99])

	var slicenothing []int
	fmt.Println(slicenothing) // nil pointer, len 0, cap 0

	slicezero := make([]int, 0) //address contains value pointer, len 0, cal 0
	fmt.Println(slicezero)

	slicezero1 := []int{} // memonry address contains value pointer, len 0 , cal 0
	fmt.Println(slicezero1)

	sliceexample := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}
	fmt.Println(len(sliceexample), cap(sliceexample))

	newSlice := sliceexample[3:5]
	fmt.Println(newSlice)
	fmt.Println(len(newSlice), cap(newSlice))
	// slice[i,j] underlying array of capacity k
	// Length: j - 1
	// Capacity: k - 1

	//underlying slice will get updated if the newslice is changed
	newSlice[1] = 90
	fmt.Println(sliceexample, newSlice)

	//Using append to add an element to a slice
	newSlice = append(newSlice, 60)
//	newSlice = append(newSlice, 100)
//	newSlice = append(newSlice, 200)
//	newSlice = append(newSlice, 300)
//	newSlice = append(newSlice, 400)
	fmt.Println(cap(newSlice))
	fmt.Println(newSlice, cap(newSlice))
	fmt.Println(sliceexample)
	newSlice[0] = 900
	newSlice[1] = 900
	newSlice[2] = 900
	//newSlice[3] = 900
	//newSlice[4] = 900
	//newSlice[5] = 900
	//newSlice[6] = 900
	//newSlice[7] = 900
	//newSlice[8] = 900
	//newSlice[9] = 900
	fmt.Println(newSlice)
	fmt.Println(sliceexample)


	// Create a slice of integers
	// Contains a length and capacity of 4 elements
	sliceEx := []int{10, 20, 30, 40}

	// Append a new value to the slice
	// Assign the value of 50 to the new element
	newSliceEx := append(sliceEx, 50)
	//newSliceEx = append(sliceEx, 80)
	fmt.Println(sliceEx, cap(sliceEx))
	fmt.Println(newSliceEx, cap(newSliceEx))

	// Create a slice of strings
	// Contains a length and capacity of 5 elements
	source := []string{"Apple", "Orange", "Plum", "Grapes", "Banana"}
	fmt.Println(source, cap(source), len(source))

	newSliceSource := source[0:len(source) -2:len(source)]
	// For slice[i:j:k] or [2:3:4]
	// Length: j - i or 3 - 2 = 1
	// Capacity: K - i or 4 - 2 = 2
	fmt.Println(newSliceSource, len(newSliceSource), cap(newSliceSource))

	// Benefits of setting length and capacity to be the same
	// Create a slice of strings
	// Contains a length and capacity of 5 elements
	sourceBenefit := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	// slice the third element and restrict the capacity
	// Contains a length and capacity of 1 element
	newSourceBenefit := sourceBenefit[2:3:3]
	//newSourceBenefit = append(newSourceBenefit, "Kiwi")
	fmt.Println(sourceBenefit, len(sourceBenefit), cap(sourceBenefit))
	fmt.Println(newSourceBenefit, len(newSourceBenefit), cap(newSourceBenefit))
    newSourceBenefit[0] = "Cat"
    //newSourceBenefit[1] = "Dog"
	fmt.Println(sourceBenefit, len(sourceBenefit), cap(sourceBenefit))
	fmt.Println(newSourceBenefit, len(newSourceBenefit), cap(newSourceBenefit))


	// Appending to a slice from another slice
	s1 := []int{1, 2, 3, 4}
	s2 := []int{3, 4, 5, 6}
	fmt.Printf("%v\n", append(s1, s2...))

	// Iterating over Slice
	sliceLoop := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	for index, monthName := range sliceLoop{
		fmt.Printf("Month name is %v and number is %d\n",monthName, index+1)
	}

	// Declaring a multidimensioanl slice
	// Create a slice of a slice of integers
	sliceCompose1 := [][]int{{10}, {100, 200}, {20}, {80, 900}}
	fmt.Println(sliceCompose1)
	// Append the value of 20 to the first slice of integers
	sliceCompose1[0] = append(sliceCompose1[0], 20)
	fmt.Println(sliceCompose1)

	sliceCompose1[2] = append(sliceCompose1[0], sliceCompose1[1]...)
	fmt.Println(sliceCompose1)

	fmt.Println(sliceCompose1[3][0])

	sliceC1 := [][]int{{10}, {100, 200}}
	sliceC1[0] = append(sliceC1[0], 20)
	fmt.Println(sliceC1)

	//Passing slices between functions
	// Allocate a slice of 1 million integers
	sliceFun := make([]int, 1e6)

	// Pass the slice to the function too
	// On 64-bit architecture, a slice requires 24 bytes of memory.
	// The pointer field requires 8 bytes, and the length and capacity fields requires 8 bytes respectively
	// Passing the 24 byres between functions is fast and easy
	sliceFun = foo(sliceFun)

}

func foo(slice []int) []int {
	fmt.Println("Slice - func foo")
	return slice
}


















