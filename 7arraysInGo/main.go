package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in Go")

	var fruitList [4]string // Declaring in the array

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Peach"
	fruitList[3] = "Banana"

	// Another way to declare in the list is as follows

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println(vegList)

}
