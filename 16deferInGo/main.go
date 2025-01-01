package main

import "fmt"

// Refer to the notes aryan
// What will be happening (just think bro )
func main() {
	defer fmt.Println("World of Go One")
	defer fmt.Println("World of Go Two")
	defer fmt.Println("World of Go Three")
	fmt.Println("Hello")
	myDefer()

}

func myDefer() {

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
