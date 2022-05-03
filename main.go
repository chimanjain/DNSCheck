package main

import (
	"github.com/chimanjain/dnscheck/router"
	"github.com/gofiber/fiber/v2"
)

//For testing the API in local machine GET:[http://localhost:3000/dns/{url}]
func main() {
	app := fiber.New()
	router.InitializeDNSRoutes(app)
	app.Listen(":3000")
}
