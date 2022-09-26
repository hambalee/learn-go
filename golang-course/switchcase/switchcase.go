package main

import (
	"fmt"
)

func main() {
	input := 3
	switch input {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Invalid value.")
	}

	var color string
	fmt.Printf("Enter color : ")
	fmt.Scanf("%s", &color)
	switch color {
	case "blue":
		fmt.Println("#0000FF")
	case "green":
		fmt.Println("#008000")
	case "pink":
		fmt.Println("#FFC0CB")
	case "yellow":
		fmt.Println("#FFFF00")
	default:
		fmt.Println("No color.")
	}
}
