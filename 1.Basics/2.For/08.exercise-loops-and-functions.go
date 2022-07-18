// sqare root function (z -= (z*z - x) / (2*z))
package main

import (
	"fmt"
)

func Sqrt1(x float64) float64 {
	z := float64(1) // z := 1.0
	for i:= 0 ; i < 10; i++ {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func Sqrt2(x float64) float64 {
	z := float64(1) // z := 1.0
	j := 10
	for i:= 0 ; i < j; i++ {
		if z*z == x {
			return z
			} else {
				j += 10
			}
		z -= (z*z - x) / (2*z)
	}
	return z
}


func main() {
	fmt.Println(Sqrt1(8100))
	fmt.Println(Sqrt2(25000000))
}
