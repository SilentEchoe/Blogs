package main

import "github.com/go-ping/ping"

func main() {
	pingDevices()
}

func pingDevices() {
	pinger, err := ping.NewPinger("www.baidu.com")
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	pinger.SetPrivileged(true)
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/rtt stats
	println("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	println("状态为：", stats.PacketsSent)
}
