package files

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Download(c *fiber.Ctx) error {
	file, err := os.OpenFile("./uploads/" + c.Query("file"), os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	io.Copy(c.Response().BodyWriter(), file)

	info, err := file.Stat()
	if err != nil {
		return err
	}

	Type := http.DetectContentType(c.Response().Body())
  Type = strings.Split(Type, ";")[0]
  size := strconv.FormatInt(info.Size(), 10)
  name := strings.Split(file.Name(), "/")[2]

  log.Printf("Downloaded a file of type %s with a size of %s with name %v", Type, size, name)

	c.Response().Header.Set("Content-Language", "pt-br")
	c.Response().Header.Set("Content-Type", Type + ";" + name)
	c.Response().Header.Set("Content-Length", size)

	return nil
}
