package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/astaxie/beego/validation"
	"github.com/Unknwon/com"

	"LearningNotes-Go/pkg/e"
	"LearningNotes-Go/models"
	"LearningNotes-Go/pkg/util"
	"LearningNotes-Go/pkg/setting"
)

//获取多个文章标签
func GetTags(c *gin.Context) {

	// c.Query 用户获取?name=test Url的参数
	name := c.Query("name")

	// 这里创建了两个string类型的map interface代表空接口
	// 因为所有类型都能实现空接口 所以Key是string类型 values 为任意类型

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}