package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil{
		message, _ := fmt.Scanf("%v must be number.", promt)
		panic(message)
	}

	return value
}

func getOperator() string {
	fmt.Print("Operator is in ( + - * /):")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func add(val1 float64, val2 float64) float64 {
	return val1 + val2;
}

func minus(val1 float64, val2 float64) float64 {
	return val1 + val2;
}

func divide(val1 float64, val2 float64) float64 {
	return val1 + val2;
}

func multiply(val1 float64, val2 float64) float64 {
	return val1 + val2;
}
func main() {
	value1 := getInput(" Enter first number: ")
	value2 := getInput(" Enter secound number: ")
	
	var result float64

	switch operator := getOperator(); 
	operator {
		case "+":
			result = add(value1, value2)
		case "-":
			result = minus(value1, value2)
		case "*":
			result = multiply(value1, value2)
		case "/":
			result = divide(value1, value2)
	}

	fmt.Printf("%v is the result", result)
}	