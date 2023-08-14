package main

import (
	"fmt"
)

func main() {

	// Three wWays to declare a map. Both ways below will initialize with no values inside

	// 3˚ Built in function
	// colors := make(map[int]string)

	// ---------------------------------------

	// 2˚
	// var colors map[string]string

	// ---------------------------------------
	/* 1˚
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
	}
	*/

	// Add values to a map - Key and Value
	// colors[10] = "#ffffff"

	// Delete keys and values on an existing map. Built in function
	// delete(colors, 10)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"white": "#ffffff",
	}

	printMap(colors)
}

// Iterate through the map and print the keys and values
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}

}
