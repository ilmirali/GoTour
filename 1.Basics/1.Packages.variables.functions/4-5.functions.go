// simple multiplication function with output of variables types
package main

import (
	"fmt"
)

func mp(x, y int) int {
	return x * y
}

func main() {
	a := 42
	b := 13
	fmt.Printf("Multiplying %d by %d is %d n\n", a, b, mp(a, b))
	fmt.Printf("The type of a is: %T\n", a)
	fmt.Printf("The type of b is: %T\n", b)
}
