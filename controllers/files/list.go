package files

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

type File struct {
	Name     string
	Size     float64
	Modified string
}

func List(c *fiber.Ctx) error {
	var list []File

	files, err := os.ReadDir("./uploads/")
	if err != nil {
    return err
	}

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			continue
		}

		list = append(list, File{
			Name:     file.Name(),
			Size:     float64(info.Size() / 1024),
			Modified: info.ModTime().Format("02-01-2006"),
		})
	}

	return c.Render("Home", list)
}
