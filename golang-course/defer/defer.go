package main

import "fmt"

func deferloop() {
	for j:= 0; j < 10; j++ {
		defer fmt.Println("j :", j)
	}
}

func main() {
	fmt.Println("Welcome")
	defer fmt.Println("End")
	deferloop()
}
// Welcome
// j : 9
// j : 8
// j : 7
// j : 6
// j : 5
// j : 4
// j : 3
// j : 2
// j : 1
// j : 0
// End