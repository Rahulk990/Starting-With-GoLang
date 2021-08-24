package main

import "fmt"

func main() {
	// Printing String
	fmt.Println("Hello World!")

	// Declaring Variables
	var item string = "Item"
	item2 := "Item2"
	fmt.Println(item, item2)

	// Declaring Arrays/Slices
	items := []string{"I1", "I2", "I3", "I4", "I5"}
	for i, item := range items {
		fmt.Println(i, item)
	}

	// Indexing Arrays/Slices
	fmt.Println(items[0])
	fmt.Println(items[0:2])

	// Type Conversions
	str := "Hello World!"
	fmt.Println([]byte(str))

}
