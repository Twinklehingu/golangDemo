package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
) // used for displaying/print

// entry point
func main() {
	// var name string = "Hello Twinkle!"
	// var number uint16 = 255
	// number = number + 34
	// var name string
	// name = "Hello Twinkle!" // assign value to variable
	// name = "Bill"
	// fmt.Println("Hello, world!") //print new line for printing "" and '' to show characters
	// fmt.Println(name, number)

	// var number = 20000.98
	// num := "hello"
	// fmt.Printf("%T", num)

	// var numbers uint64
	// var bl bool
	// fmt.Println(numbers, bl)

	// fmt.Printf("%T", 10) // will show the type which is int
	// fmt.Printf("Hello %t", 0)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year you were born: ")
	scanner.Scan()

	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d years old at the end of 2024", 2024-input)
}

/*

The go build command compiles the Go code and produces an executable file (tutorial.exe on Windows). It does not run the program; it only creates an executable you can run later.
This step is useful if you want to compile the program once and run it multiple times without having to recompile each time.



*/
