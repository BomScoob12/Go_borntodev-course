package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func toMarshal(emp *employee) string {
	outputData, _ := json.Marshal(emp)
	fmt.Println(string(outputData))
	return string(outputData)
}

// unMarshal is a function to convert json to object
// recieve a data in json format and set it to the object
func unMarshalAndChangeData(jsonString string,emp *employee) employee {
	err := json.Unmarshal([]byte(jsonString), &emp)
	if err != nil {
		log.Fatal(err)
	}
	emp.Email = "bobidubiduwa@mail.com"
	return *emp
}

func main() {
	emp1 := employee{12, "Bobby", "0658884547", "brogot@email.com"}
	// send an object to marshal (with a pointer)
	jsonString := toMarshal(&emp1)
	fmt.Println("Employee 1 after marshal: ", jsonString)
	
	// after unmarchal, emp1 will be changed
	// because emp1 send a pointer to the function
	unMarshalAndChangeData(jsonString, &emp1)
	fmt.Println("Employee 1 after un marchal: ", emp1)
}