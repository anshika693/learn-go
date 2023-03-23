package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(updateString("anshika", "is the best"))

	fmt.Println("Sum is ", calculateSum(2, 3, 1, 4, 5, 6))
}

func updateString(s1, s2 string) (str1, str2 string) {
	s1 = strings.ToUpper(s1)
	s2 = strings.Title(s2)
	return s1, s2
}

func calculateSum(values ...int) int {
	sum := 0
	for _, i := range values {
		sum = sum + i
	}
	return sum
}
