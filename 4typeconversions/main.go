package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Aryan shopping mall")
	fmt.Println("Rate the mall experience from 1-5")

	// I have to take in input from the user here
	// so we will be in the need of the input package here

	// To take in input initializing a  reader here

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)

	// numRating = input + 1;
	// The above is the thing we are trying to achieve in

	// See the notes once
	numRating, error := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// The first para it takes in is the input and the second it takes in is the bit size

	// The below code means there is some kind of error in it
	if error != nil {
		fmt.Println(error)

	} else {

		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
