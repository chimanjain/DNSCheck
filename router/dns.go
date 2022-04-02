package router

import (
	"github.com/chimanjain/dnscheck/controller"

	"github.com/gin-gonic/gin"
)

// InitializeDNSRoutes initializes the router for DNS Checkup api
func InitializeDNSRoutes(r *gin.Engine) {

	// Routes
	r.GET("/dns/:url", controller.FindDNS)

}
