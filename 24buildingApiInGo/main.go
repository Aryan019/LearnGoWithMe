package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	// off the id check while using in the create one course method
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

// Get one course
func getOneCourse(w http.ResponseWriter, r *http.Request) {

	//Setting up the headers of the file
	w.Header().Set("Content-Type", "application/json")

	// Grabbing in the course id
	//gorilla mux router also provides us the way to grab in the params here
	params := mux.Vars(r)

	// Looping in the courses and finding in the course with that particular id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			// returning in the function
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with this id")
	return

}

// Performing in a create operation in go lang

func createOneCourse(w http.ResponseWriter, r *http.Request) {

	// json is coming in in this function
	// so this time we are creating in the data through in the request json object and body

	// So firstly we have to decode in the json here
	fmt.Println("Create one course")
	// Setting up the headers of the file
	w.Header().Set("Content-Type", "application/json")

	// What if the body received is nil and is not defined ;
	// r.body gives us the request body to us
	// That body can be used to find in some attributes or have some validation check on the data
	if r.Body == nil {
		// Sending in a json response back to the client that he is entering in the empty one here
		// for sending in the data in the json back we use in some newEncoder and then in
		// encoder we pass in the writer object and then we pass in the data
		json.NewEncoder(w).Encode("Please send in some data")
	}

	// what about when data is {}

	// Creating in a variable called course of Course type here
	var course Course
	//While decoding in we utilize in the reader here
	// and we are decoding in the request body here
	// so see it below

	//When utilizing in a encoder or decodder we firstly create in teh
	// newEncoder and decoder and then apply in the encode and decode method to it

	// In decode method we specify in where to receive in the decoded data
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		// Sending in a json response back to the client that he is entering in the empty one here
		// for sending in the data in the json back we use in some newEncoder and then in
		// encoder we pass in the writer object and then we pass in the data
		json.NewEncoder(w).Encode("No data inside the json ")
		return
	}

	// The decoded data is in the course variable
	// Now pushing in the data in the courses slie that is in our fake db

	// Generating in the unique id and converting it into string
	rand.Seed(time.Now().Unix())

	// Itoa method converts in the pased integert to the string
	course.CourseId = strconv.Itoa(rand.Intn(1000))

	// we have a stringconv package in go used for string conversion here

	// pushing in the fake db
	// using in the append method to append in the course into thatc
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)

	return

}

// Updating in one course in go lang
// The u operation of the crud
func updateOne(w http.ResponseWriter, r *http.Request) {
	// The values which are to be updated
	// and the course which is to be updated using in the unique course id

	// to imitate we are gonna loop through the slice
	// and find in the course with that particular id
	// and if found we are going to update the course

	fmt.Println("Updating in a particular course with the unique course id ")
	w.Header().Set("Content-Type", "application/json")

	var infoFromReqToBeUpdated Course
	_ = json.NewDecoder(r.Body).Decode(&infoFromReqToBeUpdated)

	// Grabbing in the course id
	params := mux.Vars(r)

	// Looping through the courses data structure
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// Removing in the old course
			// db behavior imitiation

			// Forming in the slice once again now
			courses = append(courses[:index], courses[index+1:]...)

			// Decoding in the data which is received in the request body
			_ = json.NewDecoder(r.Body).Decode(&infoFromReqToBeUpdated)

			infoFromReqToBeUpdated.CourseId = params["id"]
			courses = append(courses, infoFromReqToBeUpdated)

			json.NewEncoder(w).Encode(courses)
			return
		}
	}
}
