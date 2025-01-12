package main

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
