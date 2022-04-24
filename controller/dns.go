package controller

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"

	"github.com/chimanjain/dnscheck/cache"
	"github.com/chimanjain/dnscheck/model"
	"github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"
)

var (
	mxEntity model.MxEntity
	dns      model.Dns
	wg       sync.WaitGroup
)

var ctx = context.Background()

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
	encodedURL := url.QueryEscape(c.Param("url"))
	redisClient := cache.GetRedisClient()
	val, err := redisClient.GetDNS(ctx, encodedURL)
	if err == redis.Nil {
		fmt.Println(encodedURL, " record does not exist")
	}
	if val.URL != "" {
		c.JSON(http.StatusOK, val)
		return
	}

	wg.Add(4)
	dns.URL = encodedURL

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

	err = redisClient.SetDNS(ctx, dns)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, dns)
}
