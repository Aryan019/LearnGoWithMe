// The previous context here is that we are creating in a node server which is capable of hadling in the requests
// in the backend
// and here is the code of how we can send in the request in go laang to the backend server

// This is for the sending in of the get requests
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("Hitting in the requests in go lang ")
	performGetRequests()
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

func performPostRequestsWithJson() {
	// The url where you want to send in the post request

	const myurl = "<enter the url here>"

	//  Creating in the fake json payload
	// after applying back ticks create in any kind of data
	requestBody := strings.NewReader(`
	"name" : "Aryan019",
	"place": "Indore"
	`)

	// Actually sending in the post request
	// While sending in post req three types of data to be included in
	// the url you are sending it to
	// the type of data you are sending in
	// and the request body
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	//Reading in the response
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

	// Sending in the form data in go lang
	// Useful in the case of image upload and all

	//Performing in the post requests of the form

}

// Sending in the form data as in the request
func performFormRequest() {
	const myUrl = "<Some random url in which you can send in the form data"
	// Creating in the fake form data

	// This url package is used to create in the form data
	data := url.Values{}

	data.Add("name", "Aryan019")
	data.Add("place", "Indore")

	// PostForm method is used to sending in the form data request in go
	// Also handling in the response in here
	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)

	}

	// Reading in the response
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
