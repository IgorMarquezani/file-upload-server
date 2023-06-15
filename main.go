package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/wifi-transfer/routes"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.RegisterAll(app)

	log.Fatal(app.Listen("10.0.0.212:80"))
}
