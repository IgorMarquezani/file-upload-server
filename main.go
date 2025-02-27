package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/wifi-transfer/controllers"
	"github.com/wifi-transfer/controllers/users"
	"github.com/wifi-transfer/routes"
)

const (
	help = `--help: show this message.

-p {port-number}: especifies the port number.

-a {ip-address}: especifies the ip addres to be used. (Does not support ipv6).

--use-wifi: uses your machine ip addres in the local wi-fi network. (only works on linux and windows 11).

--add-user {name@password}: creates a user with the especified name and password. `
)

var (
	helped bool
	regex  *regexp.Regexp
	addr   = "127.0.0.1"
	port   = ":80"
)

func init() {
	if runtime.GOOS == "linux" {
		regex = regexp.MustCompile("(wlan.*)|(wlp2s.*)")
	}

	if runtime.GOOS == "windows" {
		regex = regexp.MustCompile("Wi-Fi.*")
	}
}

func getWifiNetworkIP() string {
	networks, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, network := range networks {
		if regex.MatchString(network.Name) {
			addrs, err := network.Addrs()
			if err != nil {
				panic(err)
			}

			for _, addr := range addrs {
				addrStr := strings.Split(addr.String(), "/")[0]

				parsed := net.ParseIP(addrStr)
				if parsed.To4() == nil {
					continue
				}

				return addrStr
			}
		}
	}

	return ""
}

func executeArgs(args []string) error {
	for i := 1; i < len(args); i += 2 {
		if args[i] == "--help" {
			fmt.Println(help)
			helped = true
			return nil
		}

		if args[i] == "-p" {
			if len(args) <= i+1 {
				return fmt.Errorf("Invalid argument for port number")
			}

			if _, err := strconv.Atoi(args[i+1]); err != nil {
				return fmt.Errorf("Invalid argument for port number")
			}

			port = ":" + args[i+1]
		}

		if args[i] == "-a" {
			if len(args) <= i+1 {
				return fmt.Errorf("Invalid argument for ip addr using ipv4")
			}

			parsed := net.ParseIP(args[i+1])
			if parsed.To4() == nil {
				return fmt.Errorf("Invalid ipv4 format")
			}

			addr = args[i+1]
		}

		if args[i] == "--use-wifi" {
			addr = getWifiNetworkIP()
			i--
		}

		if args[i] == "--add-user" {
			if len(args) <= i+1 {
				return fmt.Errorf("Invalid argument for user creation")
			}

			credentials := strings.Split(args[i+1], "@")
			if len(credentials) < 2 {
				return fmt.Errorf("Invalid argument for user creation")
			}

			users.AddUser(credentials[0], credentials[1])
		}
	}

	return nil
}

func main() {
	temp := html.New("views", ".html")

	temp.Reload(true)

	app := fiber.New(fiber.Config{
		Views:   temp,
		AppName: "http file transfer",
	})

	app.Server().StreamRequestBody = true

	routes.RegisterAll(app)

	err := executeArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	controllers.Ip = addr
	controllers.Port = port

	if !helped {
		log.Fatal(app.Listen(addr + port))
	}
}
