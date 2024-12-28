package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Course struct {
	ID         int
	Title      string
	Price      float64
	Instructor string
}

var courseList []Course

func init() {
	CourseJSON := `[
		{
			"ID": 1,
			"Title": "Golang",
			"Price": 1000,
			"Instructor": "Bobby"
		},
		{
			"ID": 2,
			"Title": "Python",
			"Price": 1500,
			"Instructor": "Bobby"
		}
	]`

	// Unmarshal the JSON byte slice to a Go data structure
	err := json.Unmarshal([]byte(CourseJSON), &courseList)
	if err != nil {
		log.Fatal(err)
	}

}

func findId(ID int) (*Course, int) {
	for i, course := range courseList {
		if course.ID == ID {
			return &course, i
		}
	}

	return nil, -1
}

func getNextID() int {
	highestID := -1
	for i := 0; i < len(courseList); i++ {
		if highestID < courseList[i].ID {
			highestID = courseList[i].ID
		}
	}
	return highestID + 1
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	courseJSON, err := json.Marshal(courseList)

	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)
	case http.MethodPost:
		var newCourse Course

		bodyByte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// assign new value to newCourse
		err = json.Unmarshal(bodyByte, &newCourse)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		newCourse.ID = getNextID()
		courseList = append(courseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func courseHandlerID(w http.ResponseWriter, r *http.Request) {
	pathSplit := strings.Split(r.URL.Path, "/")
	ID := pathSplit[len(pathSplit)-1]

	IDInt, err := strconv.Atoi(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	course, itemIndex := findId(IDInt)

	if course == nil {
		http.Error(w, "Course not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		courseJSON, err := json.Marshal(course)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)

	case http.MethodPut:
		var newCourse Course

		bodyByte, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(bodyByte, &newCourse)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if newCourse.ID != IDInt {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		course = &newCourse
		courseList[itemIndex] = *course
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func main() {
	http.HandleFunc("/courses/", courseHandlerID)
	http.HandleFunc("/courses", courseHandler)
	http.ListenAndServe(":8080", nil)
}
