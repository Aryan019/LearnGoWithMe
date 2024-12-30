package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Enter your name")
	// var name string
	// fmt.Scanln(&name)
	// fmt.Println("Hello", name)

	welcome := "Welcome to the world of Go"
	fmt.Println(welcome)

	// Sample and simply for the testing purpose it is for
	fmt.Println("Enter your name")
	reader := bufio.NewReader(os.Stdin)

	// Comma ok or reader ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Hello", input)

}
