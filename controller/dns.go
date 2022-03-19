package controller

import (
	"dnscheck/model"
	"net"
	"net/http"
	"net/url"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	mxEntity model.MxEntity
	dns      model.Dns
	wg       sync.WaitGroup
)

//FindIP Fetches the IP records
func FindIP(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	iprecords, _ := net.LookupIP(url)
	dns.IPAddress = append(dns.IPAddress, iprecords...)
}

//FindCName Fetches the CName records
func FindCName(url string, cCName chan string) {

	cname, _ := net.LookupCNAME(url)
	cCName <- cname
}

//FindNS Fetches the NS records
func FindNS(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	nameserver, _ := net.LookupNS(url)
	for _, ns := range nameserver {
		dns.NS = append(dns.NS, ns.Host)
	}
}

//FindMX Fetches the MX records
func FindMX(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	mxrecords, _ := net.LookupMX(url)
	for _, mx := range mxrecords {
		mxEntity.Host, mxEntity.Pref = mx.Host, mx.Pref
		dns.MX = append(dns.MX, mxEntity)
	}
}

//FindTXT Fetches the TXT records
func FindTXT(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	txtrecords, _ := net.LookupTXT(url)
	dns.TXT = append(dns.TXT, txtrecords...)
}

//FindDNS function is used for fetching the IP, CName, NS, MX and TXT records of the URL
func FindDNS(c *gin.Context) {
	wg.Add(4)
	encodedURL := url.QueryEscape(c.Param("url"))
	dns.URL = c.Param("url")

	go FindIP(encodedURL, &wg)

	// Channel for experimental purpose
	cCName := make(chan string)
	go FindCName(encodedURL, cCName)
	dns.Cname = <-cCName
	close(cCName)

	go FindNS(encodedURL, &wg)

	go FindMX(encodedURL, &wg)

	go FindTXT(encodedURL, &wg)

	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"data": dns})
}
