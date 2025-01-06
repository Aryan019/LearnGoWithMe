package main

import "fmt"

func main() {
	fmt.Println("Rough work here of go")

}

// type defining here
type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Declaring in a slice to use it in as in-store data storage
var items []Item
