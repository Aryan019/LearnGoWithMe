package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to my modules in Go")
	gretter()

	// Creating in a router in gorilla mux
	// Creating in a new router
	//It is optional but also specifying in the type of method that will enter in
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	// Running in my first server on golang
	// log fatal for logging in the different files and all
	log.Fatal(http.ListenAndServe(":8080", r))

}

func gretter() {
	fmt.Println("Hey there mod users")
}

// Another method
// Notes here
// Common syntax used mostly
//The request related things and all are inside the r
// the params and all are inside of the r

// But if you want to send in some response back to the client
// that is through the w the response writer object
func serverHome(w http.ResponseWriter, r *http.Request) {

	// Sending in a response and as everything is in byte
	// sending in the response in the byte format

	w.Write([]byte("<h1>Welcome to go lang aryan</h1>"))

}
