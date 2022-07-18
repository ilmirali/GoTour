// just defer
package main

import "fmt"

func main() {
	defer fmt.Printf ("world!\n")
	fmt.Printf ("Hello, ")	
}
