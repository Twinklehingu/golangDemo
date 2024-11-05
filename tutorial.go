package main

import (
	"fmt"
) // used for displaying/print

// entry point
func main() {

	var mp map[string]int = map[string]int{
		"apple":  1,
		"pear":   2,
		"banana": 3,
	}

	// val, ok := mp["apple"]
	// fmt.Println(val, ok)
	val, ok := mp["tim"]
	fmt.Println(val, ok)

	// mps := make(map[string]int)
	// fmt.Println(mp["apple"])
	// mp["orange"] = 100
	// delete(mp, "orange")
	fmt.Println(mp)

}

/*

The go build command compiles the Go code and produces an executable file (tutorial.exe on Windows). It does not run the program; it only creates an executable you can run later.
This step is useful if you want to compile the program once and run it multiple times without having to recompile each time.



*/
