package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	data, _ := json.Marshal(&employee{1, "abc def", "123456", "abc@def.com"})
	fmt.Println(string(data)) // {"ID":1,"EmployeeName":"abc def","Tel":"123456","Email":"abc@def.com"}
}
