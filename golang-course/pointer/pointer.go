package main

import "fmt"

func zerovalue(ivalue int)  {
	ivalue = 0
}

func zeropointer(ipointer *int)  {
	*ipointer = 0
}

func main() {
	i := 1
	fmt.Println("i =", i) // i = 1

	zerovalue(i)
	fmt.Println("i from function zerovalue =", i) // i from function zerovalue = 1

	zeropointer(&i)
	fmt.Println("i from function zero pointer =", i) // i from function zero pointer = 0
	fmt.Println("i from function zero pointer =", &i) // i from function zero pointer = 0xc000016098
}
