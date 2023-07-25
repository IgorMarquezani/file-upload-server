package files

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/wifi-transfer/controllers"
)

func Delete(c *fiber.Ctx) error {
	if err := controllers.AuthenticateSession(c); err != nil {
		return fmt.Errorf("Você não fez login")
	}

	name := c.Query("file")

	err := os.Remove("./uploads/" + name)
	if err != nil {
		return err
	}

	return nil
}
