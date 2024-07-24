package main

import "fmt"

const count = 10

var name string = "Bom"

func test() {
	i := 0
	//while loop
	for i < 10 {
		fmt.Print("Hello World")
		i++
	}

	//for loop
	for j := 0; j < count; j++ {
		fmt.Println(name, " ", j)
	}

	//for loop with range
	slice := []string{"a", "b", "c"}
	//i is index, s is value
	for i, s := range slice {
		fmt.Println(i, s)
	}

	//for loop with range for map
	m := make(map[string]string)
	m["a"] = "apple"
	m["b"] = "banana"
	//k is key, v is value
	for k, v := range m {
		fmt.Println(k, v)
	}

	//for loop with range for string
	//i is index, c is unicode
	for i, c := range "go" {
		// unicode is 103, 111
		fmt.Println(i, c + 2)
		fmt.Println(string(c + 2))
	}

	// infinite loop\
	var name string
	for {
		fmt.Println("Infinite loop")
		fmt.Println("Enter your name: ")
		fmt.Scanf("%s", &name)
		fmt.Println("Your name is ", name)
		if name == "exit" {
			break
		}
	}
}

// func main() {
// 	test()
// }