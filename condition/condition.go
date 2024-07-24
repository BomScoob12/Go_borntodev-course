package main

import "fmt"

var score int

func calculateGrade(score int) string {
	if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 60 {
		return "C"
	} else if score >= 50 {
		return "D"
	} else {
		return "E"
	}

}

func main() {
	fmt.Println("Enter your score: ")
	fmt.Scanf("%d",&score)
	fmt.Print("Your score is ", score, "\n")
	yourGrade := calculateGrade(score)
	fmt.Println("Your grade is", yourGrade)
}