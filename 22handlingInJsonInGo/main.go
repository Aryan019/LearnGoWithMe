package main

import (
	"encoding/json"
	"fmt"
)

// Creating in the json data in go lang

// We were reciving in the json data but not treating it as in the json data structure
// we were treating it more like a string not a full flesh json strin

// Creating and handling in json requests response while it is coming from the backend or we
// are sending in a json response

// Making in the structure using struct
type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

// Working with the encoding and decoding in the json

// Encoding in the json

func main() {
	fmt.Println("Hare Krishna")
}

// Making in data through the struct
func encodeJson() {
	differentCourses := []course{
		{"Go", 100, "Aryan", "1234", []string{"Backend", "Frontend"}},
		{"Go", 100, "Aryan", "1234", []string{"Backend", "Frontend"}},
		{"Go", 100, "Aryan", "1234", []string{"Backend", "Frontend"}},
	}

	// Packaging in the above data as the json data

	// The library name is json only which is used to pack in the data in the json format
	// marshal is kind of the way to implement of the json

	// While marshalling we are sending in the data as well and also an interface
	// Interface is nothing but a different alternative for struct
	finalJson, err := json.Marshal(differentCourses)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
	// Marshal indent is nothing but organizing the data in the json format
	// with the help of the indent we can make in the json data more readable

	// See in the notion notes aryan
	// Below is the example of the marshal indent

	// finalJson, err := json.MarshalIndent(differentCourses, "", "\t")

	// so basically first declare in rhe struct and then put in the data into it and then by using in the json marshal
	// method convert to json and then send in the data also the - and omitempty are also a thing here
}
