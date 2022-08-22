package main

import "fmt"

func main() {
	fmt.Println("Welcome to GoLang Pointers class")

	// var ptr *int
	// fmt.Println("Value of pointers: ", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value of pointers: ", ptr)
	fmt.Println("Value of pointers: ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value of the pointer is: ", *ptr)
	fmt.Println("New value of the pointer is: ", myNumber)

}
