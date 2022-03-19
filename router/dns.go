package router

import (
	"dnscheck/controller"

	"github.com/gin-gonic/gin"
)

// InitializeDNSRoutes initializes the router for DNS Checkup api
func InitializeDNSRoutes(r *gin.Engine) {

	// Routes
	r.GET("/dns/:url", controller.FindDNS)

}
