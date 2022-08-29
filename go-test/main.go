package main

import "fmt"

func CalSum(x int, y int) (result int) {
	result = x + y
	return result
}

func main() {
	fmt.Println("My Simple Testing Tutorial")
	result := CalSum(2, 3)
	fmt.Println(result)

}
