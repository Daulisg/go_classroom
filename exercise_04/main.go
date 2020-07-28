package main

import (
	"fmt"
)

type numero int

var x numero

func main() {
	x = 10
	fmt.Printf("The TYPE for the X variable is: %T\n", x)
	fmt.Printf("The VALUE for the X variable is: %v\n", x)
	x = 42
	fmt.Printf("The new VALUE for the X variable is: %v\n", x)
}
