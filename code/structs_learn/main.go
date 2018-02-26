package main

import (
	"fmt"
)

type person struct { // a custom type
	firstName string
	lastName  string
}

func main() {
	// alex := new person{"Alex", "Anderson"}  // method one: struct property order matters

	alex := person{firstName: "Alex", lastName: "Anderson"}

	fmt.Println(alex.firstName)
	fmt.Println(alex.lastName)
}
