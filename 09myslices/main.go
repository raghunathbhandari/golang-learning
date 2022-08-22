package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to My Slices !")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 567
	highScores[3] = 789

	highScores = append(highScores, 555, 666, 444)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

}
