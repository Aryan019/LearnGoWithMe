package main

import "fmt"

func main() {
	fmt.Println("Maps in Go lang")

	// Declaring in the maps
	// Thats how we declare in the maps in go lang
	languages := make(map[string]string)

	// Putting in the values in the map
	languages["JS"] = "Javascript"
	languages["Py"] = "Python"

	fmt.Println("Listing down all the languages in the map ", languages)
	//Accessing a key-value pair in the map by using in the key
	fmt.Println("JS stands for ", languages["JS"])

	delete(languages, "Py")
	fmt.Println("All the languages here are ", languages)

	// Starting with control flow over maps in go
	// for loop in go

	for key, value := range languages {
		fmt.Printf("For the key %v the value is %v\n", key, value)
	}
}
