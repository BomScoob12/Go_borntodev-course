package main

import "fmt"

var number1, number2 int = 1000,2000
var numberFloat float32 = 82.5
var numberFloat64 float64 = 924.23

var msg string = "BomScoob"

func main() {
	//default is float64 (":=" is short format)
	newNumber := 22.5

	fmt.Println("Bom Hello world.")
	fmt.Println("Testing print.")
	fmt.Println(number1 + number2)
	fmt.Println(msg)
	fmt.Println(float32(number1) + numberFloat)
	fmt.Println(numberFloat64 + newNumber)
	fmt.Println("-------------")
	fmt.Println(msg, newNumber)
}
