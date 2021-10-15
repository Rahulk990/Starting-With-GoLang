package main

import "fmt"

func main() {

	// Defining Map
	colors := map[string]string{
		"red":    "apple",
		"yellow": "banana",
	}
	fmt.Println(colors)

	// Adding Values
	colors["orange"] = "orange"
	fmt.Println(colors)

	// Deleting Values
	delete(colors, "red")
	fmt.Println(colors)

	// Iterating Map
	for k, v := range colors {
		fmt.Println(k, v)
	}
}
