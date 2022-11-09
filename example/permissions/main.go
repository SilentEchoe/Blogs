package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	e   *casbin.Enforcer
	err error
)

func init() {
	e, err = casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		logrus.Fatal("load file failed, %v", err.Error())
	}
}

func main() {
	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		sub := ctx.Query("username")
		obj := ctx.Request.URL.Path
		act := ctx.Request.Method
		checkPermission(ctx, sub, obj, act)
	})
	r.POST("/users", func(ctx *gin.Context) {
		sub := ctx.Query("username")
		obj := ctx.Request.URL.Path
		act := ctx.Request.Method
		checkPermission(ctx, sub, obj, act)
	})

	r.Run()
}

func checkPermission(ctx *gin.Context, sub, obj, act string) {
	logrus.Infof("sub = %s obj = %s act = %s", sub, obj, act)
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		logrus.Print("enforce failed %s", err.Error())
		ctx.String(http.StatusInternalServerError, "内部服务器错误")
		return
	}
	if !ok {
		logrus.Println("权限验证不通过")
		ctx.String(http.StatusOK, "权限验证不通过")
		return
	}
	logrus.Println("权限验证通过")
	ctx.String(http.StatusOK, "权限验证通过")
}
