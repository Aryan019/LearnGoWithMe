package main

import "fmt"

func main() {
	// In the main function curly braces
	// we can pass on some params
	// No issues in that
	fmt.Println("functions in go lang ")
	greeter()

	sum := adder(2, 4)
	fmt.Println("The sum is ", sum)
}

func greeter() {
	fmt.Println("Good afternoon Aryan ")
}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

// See the notes once
