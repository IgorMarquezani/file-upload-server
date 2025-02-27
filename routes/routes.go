package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/wifi-transfer/controllers"
	"github.com/wifi-transfer/controllers/files"
	"github.com/wifi-transfer/controllers/users"
)

func RegisterAll(app *fiber.App) {
	app.Get("/", users.LoginPage)
	app.Get("/user/login.html", users.LoginPage)
	app.Post("/user/validate", users.ValidateLogin)

	app.Get("/files/upload.html", controllers.Index)
	app.Get("/files/list", files.List)
	app.Get("/files/download", files.Download)
	app.Post("/files/upload", files.Upload)
	app.Delete("/files/delete", files.Delete)
	app.Static("/static", "public")
}
