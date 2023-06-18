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
	file, err := os.OpenFile("./uploads/"+c.Query("file"), os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	io.Copy(c.Response().BodyWriter(), file)

	info, err := file.Stat()
	if err != nil {
		return err
	}

	Type := strings.Split(http.DetectContentType(c.Response().Body()), ";")[0]
	size := strconv.FormatInt(info.Size(), 10)
	name := strings.Split(file.Name(), "/")[2]

	log.Printf("Downloaded a file of %s bytes named %v of the type %s", size, name, Type)

	c.Response().Header.Set("Content-Language", "pt-BR")
	c.Response().Header.Set("Content-Type", Type+";"+"charset=utf-8")
	c.Response().Header.Set("Content-Disposition", `attachment; filename="`+name+`"`)
	c.Response().Header.Set("Content-Length", size)

	return nil
}
