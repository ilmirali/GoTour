// switch, runtime package using
// cmd - go tool dist list - to get full list of options
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd
		// plan9, windows
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("runtime.GOOS", runtime.GOOS)
//	fmt.Println("runtime.GOOS", runtime.GOGC)
	fmt.Println("runtime.GOMAXPROCS", runtime.GOMAXPROCS)
	fmt.Println("runtime.NumCPU", runtime.NumCPU)
	fmt.Println("runtime.ReadMemStats", runtime.ReadMemStats)
}
