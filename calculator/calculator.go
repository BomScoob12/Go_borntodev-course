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
	if err != nil {
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
	return val1 + val2
}

func minus(val1 float64, val2 float64) float64 {
	return val1 - val2
}

func divide(val1 float64, val2 float64) float64 {
	return val1 / val2
}

func multiply(val1 float64, val2 float64) float64 {
	return val1 * val2
}

func main() {
	round := int(getInput("How many number do you want to calculated? :"))

	var result float64
	result = getInput("number 1 :")
	
	for i := 1; i <= round-1; i++ {
		operator := getOperator()
		
		tmp := getInput("number" + strconv.Itoa(i+1) + " :")
		
		switch operator {
		case "+":
			result = add(result, tmp)
		case "-":
			result = minus(result, tmp)
		case "*":
			result = multiply(result, tmp)
		case "/":
			result = divide(result, tmp)
		default:
			panic("Wrong operator")
		}
	}

	fmt.Printf("%v is the result!", result)
}
