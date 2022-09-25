package main

import "fmt"

func zeroValue(value int) {
	value = 0
}

func zeroPointer(pointer *int) {
	*pointer = 0
}

func main() {
	i := 1
	fmt.Println("i =", i) // i = 1

	zeroValue(i)
	fmt.Println("i from function zeroValue =", i) // i from function zeroValue = 1

	zeroPointer(&i)
	fmt.Println("i from function zero pointer =", i)  // i from function zero pointer = 0
	fmt.Println("i from function zero pointer =", &i) // i from function zero pointer = 0xc000016098
}
