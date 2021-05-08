
package main

import "fmt"

func addNums(x int64, y int64) int64 {
	return x + y
}

func main() {
	// if-else
	if true {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// ***** There is no ternary in Go ****

	// "if" statement allows you to have two expressions:
	// one assignment and one condition

	limit := 10
	result := addNums(10, 20)
	sum := result
	if  result <= limit {
		fmt.Println(sum)
	} else {
		fmt.Println(limit)
	}
//	result := addNums(10, 20)
	fmt.Println(result)
}


