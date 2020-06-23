package model

import "net"

type Dns struct {
	URL       string `json:"url" gorm:"primary_key"`
	IPAddress []net.IP
	Cname     string
	NS        []string
	MX        []MxEntity
	TXT       []string
}

type MxEntity struct {
	Host string
	Pref uint16
}
