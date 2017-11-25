package main

import "fmt"

func main() {
	age := 24
	changeMe(age)
	fmt.Println(age) // 24
}

func changeMe(z int) {
	fmt.Println(z)
	z = 24
}

// when changeMe is called on line 8
// the value 44 is being passed as an argument
