package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	e := employee{}
	err := json.Unmarshal([]byte(`{"ID": 1, "EmployeeName": "abc def", "Tel": "123456", "Email": "abc@def.com"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e) // {1 abc def 123456 abc@def.com}
}
