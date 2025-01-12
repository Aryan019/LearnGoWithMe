// The previous context here is that we are creating in a node server which is capable of hadling in the requests
// in the backend
// and here is the code of how we can send in the request in go laang to the backend server

// This is for the sending in of the get requests
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("Hitting in the requests in go lang ")
}

func performGetRequests() {
	url := "https://localhost:8080:"

	// Sending in the get requests
	// The below might give you response or error
	// http.Get(url)

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// It is my job to close in the requests as well
	defer response.Body.Close()

	fmt.Println("The response code is %T\n", response.StatusCode)

	// Reading in the response of getting in the data
	content, _ := io.ReadAll(response.Body)

	// After receivin in the response firstly use the io library to
	//read in it the response
	// and then convert and print in the string received
	fmt.Println(string(content))

	// Making in the post request with go lang
	// Assume in the endpoint we are sending in the data only accepts in the json format only

}
