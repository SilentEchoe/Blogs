package main

import (
	"LearningNotes-Go/Servicelmpl"

	Service "LearningNotes-Go/Services"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	consulReg := consul.NewRegistry( //新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
		registry.Addrs("127.0.0.1:8500"),
	)

	prodService := micro.NewService(
		micro.Name("prodservice"),
		micro.Address(":8011"),
		micro.Registry(consulReg),
	)
	prodService.Init()
	Service.RegisterProdServiceHandler(prodService.Server(), new(ServiceImpl.ProdService))
	prodService.Run()
}
