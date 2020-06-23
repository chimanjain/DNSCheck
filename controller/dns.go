package controller

import (
	"DNSCheck/model"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

var mxEntity model.MxEntity

var dns model.Dns

//FindIP Fetches the IP records
func FindIP(url string) {
	iprecords, _ := net.LookupIP(url)
	for _, ip := range iprecords {
		dns.IPAddress = append(dns.IPAddress, ip)
	}
}

//FindCName Fetches the CName records
func FindCName(url string, cCName chan string) {
	cname, _ := net.LookupCNAME(url)
	cCName <- cname
}

//FindNS Fetches the NS records
func FindNS(url string) {
	nameserver, _ := net.LookupNS(url)
	for _, ns := range nameserver {
		dns.NS = append(dns.NS, *&ns.Host)
	}
}

//FindMX Fetches the MX records
func FindMX(url string) {
	mxrecords, _ := net.LookupMX(url)
	for _, mx := range mxrecords {
		mxEntity.Host, mxEntity.Pref = mx.Host, mx.Pref
		dns.MX = append(dns.MX, mxEntity)
	}
}

//FindTXT Fetches the TXT records
func FindTXT(url string) {
	txtrecords, _ := net.LookupTXT(url)
	for _, txt := range txtrecords {
		dns.TXT = append(dns.TXT, txt)
		//fmt.Printf("%T\n", txt)
	}
}

//FindDNS function is used for fetching the IP, CName, NS, MX and TXT records of the URL
func FindDNS(c *gin.Context) {

	encodedURL := url.QueryEscape(c.Param("url"))
	dns.URL = c.Param("url")

	go FindIP(encodedURL)

	cCName := make(chan string)
	go FindCName(encodedURL, cCName)
	dns.Cname = <-cCName

	go FindNS(encodedURL)

	go FindMX(encodedURL)

	go FindTXT(encodedURL)

	time.Sleep(1 * time.Second)
	c.JSON(http.StatusOK, gin.H{"data": dns})
}
