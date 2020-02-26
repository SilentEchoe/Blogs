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
		// 获取标签
		apiv1.GET("/tags",v1.GetTags)
		// 添加标签
		apiv1.POST("/tags",v1.AddTag)
		// 修改标签
		apiv1.PUT("/tags",v1.EditTag)
		// 删除标签
		apiv1.DELETE("/tags",v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return r
}