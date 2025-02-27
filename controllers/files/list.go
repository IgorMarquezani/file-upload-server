package files

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/wifi-transfer/controllers"
)

func List(c *fiber.Ctx) error {
	if err := controllers.AuthenticateSession(c); err != nil {
		return c.SendString("Você não fez login")
	}

	server := ServerFiles{
		Ip:    controllers.Ip,
		Port:  controllers.Port,
		Files: make([]File, 0),
	}

	files, err := os.ReadDir("./uploads")
	if err != nil {
		return err
	}

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			continue
		}

		server.Files = append(server.Files, File{
			Name:     file.Name(),
			Size:     float64(info.Size() / 1024),
			Modified: info.ModTime().Format("02-01-2006"),
		})
	}

	return c.Render("Home", server)
}
