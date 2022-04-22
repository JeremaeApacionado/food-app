package main

import (
	"fmt"
	"log"
	"fooddel-app/database"
	"fooddel-app/route"

	"github.com/gofiber/fiber/v2"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome!")
}

func Routes(app *fiber.App) {

	//users
	app.Post("/user", route.Registration)
	app.Get("/user", route.GetUsers)
	//foodlist
	app.Post("/product", route.AddProduct)
	app.Get("/product", route.GetProductName)
	app.Get("/product/:id", route.GetProduct)
	app.Delete("/product/:id", route.Delete)
	app.Put("/product/:id", route.Update)
	//image
	app.Get("/images", route.GetImages)
	app.Get("/images/:id", route.GetImage)

}

func main() {
	database.Migration()
	app := fiber.New()
	Routes(app)
	log.Fatal(app.Listen(":3000"))
	fmt.Println("instaCart")
	app.Get("/", Welcome)
}
