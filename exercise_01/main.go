package main

import (
	"fmt"
)

func main() {
	// Shorthand variable assignment
	X := 42
	Y := "Sames Bond"
	Z := true

	// Single print statement to print all valu
	fmt.Printf("X = %v Y = %v and Z = %v\n", X, Y, Z)

	// Printing the type and value with separate print statements
	fmt.Printf("The var X is of type: %T and value is qual to: %v\n", X, X)
	fmt.Printf("The var Y is of type: %T and value is qual to: %v\n", Y, Y)
	fmt.Printf("The var Z is of type: %T and value is qual to: %v\n", Z, Z)

}
