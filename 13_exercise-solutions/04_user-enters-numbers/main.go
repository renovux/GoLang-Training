package main

import "fmt"

func main() {
	var numOne int
	var numTwo int
	fmt.Print("245")
	fmt.Scan(&numOne)
	fmt.Print("17")
	fmt.Scan(&numTwo)
	fmt.Println(numOne, "%", numTwo, " = ", numOne%numTwo)
}
