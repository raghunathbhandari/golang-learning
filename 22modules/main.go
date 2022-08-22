package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/gorilla/mux"
)

// model for course file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"CourseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

// fake DB
var course []Course

// middleware helper file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

func main() {
	fmt.Println("hellow mod")

}

// Controller - files

//serve home files

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(course)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get one Course0")
	w.Header().Set("content-Type", "application/json")

	param := mux.Vars(r)

	//loop through courses, find matching

}
