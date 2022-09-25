package main

import (
	"fmt"
)

func printGrade(score int) {
	if score >= 80 {
		fmt.Println("A")
	} else if score >= 70 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else if score >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}

func main() {
	myMoney := 100
	if myMoney >= 100 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	fmt.Println("Grade Calculator")
	var score int
	fmt.Scanf("%d", &score)
	fmt.Println("Score =", score)
	printGrade(score)
}
