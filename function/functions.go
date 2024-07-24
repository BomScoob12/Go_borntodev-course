package main

import (
	"fmt"
	"strings"
)

// void
func hello() {
	fmt.Println("Hello World")
}

// with param
func toLowerCaseString(str string, str2 string) {
	lowerStr := strings.ToLower(str)
	lowerStr2 := strings.ToLower(str2)
	fmt.Println(lowerStr, lowerStr2)
}

// return value
func addingNumber(numbers ...int) int {
	if (len(numbers) == 0) {
		return -1
	}

	tempNumber := 0
	for number := range numbers {
		tempNumber += number
	}
	return tempNumber
}

func main() {
	hello()
	toLowerCaseString("BomSCoob", "#Hey I'M A GOOD PERSON")

	fmt.Println(addingNumber(0))
	fmt.Println(addingNumber())
	fmt.Println(addingNumber(1,2,3,4,5,6,7,8,9,10))
}