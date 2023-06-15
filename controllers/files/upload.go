package files

import (
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	header, err := c.FormFile("file")
	if err != nil {
		return err
	}

	buffer, err := header.Open()
	if err != nil {
		return err
	}
	defer buffer.Close()

	file, err := os.OpenFile("./uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	io.Copy(file, buffer)

	return c.Redirect("http://10.0.0.212:80/", http.StatusOK)
}
