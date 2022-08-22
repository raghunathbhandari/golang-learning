package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Post struct {
	Id          uint
	Title       string
	Description string
}

func main() {

	//fmt.Println("Microservices POST Method!")

	db, err := gorm.Open(mysql.Open("root:'Raghu@123'@tcp(127.0.0.1:3306)/posts_ms"), &gorm.Config{})
	fmt.Println("DB Connection!")

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Post{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8000")
}
