package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Structs classes")

	/*
		hitesh := User{"hitesh", "hitesh@go.dev", true, 16}

		fmt.Println(hitesh)
		fmt.Printf("hitesh details are : %+v \n", hitesh)
		fmt.Printf("hitesh is  %v and Email is %v", hitesh.Name, hitesh.Email)

		rand.Seed(time.Now().UnixNano())
		diceNumber := rand.Intn(6) + 1
		fmt.Println(diceNumber)

		switch diceNumber {
		case 1:
			fmt.Println("Dice Value is 1")
		case 2:
			fmt.Println("Dice Value is 2")
		case 3:
			fmt.Println("Dice Value is 3")
			fallthrough
		default:
			fmt.Println("What is that")
		}
	*/

	//days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	//fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	roughValue := 1

	for roughValue < 10 {

		if roughValue == 5 {
			roughValue++
			goto raghu
		}

		fmt.Println("rough value is: ", roughValue)
		roughValue++
	}

raghu:
	fmt.Println("Jumping in the last line Raghu!")
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
