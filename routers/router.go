package routers

import (
	"LearningNotes-Go/pkg/setting"
	v1 "LearningNotes-Go/routers/api"
	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter() *gin.Engine  {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags",v1.GetTags)
		apiv1.POST("/tags",v1.AddTag)
		apiv1.PUT("/tags",v1.EditTag)
		apiv1.DELETE("/tags",v1.DeleteTag)
	}
	return r
}