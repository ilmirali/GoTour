// explicit conversion
package main

import (
	"fmt"
)

func main() {
	v1 := 42
	v2 := 3.142
	v3 := "some text"
	fmt.Printf("v1 is of type %T\n", v1)
	fmt.Printf("v2 is of type %T\n", v2)
	fmt.Printf("v3 is of type %T\n", v3)
}
