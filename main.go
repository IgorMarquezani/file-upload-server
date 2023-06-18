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

	"github.com/wifi-transfer/routes"
)

var (
	regex *regexp.Regexp
	addr  = "0.0.0.0"
	port  = ":80"
)

func init() {
	if runtime.GOOS == "linux" {
		regex = regexp.MustCompile("wlan.*")
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
	}

	return nil
}

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:   engine,
		AppName: "http file transfer",
	})

	app.Server().StreamRequestBody = true

	routes.RegisterAll(app)

	err := executeArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	log.Fatal(app.Listen(addr + port))
}
