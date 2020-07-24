package main

import (
	"fmt"
)

// Package level scope variables with the zero value assignment
var x int
var y string
var z bool

func main() {

	// Printing the variable types
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
