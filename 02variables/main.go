package main

import "fmt"

const LoginToken string = "ghashdf"

func main() {
	//fmt.Println("Variables")
	var username string = "OM Nama Shivaya!"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var IsLoggedIn bool = false
	fmt.Println(IsLoggedIn)
	fmt.Printf("Variable is of type: %T \n", IsLoggedIn)

	var smallVal int8 = 25
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 25.1231233223232333
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default value and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implecit type
	// var website string = "OM Nama"
	// x, x1 := fmt.Printf(website)
	// x, x1

	//no var style

	var numberOfUser int32 = 300000
	fmt.Println(numberOfUser)

	//public variable

}
