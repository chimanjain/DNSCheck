package controller

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/chimanjain/dnscheck/service"
	"github.com/gofiber/fiber/v2"
)

// FindDNS function is used for fetching the IP, CName, NS, MX and TXT records of the URL
func GetDNS(c *fiber.Ctx) error {
	encodedURL := url.QueryEscape(strings.TrimSpace(c.Params("url")))
	dnsresp := service.FetchDNS(encodedURL)
	return c.Status(http.StatusOK).JSON(dnsresp)
}
