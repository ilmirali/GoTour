// Clean switch
//  Switch without a condition is the same as switch true. 
package main

import ("fmt"; "time")

func main() {
	n := time.Now().UnixNano()
	t := time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good day!")
	default:
		fmt.Println("Good afternoon!")	
	}
	fmt.Println("time.Now().UnixNano():", n)
}
