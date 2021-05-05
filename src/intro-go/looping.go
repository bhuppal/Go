package main

import "fmt"

func main() {
	for i := 1; i <= 2; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 1; j <=i; j++ {
			fmt.Printf("\t\t The inner loop: %d\n", j)
		}
	}

	x := 83 % 40
	fmt.Println(x)

	y := 1
	for {
		y++
		if y > 25 {
			break
		}
		if y%2 != 0 {
			continue
		}
		fmt.Println(y)

	}
	fmt.Println("Done.")

	asciiKeyValues := 33
	for asciiKeyValues<=122 {
		fmt.Printf("%v \t %#x \t %#U\n",asciiKeyValues, asciiKeyValues, asciiKeyValues)
		asciiKeyValues++
	}

	str := []byte("Bhuppal")
	fmt.Println(str)

	str1 := []rune("Bhuppal") // just an alias name for int32
	fmt.Println(str1)

	str2 := []int32("Bhuppal")
	fmt.Println(str2)

	fmt.Println(true && true) // true
	fmt.Println(true && false) // false
	fmt.Println(true || true) // true
	fmt.Println(true || false) // true
	fmt.Println(!true) //false

}
