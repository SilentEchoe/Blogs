package Weblib

import (
	"LearningNotes-Go/Services"
	"github.com/gin-gonic/gin"
)

func NewGinRouter(prodService Services.ProdService) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(InitMiddleware(prodService))
	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", GetProdsList)
	}
	return ginRouter
}
