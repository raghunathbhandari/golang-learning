package main

import "fmt"

func main() {
	fmt.Println("welcome to array session!")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Peach"
	fruitList[3] = "mango"

	fmt.Println("Array Value: ", fruitList)
	fmt.Println("Array Value: ", len(fruitList))

	var vegList = [5]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Vegy List are:", len(vegList))

}
