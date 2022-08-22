package main

import "fmt"

func main() {
	fmt.Println("My FUnction class")
	greeter()

	fmt.Println("Added result: ", adder(3, 5))

	proRes, proRes2 := proAdder(2, 3, 4, 5)
	fmt.Printf("Proadder Values %v , %v : ", proRes, proRes2)
}

func adder(firstVal int, secondVal int) int {
	return firstVal + secondVal
}

func greeter() {
	fmt.Println("My Greeter Fucntion")
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "This is string return in second value"
}
