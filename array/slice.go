package main

import "fmt"

//slice values are dynamic (this call a slice variable)
// it look like an ArrayList in Java
var courseName []string

func main() {
	courseName = [] string {"Python", "Java", "Golang", "Javascript"}
	fmt.Println("Course Name: ", courseName)
	fmt.Println("Course Name: ", courseName[2])
	courseName = append(courseName, "Ruby", "C++", "C#", "HTML", "CSS", "JavaScript")
	fmt.Println("Course Name: ", courseName)

	// start 7 end at n-1
	courseWeb:= courseName[7:10]
	fmt.Println("Course Web: ", courseWeb)

	// first 5 element
	first5Element := courseName[:5]
	fmt.Println("First 5 Element: ", first5Element)

	//slice with make
	// make([]type, length, capacity)
	courseName = make([]string, 5, 10)
	fmt.Println("Course Name: ", courseName)
}