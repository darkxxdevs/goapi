package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"name"`
	CoursePrice string  `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	AuthorName string `json:"author_name"`
	Website    string `json:"auhtor_websitr"`
}

func (c *Course) isEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == " "
}

var courses []Course

func main() {
	fmt.Println("API -drxxdevs")

	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "GoLang",
		CoursePrice: "1000",
		Author: &Author{
			AuthorName: "arpit",
			Website:    "https://golang.org",
		},
	})

	// routing

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/post", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{courseid}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses/update/{courseid}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/delete/{courseid}", deleteOneCourse).Methods("DELETE")
	r.HandleFunc("/courses/deleteall", deleteAllCourse).Methods("DELETE")

	http.ListenAndServe(":8080", r)

}

// handler functions

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>api from darxxdevs</h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recived request!")
	w.Header().Set("Content-Type", "application/json")

	if len(courses) == 0 {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("<h1>no courses found </h1>"))
		return
	}
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course-")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// fmt.Println(params)

	for _, c := range courses {
		if (c.CourseId) == params["courseid"] {
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	json.NewEncoder(w).Encode("no course found !")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("body is empty")
		return
	}
	var c Course
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	if c.isEmpty() {
		json.NewEncoder(w).Encode("course is empty")
		return
	}

	courses = append(courses, c)
	json.NewEncoder(w).Encode(c)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating one course...")
	w.Header().Set("Content-Type", "application/json")

	// Grabbong the id to update the course with the given id

	params := mux.Vars(r)

	// loop through the courses and delete the course with the given id and add the course again with the update details the we recive from request body

	for index, course := range courses {
		if course.CourseId == params["courseid"] {
			courses = append(courses[:index], courses[index+1:]...)
			var crs Course

			err := json.NewDecoder(r.Body).Decode(&crs)

			if crs.isEmpty() {
				json.NewEncoder(w).Encode("course is empty")
				return
			}
			if err != nil {
				json.NewEncoder(w).Encode(err)
				return
			}
			fmt.Println(crs)
			crs.CourseId = params["courseid"]
			courses = append(courses, crs)
			json.NewEncoder(w).Encode(crs)
			return
		}
	}

	// adding the course from the request parameters

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting ...")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["courseid"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}
}

func deleteAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting all courses...")
	w.Header().Set("Content-Type", "application-json")
	courses = []Course{}
	json.NewEncoder(w).Encode(courses)
}
