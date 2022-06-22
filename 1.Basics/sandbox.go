package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Year:", t.Year())
	fmt.Println("UnixNano:", time.Now().UnixNano())	// 1655282909309318659
	fmt.Println("RFC3339", t.Format(time.RFC3339))	// 2022-06-15T15:13:38+06:00
	fmt.Println(t.Format("15:04:05"))
}	
