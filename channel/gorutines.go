package main

import (
	"fmt"
	"time"
)

func form(textForm string) {
	for i := 0; i < 100; i++ {
		fmt.Println(textForm, " :", i)
	}
}

func main() {
	fmt.Println("This is a concerrent runing, switching to another task")
	go form("task 1")
	go form("task 2")
	time.Sleep(1 * time.Second)
}