package main

import (
	"fmt"
) // used for displaying/print

// entry point
func main() {
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)

	nums := [3]int{5, 62, 2}
	// fmt.Println(len(nums))

	sum := 0

	arr2D := [2][3]int{{5, 62, 4}, {4, 2, 1}}
	fmt.Println(arr2D)

	for i := 0; i <= len(nums)-1; i++ {
		sum += nums[i]
	}
	fmt.Println(sum)

}

/*

The go build command compiles the Go code and produces an executable file (tutorial.exe on Windows). It does not run the program; it only creates an executable you can run later.
This step is useful if you want to compile the program once and run it multiple times without having to recompile each time.



*/
