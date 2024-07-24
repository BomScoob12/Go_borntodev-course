package main

import "fmt"

var nameColor string

func main() {

	fmt.Println("Enter a number: ")
	fmt.Scanf("%s", &nameColor)

	switch nameColor {
	case "blue":
		fmt.Println("#0000FF")
	case "green":
		fmt.Println("#00FF00")
	case "pink":
		fmt.Println("#FFC0CB")
	default:
		fmt.Println("Invalid color")
	}
}
