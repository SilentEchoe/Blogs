package main

import (
	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
	"log"
)

type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:512;uniqueIndex:unique_index"`
	V0    string `gorm:"size:512;uniqueIndex:unique_index"`
	V1    string `gorm:"size:512;uniqueIndex:unique_index"`
	V2    string `gorm:"size:512;uniqueIndex:unique_index"`
	V3    string `gorm:"size:512;uniqueIndex:unique_index"`
	V4    string `gorm:"size:512;uniqueIndex:unique_index"`
	V5    string `gorm:"size:512;uniqueIndex:unique_index"`
}

func main() {
	// Increase the column size to 512.

	a, _ := xormadapter.NewAdapter("mysql", "root:root@tcp(localhost:3306)/")

	e, _ := casbin.NewEnforcer("model.conf", a)

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	e.Enforce("alice", "data1", "read")

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	e.SavePolicy()

	sub := "alice"  // 想要访问资源的用户
	obj := "/users" // 将要被访问的资源
	act := "GET"    // 用户对资源实施的操作

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// 处理错误
		log.Fatalln("错误:", err)
	}

	if ok == true {
		// 允许 alice 读取 data1
		log.Println("验证成功")
	} else {
		// 拒绝请求，抛出异常
		log.Println("验证失败，拒绝请求")
	}

}
