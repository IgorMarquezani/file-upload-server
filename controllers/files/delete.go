package files

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Delete(c *fiber.Ctx) error {
	name := c.Query("file")

	err := os.Remove("./uploads/" + name)
	if err != nil {
    return err
	}

	return nil
}
