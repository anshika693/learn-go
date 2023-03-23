package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Value 1: ")
	reader := bufio.NewReader(os.Stdin)
	value1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		panic("Value 1 is not integer")
	}
	fmt.Print("Value 2: ")
	value2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err != nil {
		panic("Value 2 is not integer")
	}
	sum := float1 + float2
	fmt.Println("Sum is ", sum)
}
