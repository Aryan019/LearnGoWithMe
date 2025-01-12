package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Creating in models for our file
// The model should go in a file
// Remember to import in the file

type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"course_price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"full_name"`
	Website  string `json:"website"`
}

// main is the entry point of the application.
// Creating in the fake db here

var courses []Course

// The middlewares and the helpers in go -- also in a separate file
// This method must be part of that structure so let us pass that structure here in

// Attaching this method to my struct
func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""

}

func main() {

}

// Controllers in go lang -- also in a separate file
// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to go lang aryan It is my home page   </h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all the courses")

	// Setting up the headers of the file
	// Writing up the headers of the file

	w.Header().Set("Content-Type", "application/json")

	// Sending in the response with the help of the response writer
	json.NewEncoder(w).Encode(courses)
}
