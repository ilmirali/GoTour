// var statement (with initializers) and short assignment statement (:=)
// Outside a function, every statement begins with a keyword (var, func, and so on)
//  and so the := construct is not available
package main

import (
	"fmt"
)

var i, j int = 1, 2

func main() {
	var c, py, java = true, false, "no!"
	k := 3
	a, b := 23, 45
	fmt.Println(i, j, c, py, java, k, a, b)
}
