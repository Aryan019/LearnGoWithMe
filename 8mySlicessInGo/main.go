package main

import "fmt"

func main() {
	fmt.Println("playing with slices in go language")

	// Declaring in the slices
	var fruitList = []string{"apple", "orange"}

	fmt.Println(fruitList)

	fruitList = append(fruitList, "banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// Making in the slices with different approach
	// highScores:= make([]int,4)

	// The strange case of go on memory allocation and reallocation
	// check in with the notes

	var courses = []string{"reactJs", "Angular", "VueJs"}
	fmt.Println(courses)

	var index int = 2
	fmt.Println(index)

	// Removing in the values from the slices
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

	// more()

}
