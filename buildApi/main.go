package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// model for course -file
type Course struct {
     CourseId string  `json:"courseid"`
	 CourseName string `json:"coursename"`
	 CoursePrice  int   `json:"price"`
	 Author  *Author  `json:"author"` // type is a pointer
     
}

type Author struct {
	FullName string  `json:"fullname"`
	Website  string `json:"website"`

}

//fake db

var courses []Course
  //middleware and helpers - file separate 
func (c *Course) IsEmpty() bool {
  return c.CourseId == "" && c.CourseName == ""
}
func main() {
   fmt.Println("welcome to API by LearnCodeOnline")
   r := mux.NewRouter()
   courses = append(courses, Course{CourseId: "1", CourseName: "reactjs bootcamp", CoursePrice: 299, Author: &Author{FullName: "John Doe", website: "lco.dev"}})
    courses = append(courses, Course{CourseId: "2", CourseName: "MERN bootcamp", CoursePrice: 399, Author: &Author{FullName: "Mahima singh", website: "100x.dev"}})	 
	
	
	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - separate file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("<h1>welcome to API by learningCode Online</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) 
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")
	courseId := mux.Vars(r)

	//loop through courses, find matching courseId and return the response
	for_, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
            return
		}
	}
	//if course not found
	json.NewEncoder(w).Encode("course not found")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("create course")
	w.Header().Set("Content-Type", "application/json")


	//what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("pls send some data")
		
	}

	//what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
        json.NewEncoder(w).Encode("No data inside json")
		return
	}

	rand.Seed(time.Now().UnixNano())
course.CourseId  = strconv.Itoa(rand.Intn(100))
courses = append(courses, course)
json.NewEncoder(w).Encode(course)
return
}

func updateCourse( w http.ResponseWriter, r *http.Request) {
	fmt.Println("update course")
	w.Header().Set("Content-Type", "application/json")
	courseId := mux.Vars(r)

	//loop through courses, find matching id and remove it then add with my ID
	for index, course := range courses{
      if course.CourseId == params["id"] {
		courses = append(courses[:index], courses[index+1:]...)
		var course Course
		_ = json.NewDecoder(r.Body).Decode(&course)
		course.CourseId = params["id"]
		courses = append(courses, course)
		json.NewEncoder(w).Encode(course)
		return
	}

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete course")
	w.Header().Set("Content-Type", "application/json")
	courseId := mux.Vars(r)
   for index, course := range courses {
		if course.CourseId == params["id"] {
            courses = append(courses[:index], courses[index+1]...)
			break;
		}
	}
}
 