package main

import "fmt"

func hello()  {
	fmt.Println("Hello") // Hello
}

func plus(value1 int, value2 int)  {
	result := value1 + value2
	fmt.Println("result =", result)  // result = 3
}

func plus2(value1 int, value2 int) int {
	return value1 + value2
}

func plus3vvalue(value1, value2, value3 int) int {
	return value1 + value2 + value3
}

func main() {
	hello()
	plus(1, 2)
	result2 := plus2(1, 2)
	fmt.Println("result2 =", result2) // result2 = 3
	result3 := plus3vvalue(5, 5, 10)
	fmt.Println("result3 =", result3) // result3 = 20
}
