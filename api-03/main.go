package main

import (
	"api-03/database"
	"api-03/routes"

	"github.com/gofiber/fiber/v2"
)

func main() { 
	database.ConnectDb()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the API!")
	})
	app.Get("/users", routes.GetUsers)
	app.Post("/users", routes.CreateUser)
	app.Get("/users/:id", routes.GetUser)

	app.Listen(":8080")
 }