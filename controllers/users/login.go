package users

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/wifi-transfer/controllers"
)

type User struct {
	Name   string
	Passwd string
}

var (
	users = make(map[string]User)
)

func AddUser(name, passwd string) error {
	if _, ok := users[name]; ok {
		return fmt.Errorf("User already exists")
	}

	users[name] = User{
		Name:   name,
		Passwd: passwd,
	}

	return nil
}

func getUser(name string) User {
	user, ok := users[name]
	if !ok {
		return User{}
	}

	return user
}

func ValidateLogin(c *fiber.Ctx) error {
	var (
		name   = c.FormValue("name")
		passwd = c.FormValue("passwd")
		user   = getUser(name)
	)

	if user.Name == "" || user.Passwd != passwd {
		c.Status(http.StatusUnauthorized)
		return c.SendString("Crendências de login inválidas")
	}

	se := controllers.NewSession(name)

	c.Cookie(&fiber.Cookie{
		Name:    "session",
		Value:   se.Id,
		Expires: time.Now().Add(time.Hour * 24),
		Path:    "/",
	})

	return c.SendString("loged in")
}

func LoginPage(c *fiber.Ctx) error {
	return c.Render("Login", nil)
}
