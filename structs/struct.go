package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {

	employee1 := employee{
		employeeID:   "100",
		employeeName: "Sarawit BomScoob",
		phone:        "0666666666",
	}
	fmt.Println(employee1)

	// ----------- array, slice
	employeeList := [3]employee{}
	employeeList[0] = employee{
		employeeID: "101",
		employeeName: "Bob",
		phone: "0999999999",
	}
	employeeList[1] = employee{
		employeeID: "102",
		employeeName: "Bob",
		phone: "0999999999",
	}
	employeeList[2] = employee{
		employeeID: "103",
		employeeName: "Bob",
		phone: "0999999999",
	}

	employeeSlice := []employee{}

	employeeSlice = append(employeeSlice, employee1)

	fmt.Println(employeeList)
	fmt.Println(employeeSlice)
	
}