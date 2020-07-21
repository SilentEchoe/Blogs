package Weblib

import (
	"LearningNotes-Go/Services"
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

/*
func GetProdsList(ginCtx *gin.Context) {
	prodService := ginCtx.Keys["prodservice"].(Services.ProdService)
	var prodReq Services.ProdsRequest
	err := ginCtx.Bind(&prodReq)
	if err != nil {
		ginCtx.JSON(500, gin.H{"status": err.Error()})
	} else {
		prodRes, _ := prodService.GetProdsList(context.Background(), &prodReq)
		ginCtx.JSON(200, gin.H{"data": prodRes.Data})
	}
}*/

func GetProdsList(ginCtx *gin.Context) {
	prodService := ginCtx.Keys["prodservice"].(Services.ProdService)
	var prodReq Services.ProdsRequest
	err := ginCtx.Bind(&prodReq)
	if err != nil {
		ginCtx.JSON(500, gin.H{"status": err.Error()})
	} else {
		//熔断代码改造
		configA := hystrix.CommandConfig{
			Timeout: 1000,
		}
		//配置command
		hystrix.ConfigureCommand("getprods", configA)

		// 执行使用Do 方法
		var prodRes *Services.ProdListResponse

		err := hystrix.Do("getprods", func() error {
			prodRes, err = prodService.GetProdsList(context.Background(), &prodReq)
			return err
		}, nil)

		if err != nil {
			ginCtx.JSON(500, gin.H{"status": err.Error()})
		} else {
			ginCtx.JSON(200, gin.H{"data": prodRes.Data})
		}

	}
}
