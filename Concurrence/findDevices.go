package main

import (
	"github.com/go-ping/ping"
)

func main() {
	pinger, err := ping.NewPinger("www.baidu.com")
	if err != nil {
		panic(err)
	}
	pinger.SetPrivileged(true)
	pinger.Count = 3
	err = pinger.Run() // blocks until finished
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/rtt stats
	println(stats)
}
