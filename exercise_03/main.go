package main

import (
	"fmt"
)

// Using Var dto declare the values
var x = 42
var y = "James Bond"
var z = true

func main() {
	// Using fmt.Sprintf to print all of the VALUES to one single string
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)

	// Printing f
	fmt.Println(s)

}
