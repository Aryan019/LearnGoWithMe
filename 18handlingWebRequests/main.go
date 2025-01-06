package main

import (
	"fmt"
	"io"
	"net/http"
)

// See the notes once
func main() {

	url := "https://google.com"
	fmt.Println(url)
	// http package will be used to hit a get requests to an entered website

	fmt.Println("Web Requests in Go ")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The response is %T\n", response)

	// and the response body must must be closed
	defer response.Body.Close()

	// Reading in that response
	// io readall helps read the response and convert it to the databytes
	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	// Converting it to the string
	content := string(databytes)
	fmt.Println(content)
}
