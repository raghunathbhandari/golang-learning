package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
)

type Students struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func main() {

	readUSer()

}

func readUSer() {

	db, err := sql.Open("mysql", "root:`Raghu@123`@/userdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Connection success!")

	results, err2 := db.Query("SELECT id, first_name,last_name,email from students")
	if err != nil {
		panic(err2)
	}
	defer results.Close()
	fmt.Println("Query success!")

}
