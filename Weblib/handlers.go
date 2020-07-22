package Weblib

import (
	"LearningNotes-Go/Services"
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"strconv"
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

func newProd(id int32, pname string) *Services.ProdModel {
	return &Services.ProdModel{ProdID: id, ProdName: pname}
}

// 熔断的降级方法
// 降级方法尽量不要有异常，且最好不需要连接数据库，可从Redis 或文本文件读取数据
func defaultProds() (*Services.ProdListResponse, error) {
	models := make([]*Services.ProdModel, 0)
	var i int32
	for i = 0; i < 5; i++ {
		models = append(models, newProd(100+i, "prodname"+strconv.Itoa(100+int(i))))
	}
	res := &Services.ProdListResponse{}
	res.Data = models
	return res, nil
}

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
		}, func(err error) error {

			prodRes, err = defaultProds()
			return err
		})

		if err != nil {
			ginCtx.JSON(500, gin.H{"status": err.Error()})
		} else {
			ginCtx.JSON(200, gin.H{"data": prodRes.Data})
		}

	}
}
