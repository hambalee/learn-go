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
	fmt.Printf("%v\n", numberFloat)

	// Array
	var productName [4]string
	productName[0] = "A"
	productName[1] = "B"
	productName[2] = "C"
	productName[3] = "D"
	price := [4]float32{5, 6, 7, 8}
	fmt.Println(productName)
	fmt.Println(price)
}
