package main

import (
	"fmt"
)

func form(textForm string) {
	for i := 0; i < 100; i++ {
		fmt.Println(textForm, " :", i)
	}
}