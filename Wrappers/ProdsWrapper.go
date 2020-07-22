package Wrappers

import (
	"LearningNotes-Go/Services"
	"context"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
)

type ProdsWrapper struct {
	client.Client
}

func (l *ProdsWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()

	configA := hystrix.CommandConfig{
		Timeout: 1000,
	}
	hystrix.ConfigureCommand(cmdName, configA)
	return hystrix.Do(cmdName, func() error {
		return l.Client.Call(ctx, req, rsp)
	}, func(err error) error {

		defaultProds(rsp)
		return nil
	})
}

func NewProdsWrapper(c client.Client) client.Client {
	return &ProdsWrapper{c}
}

func newProd(id int32, pname string) *Services.ProdModel {
	return &Services.ProdModel{ProdID: id, ProdName: pname}
}

// 熔断的降级方法
// 降级方法尽量不要有异常，且最好不需要连接数据库，可从Redis 或文本文件读取数据
func defaultProds(rsp interface{}) {
	models := make([]*Services.ProdModel, 0)
	var i int32
	for i = 0; i < 5; i++ {
		models = append(models, newProd(20+i, "prodname"+strconv.Itoa(20+int(i))))
	}
	result := rsp.(*Services.ProdListResponse)
	result.Data = models
}
