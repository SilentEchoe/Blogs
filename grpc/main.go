package main

import (
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	consulReg := consul.NewRegistry( //新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
		registry.Addrs("127.0.0.1:8500"),
	)

}
