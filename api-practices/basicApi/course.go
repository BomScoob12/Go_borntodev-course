package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
			fmt.Println(err)
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

func getNextID() int {
	highestID := -1
	for i := 0; i < len(courseList); i++ {
		if highestID < courseList[i].ID {
			highestID = courseList[i].ID
			fmt.Println(highestID)
		}
	}
	return highestID + 1
}

func main() {
	http.HandleFunc("/courses", courseHandler)

	http.ListenAndServe(":8080", nil)
}
