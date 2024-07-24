package main

import "fmt"

func setZero(num int) {
	num = 0
}

// setting star for telling this is pointer variable
func setPointerZero(numPointer *int) {
	*numPointer = 0
}

func main() {
	i := 1
	fmt.Println("i = ", i)
	setZero(i)
	fmt.Println("i from function set Zero = ", i)

	// passing address of i by adding & symbol
	setPointerZero(&i)
	fmt.Println("i value from function set Zero (Pointer) = ", i)
	fmt.Println("i address from function set Zero (Pointer) = ", &i)
}