package main

import "fmt"

func add(a, b float64) float64 {
	return a + b
}

func logLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println("Loop:", i)
	}
}

func deferLogLoop() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("defer Loop:", i)
	}
}

func main() {
	fmt.Println("Hello, playground")
	//defer is used for excecute bottom to top order
	//defer is used for clean up resources
// 	defer fmt.Println("Result 1:", add(10, 2))
// 	defer fmt.Println("Result 2:", add(1, 2))
// 	fmt.Println("Result 3:", add(1, 88))
// 	fmt.Println("End of main")

	// first in first out
	logLoop()

	// last in first out
	deferLogLoop()
}
// Output
// Hello, playground
// Loop: 0
// Loop: 1
// Loop: 2
// Loop: 3
// Loop: 4
// Loop: 5
// Loop: 6
// Loop: 7
// Loop: 8
// Loop: 9
// defer Loop: 9
// defer Loop: 8
// defer Loop: 7
// defer Loop: 6
// defer Loop: 5
// defer Loop: 4
// defer Loop: 3
// defer Loop: 2
// defer Loop: 1
// defer Loop: 0