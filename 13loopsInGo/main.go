package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in go lang ")

	// For looping through some data creating in the slice data structure so we
	// can easily loop it through

	days := []string{"Sunday", "Monday", "Tuesday"}

	// Printing in the slices
	fmt.Println(days)

	// Using in the for loop for iterating

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// Another type of for loop which is widely used
	// In js the i gives you in the entire element there but here it is not like that
	// here it only gives you the index and not the entire element there
	for i := range days {
		fmt.Println(days[i])

	}

	// Loop where the whole element is returned is called in as an for each loop
	// lets see it

	for j, value := range days {
		fmt.Printf("The index is %v and the value is %v/n", j, value)
	}

	rougeValue := 10

	// for value as an while loop
	for rougeValue < 10 {

		if rougeValue == 2 {
			goto tar
		}

		if rougeValue == 5 {
			rougeValue++

			continue
		}

		fmt.Println(rougeValue)
		rougeValue++

	}

	// The concept of go to statements in go
	// by using in the go to statements using in the label
tar:
	fmt.Println("I am declaring in my own label in go lang")

}
