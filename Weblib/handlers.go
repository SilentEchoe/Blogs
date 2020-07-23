package Weblib

import (
	"LearningNotes-Go/Services"
	"context"
	"github.com/gin-gonic/gin"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetProdDetail(ginCtx *gin.Context) {
	var prodReq Services.ProdsRequest
	PanicIfError(ginCtx.Bind(&prodReq))
	prodService := ginCtx.Keys["prodservice"].(Services.ProdService)
	resp, _ := prodService.GetProdsDetail(context.Background(), &prodReq)
	ginCtx.JSON(200, gin.H{"data": resp.Data})
}

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
}
