package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/chimanjain/dnscheck/cache"
	"github.com/chimanjain/dnscheck/model"
	"github.com/go-redis/redis/v8"
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

func FetchDNS(url string) model.Dns {
	redisClient := cache.GetRedisClient()
	val, err := redisClient.GetDNS(ctx, url)
	if err == redis.Nil {
		fmt.Println(url, " record does not exist")
	}
	if val.URL != "" {
		return val
	}

	wg.Add(4)
	dns.URL = url

	go FindIP(url, &wg)

	// Channel for experimental purpose
	cCName := make(chan string)
	go FindCName(url, cCName)
	dns.Cname = <-cCName
	close(cCName)

	go FindNS(url, &wg)

	go FindMX(url, &wg)

	go FindTXT(url, &wg)

	wg.Wait()

	err = redisClient.SetDNS(ctx, dns)
	if err != nil {
		log.Println(err)
	}
	return dns
}
