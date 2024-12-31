package main

import "fmt"

func main() {
	fmt.Println("Structs in go lang ")
	aryan := User{
		Name:   "Aryan",
		Email:  "aryan@go.dev",
		Status: true,
		Age:    22,
	}

	fmt.Println(aryan)
	fmt.Printf("Here are the detail of the user his name is %v and email is %v", aryan.Name, aryan.Email)

}

// Out of the main function
// Below the user has capital U for User
// So it simply means that it is a public class behavior imitation
// Also same it is for N in Name E in Email

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
