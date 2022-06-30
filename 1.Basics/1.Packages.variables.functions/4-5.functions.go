// simple multiplication function
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
	fmt.Printf("Multiplying %d by %d is %d n", a, b, mp(a, b))
}
