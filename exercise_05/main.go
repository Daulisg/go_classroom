package main

import (
	"fmt"
)

type numero int

var x numero
var y numero

func main() {
	var y = 42
	fmt.Printf("The TYPE for the variable numero is: %T\n", y)
	fmt.Printf("The VALUE for the variable Y is: %v\n", y)
	var x = int(y)
	fmt.Printf("Now, The TYPE for the variable X is: (%T) but before was of TYPE numero\n", x)
}
