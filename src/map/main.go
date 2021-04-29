package main

import (
	"fmt"
)

func main() {

	//var colors map[string]string;
	/*
	       colors := make(map[string]string)
	   	colors["white"] = "#ffffff"
	   	fmt.Println(colors);

	   	delete(colors, "white")
	   	fmt.Println(colors)
	*/
	//fmt.Println("Hello, World!!!")
	/*
		colors := map[string]string{
			"red": "#ff0000",
			"green": "#4bf745",
		}

		fmt.Println(colors)
	*/
	colors := map[string]string{
		"white": "#ffffff",
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors)
	printMap(colors)

	// Declaring a map using make
	// Create a map with a key of type string and a value of type int.
	//dict := make(map[string]int)

	// Create a map with a key and value of type string
	// Initialize the map with 2 key/value pairs
	dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println(dict)
	//Delcaring a map that stores slices of strings
	// Create a map using a slice of strings as the value
	//dict := map[int]string{}

	// Assigning values to a map
	// Creating an empty map to store colors and their color codes
	colors1 := map[string]string{}

	// Add the Red color code to the map
	colors1["Red"] = "#da1337"
	fmt.Println(colors1)

	value, exits := colors["red"]
	if exits {
		fmt.Printf("my color is %v\n", value)
	}

	value1 := colors["green"]

	if value1 != "" {
		fmt.Println(value1)
	} else {
		fmt.Println(value1)
	}

	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	delete(colors, "green")

	fmt.Println("***After delete the value - green ***")

	for key, value := range colors {
		fmt.Printf("Key: %s value: %s\n", key, value)
	}
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
