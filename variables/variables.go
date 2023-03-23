package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	name = "anshika"
)

func main() {
	course := "Go-lang Course"
	fmt.Println("I am able to access variable ", name)
	fmt.Println("Course is local variable ", course)

	module := "4" //Its a sting
	clip := 2
	// sum := module + clip WILL GIVE ERROR. Need to convert module to integer
	iModule, err := strconv.Atoi(module)
	if err == nil {
		sum := iModule + clip
		fmt.Println("Sum is ", sum)
	}

	changeValue(course)
	fmt.Println("Course Value is not changed ", course)

	updateCourse(&course)
	fmt.Println("Now course value is changed ", course)

	// fmt.Println(os.Environ()) //will show all environment variables
	fmt.Println(os.Getenv("USER"))
}

func changeValue(course string) string {
	course = "New Course"
	fmt.Println("Now the value of course is", course)
	return course
}

func updateCourse(course *string) string {
	*course = "New Course"
	fmt.Println("New Course is ", *course)
	return *course
}
