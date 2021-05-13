package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 = 5 	// type inferred
	var num2 = 6	// explicitly typed
	var rates float32 = 4.5 // declare as float32 and initialize

	var raining bool = false // declare as bool and initialize

	num3 := 7 //declaring and init
	num4 := num3

	var num5, num6 int = 8, 9 // multiple declares and assignment

	//declare and intialize multiple variables
	var (
		age = 25
		name = "Bhuppal"
	)

	// String interpolation
	var num7 = 5
	var rates1 float32 = 4.5

	fmt.Println("num1 is " + strconv.Itoa(num1) + " and rates is " 	+ strconv.FormatFloat(float64(rates1), 'f', 2, 32))

	//Converting numeric/boolean values to strings
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-42, 16)
	s4 := strconv.FormatUint(42, 16)

	//strings to numeric
	//bool
	if b, err := strconv.ParseBool("true"); err == nil {
		fmt.Printf("%T, %v\n", b, b)
	}
    //float
    if f, err := strconv.ParseFloat("3.1415", 64); err == nil {
    	fmt.Printf("%T, %v\n", f,f)
	}
	//Int
	if i, err := strconv.ParseInt("-42", 10, 64); err == nil {
		fmt.Printf("%T, %v\n", i, i)
	}
	//Uint
	if u, err := strconv.ParseUint("42", 10, 64); err == nil {
		fmt.Printf("%T %v\n", u, u)
	}

	// An easy way to combine string and numeric values is to use the Sprintf() function
	// with the various format specifiers
	str := fmt.Sprintf("num1 is %d and rates is %.2f", num1, rates)
	fmt.Println(str)


	fmt.Println(num2, rates, raining, num4, num5, num6, age, name, num7, rates1, s1, s2, s3, s4)
}