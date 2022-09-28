package main

import (
	"fmt"
	"time"
)

func greeting(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, " : ", i)
	}
}

func main() {
	go greeting("Hello")
	go greeting("Hola")
	time.Sleep(1 * time.Second)
}
