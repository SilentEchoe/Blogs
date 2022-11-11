package main

import (
	gorm_mysql "LearningNotes-Go/example/permissions/initialize"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var (
	GVA_DB *gorm.DB
	e      *casbin.Enforcer
	err    error
)

func main() {
	GVA_DB = gorm_mysql.GormMysql()
	if GVA_DB == nil {
		log.Fatalln("错误")
	}

	Casbin()

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

	// 程序结束前关闭数据库连接
	db, _ := GVA_DB.DB()
	defer db.Close()

	r.Run()

}

func Casbin() {
	a, err := gormadapter.NewAdapterByDB(GVA_DB)
	if a == nil || err != nil {
		log.Fatalln("连接数据库失败:", err)
	}
	e, err = casbin.NewEnforcer("model.conf", a)

	if err != nil {
		log.Fatalln("casbin 认证失败:", err)
	}

	e.LoadPolicy()
	//e.AddPolicy("alice", "/users", "DELETE")
	e.SavePolicy()
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
