package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter Value1 : ")
	float1 := takeUserInput()
	fmt.Print("Enter value 2 : ")
	float2 := takeUserInput()

	fmt.Print("Which operation you want to perform (+,-,*,/) : ")
	reader := bufio.NewReader(os.Stdin)
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	var result float64
	switch operation {
	case "+":
		result = float1 + float2
	case "-":
		result = float1 - float2
	case "*":
		result = float1 * float2
	case "/":
		result = float1 / float2
	default:
		fmt.Println("You chose wrong operation")

	}
	fmt.Println("Your result is ", result)
}

func takeUserInput() float64 {
	reader := bufio.NewReader(os.Stdin)
	value1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		panic("You enetered wrong value")
	}
	return float1
}
