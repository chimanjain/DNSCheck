package main

import (
	"github.com/chimanjain/dnscheck/router"

	"github.com/gin-gonic/gin"
)

//For testing the API in local machine GET:[http://localhost:3000/dns/{url}]
func main() {
	r := gin.Default()

	router.InitializeDNSRoutes(r)

	// Run the server
	r.Run(":3000")
}
