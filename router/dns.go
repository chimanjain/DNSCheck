package router

import (
	"DNSCheck/controller"

	"github.com/gin-gonic/gin"
)

// InitializeDNSRoutes initializes the router for blog api
func InitializeDNSRoutes(r *gin.Engine) {

	// Routes
	r.GET("/dns/:url", controller.FindDNS)

}
