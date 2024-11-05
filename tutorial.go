package main

import (
	"fmt"
)

func test(x int) {
	fmt.Println("Test:", +x)

}

//	func add(x int, y int) (int, int) {
//		// fmt.Println("Add:", x+y)
//		return x + y, x - y
//		// return
//	}
func add(x, y int) (z1, z2 int) {
	defer fmt.Println("hello")
	z1 = x + y
	z2 = x - y
	fmt.Println("Before returning")
	return

}
func main() {
	test(4)
	test(7)
	// add(4, 9)

	ans1, ans2 := add(5, 6)
	fmt.Println(ans1, ans2)

}

/*

The go build command compiles the Go code and produces an executable file (tutorial.exe on Windows). It does not run the program; it only creates an executable you can run later.
This step is useful if you want to compile the program once and run it multiple times without having to recompile each time.



*/
