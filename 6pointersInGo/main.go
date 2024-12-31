package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in Go")

	// Declaring in simple variables
	var a int = 10
	fmt.Println(a)

	// Declaring in the pointers now
	// var b *int
	// var b *string
	// The above is also allowed

	myNumber := 23
	fmt.Println(myNumber)

	// Storing in the location of myNumber variable
	var ptr = &myNumber

	// Address of myNumber is
	fmt.Println(ptr)
}
