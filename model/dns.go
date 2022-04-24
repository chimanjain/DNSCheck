package model

import "net"

type Dns struct {
	URL       string     `json:"url"`
	IPAddress []net.IP   `json:"IPAddress"`
	Cname     string     `json:"Cname"`
	NS        []string   `json:"NS"`
	MX        []MxEntity `json:"MX"`
	TXT       []string   `json:"TXT"`
}

type MxEntity struct {
	Host string
	Pref uint16
}
