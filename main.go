package main

import (
	"log"
	"task/database"
	"task/route"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Post("/user", route.Registration)
	app.Get("/user", route.GetUsers)
}
func main() {
	database.Migration()
	app := fiber.New()
	Routes(app)
	log.Fatal(app.Listen(":3000"))
}
