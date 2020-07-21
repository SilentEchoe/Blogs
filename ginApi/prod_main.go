package main

import (
	"LearningNotes-Go/Services"
	"LearningNotes-Go/Weblib"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	consulReg := consul.NewRegistry( //新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
		registry.Addrs("127.0.0.1:8500"),
	)
	/*ginRouter := gin.Default()*/

	myService := micro.NewService(micro.Name("prodservice.client"))

	prodService := Services.NewProdService("prodservice", myService.Client())

	/*v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", func(ginCtx *gin.Context) {
			var prodReq Services.ProdsRequest
			err := ginCtx.Bind(&prodReq)

			if err != nil {
				ginCtx.JSON(500, gin.H{"status": err.Error()})
			} else {
				prodRes, _ := prodService.GetProdsList(context.Background(), &prodReq)
				ginCtx.JSON(200, gin.H{"data": prodRes.Data})
			}

		})

	}*/

	httpServer := web.NewService(
		web.Name("httpprodservice"), //注册进consul服务中的service名字
		web.Address(":8001"),        //注册进consul服务中的端口,也是这里我们gin的server地址
		web.Handler(Weblib.NewGinRouter(prodService)),
		web.Registry(consulReg),
	)

	httpServer.Init() //加了这句就可以使用命令行的形式去设置我们一些启动的配置
	httpServer.Run()
}
