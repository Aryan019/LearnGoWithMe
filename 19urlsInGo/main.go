package main

import (
	"fmt"
	"net/url"
)

func main() {
	myUrl := "https://google.com"
	fmt.Println(myUrl)
	fmt.Println("Urls in go lang ")

	// Not fully completed have some dependencies issues in the go

	// parsing an url means extracting out some data from it
	// like http library was used in to sent in the requests to the server

	// we have a package called url which is used to parse the url

	// the parsed url statement has two much to be identified and explored
	result, _ := url.Parse(myUrl)

	// Extracting out the info from the result
	// after parsing it into the result
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// to store in the rawquery params there is a better way to do it
	qparams := result.Query()

	// The above lines converts in the qparams or the query params to the key value pair when
	// performed in the url parsing
	fmt.Println(qparams["coursename"])

	// Applying in the for loop and looping around the qparams
	for _, val := range qparams {
		fmt.Println("Params is ", val)

	}

	// you have all the info and parts and you want to create in the url from it
	// & describes in that you are passing in the url in this case
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "google.com",
		Path:   "/search",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("The another url is ", anotherUrl)

}
