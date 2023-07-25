package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wifi-transfer/utils"
)

type Server struct {
	Ip   string
	Port string
}

type Session struct {
	Id       string
	UserName string
}

const (
	ErrNoSuchSession = "No such Session"
	ErrNotLoged      = "Not loged in"
)

var (
	Ip       string
	Port     string
	sessions = make(map[string]Session)
)

func getSession(id string) (Session, error) {
	session, ok := sessions[id]
	if !ok {
		return Session{}, fmt.Errorf(ErrNoSuchSession)
	}

	return session, nil
}

func NewSession(name string) Session {
	id := string(utils.RandomArray(30))
	for _, ok := sessions[id]; ok; {
		id = string(utils.RandomArray(30))
	}

	se := Session{
		Id:       id,
		UserName: name,
	}

	sessions[id] = se

	return se
}

func AuthenticateSession(c *fiber.Ctx) error {
	cookie := c.Cookies("session")
	if cookie == "" {
		return fmt.Errorf(ErrNotLoged)
	}

	_, err := getSession(cookie)
	if err != nil {
		return fmt.Errorf(ErrNotLoged)
	}

	return nil
}

func Index(c *fiber.Ctx) error {
	server := Server{
		Ip:   Ip,
		Port: Port,
	}

	return c.Render("Index", server)
}
