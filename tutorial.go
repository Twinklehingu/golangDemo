package main

import (
	"fmt"
) // used for displaying/print

// entry point
func main() {
	var a []int = []int{4, 5, 32, 113, 45, 24, 245, 6, 65, 6}
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }
	//i is index and element numbere a[i]
	for i, element := range a {
		// fmt.Printf("%d: %d\n", i, element)

		for j, element2 := range a {
			if element2 == element && i != j && i < j {
				fmt.Println(element)
				// fmt.Println(i)
				// fmt.Println(j)
			}
		}
	}
	// for _, element := range a {
	// 	fmt.Printf("%d\n", element)

	// }

}

/*

The go build command compiles the Go code and produces an executable file (tutorial.exe on Windows). It does not run the program; it only creates an executable you can run later.
This step is useful if you want to compile the program once and run it multiple times without having to recompile each time.



*/
