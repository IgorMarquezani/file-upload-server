package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/wifi-transfer/controllers"
	"github.com/wifi-transfer/controllers/files"
)

func RegisterAll(app *fiber.App) {
	app.Get("/", controllers.Index)
	app.Get("/files", files.List)
  app.Get("/download", files.Download)
	app.Post("/upload", files.Upload)
	app.Delete("/delete", files.Delete)
}
