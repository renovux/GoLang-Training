package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	student[0] = "Nicholas"
	// student = append(student, "Nicholas")
	fmt.Println(student)
	fmt.Println(students)
}
