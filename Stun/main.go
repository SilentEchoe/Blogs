package main

import (
	"github.com/ccding/go-stun/stun"
)

const (
	ServiceAddress  = "10.32.130.131"
	ServiceHostName = "Service"
	ServiceHost     = 3478
	SoftwareName    = "TestStunClient"
)

func main() {
	conn := stun.NewClient()
	conn.SetServerAddr(ServiceAddress)
	conn.SetServerHost(ServiceHostName, ServiceHost)
	conn.SetSoftwareName(SoftwareName)
	conn.Discover()
}
