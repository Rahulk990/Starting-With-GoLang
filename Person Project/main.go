package main

import "fmt"

// Defining Struct
type name struct {
	firstname string
	lastname  string
}

// Nesting Struct
type person struct {
	name
	age int
}

func main() {

	// Creating Instances
	alex := person{name{"Alex", "Timber"}, 15}
	jack := person{
		name: name{
			firstname: "Jack",
			lastname:  "Sparrow",
		},
		age: 20,
	}
	fmt.Println(alex, jack)

	// Updating Struct and Format Printing
	jack.age = 40
	fmt.Printf("%+v", jack)

	// Using Receiver Functions
	alex.print()

	// Using Pointers
	jackPtr := &jack
	jackPtr.updateName("Jacky")
	jack.print()

	// Shortcut
	jack.updateName("Jackie")
	jack.print()
}

func (personPtr *person) updateName(name string) {
	(*personPtr).name.firstname = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
