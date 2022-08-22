package main

import "fmt"

func main() {
	fmt.Println("Welcome to the method class !")

	hitesh := User{"Raghu", "Raghu@swift.com", true, 16}

	fmt.Println(hitesh)
	fmt.Printf("Raghu Details are: %v, %v \n", hitesh.Name, hitesh.Email)

	hitesh.GetStatus()
	hitesh.NewMail()

	fmt.Printf("Raghu Details are: %v, %v \n", hitesh.Name, hitesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("User is active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "raghu@go.dev"
	fmt.Println("NewEmail is : ", u.Email)

}
