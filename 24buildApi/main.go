package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Model for Course - file

type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	Price int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullnmae string `json:"fullname"`
	Website string `json:"website"`
}

// fake db
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

// controllers - file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Learncodeonline.in</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all course")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}