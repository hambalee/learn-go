package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string) float64 {
	fmt.Printf("%v", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		input = strings.TrimRight(input, "\r\n")
		message, err := fmt.Printf("%s ERR: not string must number only \n", input)
		if err != nil {
			fmt.Println(err)
		}
		panic(message)
	}
	return value
}

func getOperator() string {
	str := ""
	for str != "+" && str != "-" && str != "*" && str != "/" {
		fmt.Print("Operator is ( + - * / ) : ")
		op, _ := reader.ReadString('\n')
		str = strings.TrimSpace(op)
		if str != "+" && str != "-" && str != "*" && str != "/" {
			fmt.Println("Wrong operator")
		}
	}
	return str
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}
func subtract(value1, value2 float64) float64 {
	return value1 - value2
}
func multiply(value1, value2 float64) float64 {
	return value1 * value2
}
func divide(value1, value2 float64) float64 {
	return value1 / value2
}

func main() {
	nInput := getInput("Number of input : ")
	fmt.Printf("Number of input is %v\n", nInput)

	value1 := getInput("value1 = ")
	var result float64
	result = value1

	for i := 2; i <= int(nInput); i++ {
		operator := getOperator()
		value2 := getInput("value " + strconv.Itoa(i) + " = ")
		switch operator {
		case "+":
			result = add(result, value2)
		case "-":
			result = subtract(result, value2)
		case "*":
			result = multiply(result, value2)
		case "/":
			result = divide(result, value2)
		default:
			panic("wrong operator")
		}
	}
	fmt.Printf("result is %v", result)
}
