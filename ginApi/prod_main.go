package main

import (
	"LearningNotes-Go/Services"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	consulReg := consul.NewRegistry( //新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
		registry.Addrs("127.0.0.1:8500"),
	)



	myService := micro.NewService(micro.Name("prodservice.client"))

	prodService := Services.NewProdService("prodservice",myService.Client()))

	httpServer := web.NewService(
		web.Name("httpprodservice"), //注册进consul服务中的service名字
		web.Address(":8001"), //注册进consul服务中的端口,也是这里我们gin的server地址
		web.Handler(Weblib.NewGinRouter(prodService)),
		web.Registry(consulReg),
	)
	httpServer.Init() //加了这句就可以使用命令行的形式去设置我们一些启动的配置
	httpServer.Run()
}
