package main

import (
	"fmt"
)

var product = make(map[string]float64)

func main() {
	fmt.Println("product =", product) // product = map[]

	//add
	product["MacBook"] = 40000
	product["iPhone"] = 30000
	product["iPad"] = 25000
	fmt.Println("product =", product) // product = map[MacBook:40000 iPad:25000 iPhone:30000]

	//delete
	delete(product, "iPad")
	fmt.Println("product =", product) // product = map[MacBook:40000 iPhone:30000]

	//update
	product["iPhone"] = 20000
	fmt.Println("product =", product) // product = map[MacBook:40000 iPhone:20000]

	//value
	value1 := product["iPhone"]
	fmt.Println("value1 =", value1) // value1 = 20000
	//short form
	courseName := map[string]string{"101":"Java","102":"Python"}
	fmt.Println("courseName =", courseName) // courseName = map[101:Java 102:Python]
}
