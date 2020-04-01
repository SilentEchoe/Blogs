package v1

import (
	"LearningNotes-Go/models"
	"LearningNotes-Go/pkg/e"
	"LearningNotes-Go/pkg/logging"
	"LearningNotes-Go/pkg/setting"
	"LearningNotes-Go/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"net/http"
)

//获取型号名
func GetModuleName(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetModuleName(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)

	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
