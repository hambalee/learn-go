package main

import "fmt"

var numberInt, numberInt2 int = 1000, 2000
var msg string = "Hello"

func main() {
	numberFloat := 1.2
	fmt.Println(numberInt)
	fmt.Println(numberInt2)
	fmt.Println(numberFloat)
	fmt.Println(msg)
	fmt.Println(numberFloat + float64(numberInt))
	fmt.Println(msg + " World!")
	fmt.Println("num = ", numberFloat)
}
