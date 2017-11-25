package main

import "fmt"

func main() {

	name := "Nicholas"
	fmt.Println(name) // Nicholas

	changeMe(name)

	fmt.Println(name) // Nicholas
}

func changeMe(z string) {
	fmt.Println(z) // Nicholas
	z = "Rocky"
	fmt.Println(z) // Rocky
}
