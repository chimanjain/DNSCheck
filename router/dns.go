package router

import (
	"github.com/chimanjain/dnscheck/controller"
	"github.com/gofiber/fiber/v2"
)

// InitializeDNSRoutes initializes the router for DNS Checkup api
func InitializeDNSRoutes(app *fiber.App) {
	// Routes
	app.Get("/dns/:url", controller.GetDNS)
}
