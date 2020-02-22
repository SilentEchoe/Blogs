package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main(){
	router := gin.Default()
	router.GET("/user/:name/:password", ginHandler) //冒号加上一个参数名组成路由参数
	router.Run(":8000")
}
func ginHandler(c *gin.Context) {
	name := c.Param("name")   //读取name的值
	pwd := c.Param("password")
	c.String(http.StatusOK, " Hello  %s,%s", name, pwd)     //http.StatusOK返回状态码
}