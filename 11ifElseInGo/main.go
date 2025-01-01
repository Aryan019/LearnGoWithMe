package main

import "fmt"

func main() {
	// if and else in go
	var age int = 21
	fmt.Println(age)

	// learning in the if else here
	var loginCount int = 1

	if loginCount > 10 {
		fmt.Println("Medium heavy user")
	} else if loginCount < 5 {
		fmt.Println("Light heavy user")
	} else {
		fmt.Println("Normal user")
	}

	// Another method for using in the if and else statement
	if 9%2 == 0 {
		fmt.Println("No there modulus operation dont result in 0")
	} else {
		fmt.Println("Yes It do")
	}

	// The below method will be utilized a lot in web handling
	// Special syntax of if else

	// Declaring the variable on the go and checking it via applying in some conditions
	// Putting in the semicolon after variable declaration is must
	if num := 3; num < 10 {
		fmt.Println("Condition is satisfied ")
	} else {
		fmt.Println("Condition is not satisfied ")
	}

	// var err string;

	// if err!= nil{
	// 	fmt.Println("")
	// }
}
