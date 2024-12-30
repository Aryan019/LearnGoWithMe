package main

import "fmt"

const MyToken string = "abc"

// The above variable is now public in nature

func main() {
	// fmt.Println("Hey there my first own go structure")
	// Creating in the variables in Go

	var username string = "Aryan"
	fmt.Println("Hello", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)

	var anotherVariable int
	fmt.Println(anotherVariable)

	//Implicit way of declaring in variables

	// No var style
	numberOfUsers := 3000000
	// colon equal to operator is called walrus operator
	fmt.Println(numberOfUsers)

}
