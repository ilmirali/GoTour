// "naked" return
package main

import (
	"fmt"
)

func split(sum int) (x, y, z int) {
	x = sum * 4 / 9
	y = sum - x
	z = 0
	return
}

func main() {
	n := 17
	fmt.Println(split(n))
}
