package main

import (
	"LearningNotes-Go/Services"
	"LearningNotes-Go/Weblib"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}

func main() {
	consulReg := consul.NewRegistry( //新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
		registry.Addrs("127.0.0.1:8500"),
	)
	myService := micro.NewService(
		micro.Name("prodservice.client"),
		micro.WrapClient(NewLogWrapper),
	)

	prodService := Services.NewProdService("prodservice", myService.Client())

	httpServer := web.NewService(
		web.Name("httpprodservice"), //注册进consul服务中的service名字
		web.Address(":8001"),        //注册进consul服务中的端口,也是这里我们gin的server地址
		web.Handler(Weblib.NewGinRouter(prodService)),
		web.Registry(consulReg),
	)

	httpServer.Init() //加了这句就可以使用命令行的形式去设置我们一些启动的配置
	httpServer.Run()
}
