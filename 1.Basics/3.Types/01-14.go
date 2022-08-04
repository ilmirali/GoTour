package main

import (
	"fmt"
)

// 0
// Pointers with dereferencing (or "indirecting"). 

func simple_point() {
	var p *int
	i := 42
	p = &i
	fmt.Println("i:", i, "p:", *p)
	*p = 21
	fmt.Println("i:", i, "p:", *p)
}

// 1
func pointers_basics() {
	i, j := 42, 2701
	var p *int
	
	fmt.Println(p)
	p = &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(*p, i)
	
	p = &j
	fmt.Println(*p)
	*p = *p / 37 
	fmt.Println(*p)

}

// 2
// A struct is a collection of fields
type Vertex struct {
	X, Y int
}

// 3
// Struct fields are accessed using a dot
func Vertex3() {
	v := Vertex {1, 2}
	v.X = 4
	fmt.Println("v.X", v.X)
}

func main() {
	// 1
	fmt.Println("1.Pointers with dereferencing (or \"indirecting\")")
	simple_point()
	pointers_basics()
	// 2
	fmt.Println("2. A struct is a collection of fields")
	fmt.Println(Vertex{}, Vertex{1, 2})
	// 3
	fmt.Println("3. Struct fields are accessed using a dot")
	Vertex3()
	
}
