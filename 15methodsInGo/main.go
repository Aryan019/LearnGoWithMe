package main

import "fmt"

func main() {

	aryan := User{"Aryan", "aryan.dev@gm", 21, "GoDev"}
	fmt.Println(aryan)

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status string
}

// The below is a simple function and not a method it is simply a simple function which is defined in the
// go lang
func getStatus() {

}

// The G should be capital in the below because we have to export it so it must be capital letter
// Now defining in the method
func (u User) GetStatusMethod() {
	fmt.Println("Is user active here ", u.Status)

}
