package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("statusTodo.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 1
	indexCol := 1
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		item := strings.Split(line, ",")
		fmt.Printf("Row : %d = %s\n", count,item[indexCol])
		count++
	}
}