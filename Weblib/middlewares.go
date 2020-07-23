package Weblib

import (
	"LearningNotes-Go/Services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitMiddleware(prodService Services.ProdService) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Keys = make(map[string]interface{})
		context.Keys["prodservice"] = prodService //赋值
		context.Next()
	}
}

// 中间件异常处理
func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				context.JSON(500, gin.H{"status": fmt.Sprintf("%s", r)})
				context.Abort()
			}
		}()
	}
}
