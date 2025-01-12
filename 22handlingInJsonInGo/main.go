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

// Part 2 consuming in or decoding in the json data

func decodeJson() {
	// the data that comes in the go
	// It is by default in the bytes format
	// So we need to convert it to the string
	// and then send it to the json unmarshal method

	jsonDataFromWeb := []byte(
		`{
		"name" : "Aryan",
		"price" : 100,
		"platform" : "Aryan",
		"password" : "1234",
		"tags" : ["Backend", "Frontend"]
		
		
		}
	`)

	// Checking in the data is valid or not
	// json.valid is nothing but checking in the data is valid or not
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("The json data is valid")

		// Passing in the actual data in json
		json.Unmarshal(jsonDataFromWeb, &jsonDataFromWeb)

		// to print in the data from this
		// we use in
		// The below line is used for the pretty print
		fmt.Printf("%#v\n", jsonDataFromWeb)
	}

	//add in the data using the key value pairs in
	// basically converting in the data of json to the map
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	// makes it a map from the json data
	fmt.Printf("%#v\n", myOnlineData)

	// Order is not guaranteed in maps

	for k, v := range myOnlineData {
		fmt.Printf("The key is %v and the value is %v\n", k, v)
	}
}
