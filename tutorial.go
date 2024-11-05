package main

import (
	"fmt"
) // used for displaying/print

// entry point
func main() {
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = x[1:3]    //: all entries, 1: meaning start from one to the end, :4 meaning start from zero to 4th index
	fmt.Println(s[:cap(s)]) // len is 2 and cap is 4

	var a []int = []int{5, 6, 7, 8, 9}
	// fmt.Println(cap(a[:3]))
	b := append(a, 10)
	fmt.Println(b)

	y := make([]int, 5)
	fmt.Println(y)

}

/*

The go build command compiles the Go code and produces an executable file (tutorial.exe on Windows). It does not run the program; it only creates an executable you can run later.
This step is useful if you want to compile the program once and run it multiple times without having to recompile each time.



*/
