package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sixfwa/fiber-api/database"
	"github.com/sixfwa/fiber-api/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the API")
}

func setUpRoutes(app *fiber.App) {
	//welcome endpoint
	app.Get("/", welcome)
	//user endpoints
	app.Post("/users", routes.CreateUser)
	app.Get("/users", routes.GetUsers)
	app.Get("/users/:id", routes.GetUser)
	app.Put("/users/:id", routes.UpdateUser)
	app.Delete("/users/:id", routes.DeleteUser)

	// app.Get("/products", routes.GetProducts)
	// app.Post("/products", routes.CreateProduct)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setUpRoutes(app)

	log.Fatal(app.Listen(":8080"))
}