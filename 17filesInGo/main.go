package main

import (
	"fmt"
	"io"
	"os"
)

// We have some content that we want to put it in the text file
// and handling on those files in go lang
func main() {

	fmt.Println("Files in go lang")

	// content of the file is defined below
	content := "This is the content of the file and if you are taking in reference from my github repo welcome to you my friend keep exploring :-x>"

	// Creating in the file with the help of the os module
	// and in that we have to specify the file location where the file will be created in

	file, err := os.Create("./mySampleTextFileFromOs.txt")

	if err != nil {
		// panic is used to pass in the error
		fmt.Println("There are some errors in creating in the file (FILE CANNOT BE CREATED )")
		panic(err)

	}

	// handling in the file creation and error with the comma special syntax in go
	length, err := io.WriteString(file, content)
	if err != nil {
		// panic is used to pass in the error
		fmt.Println("There are some errors in creating in the file (FILE CANNOT BE CREATED )")
		panic(err)

	}

	fmt.Printf("Length is %v", length)
	// After the creation of file closing in the file in go lang

	// Using in defer to close in the file as it must be executed at the very very last
	defer file.Close()
}
