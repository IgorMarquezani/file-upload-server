package files

import (
	_ "bytes"
	"fmt"
	"io"
	_ "log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wifi-transfer/controllers"
)

// func UploadWithStream(c *fiber.Ctx) error {
//   body := c.Request().BodyStream()
//   defer c.Request().CloseBodyStream()

//   buffer := make([]byte, 0, 1024*1024)

//   for {
//     length, err := io.ReadFull(body, buffer[:cap(buffer)])
//     buffer = buffer[:length]
//     if err != nil {
//       if err == io.EOF {
//         break
//       }
//       if err == io.ErrUnexpectedEOF {
//         return err
//       }
//     }
//
//     log.Printf("Read %d bytes", length)
//   }

// 	file, err := os.OpenFile("./uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	io.Copy(file, bytes.NewBuffer(buffer))

// 	return c.Redirect("http://10.0.0.212:80/", http.StatusOK)
// }

func Upload(c *fiber.Ctx) error {
	if err := controllers.AuthenticateSession(c); err != nil {
		return fmt.Errorf("Você não fez login")
	}

	header, err := c.FormFile("file")
	if err != nil {
		return err
	}

	var filename string

	fmt.Println(string(c.Request().Header.Header()))
	infos := strings.Split(c.Get("Content-Disposition"), " ")
	for _, str := range infos {
		regex, err := regexp.Compile("filename.*")
		if err != nil {
			return err
		}

		if regex.MatchString(str) {
			name := strings.Split(str, "=")[1]
			name = name[1 : len(name)-1]
			filename = name
		}
	}
	fmt.Println(filename)

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
