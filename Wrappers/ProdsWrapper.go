package Wrappers

import (
	"context"

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
	}, nil)
}

func NewProdsWrapper(c client.Client) client.Client {
	return &ProdsWrapper{c}
}
