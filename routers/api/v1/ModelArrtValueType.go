package v1

import (
	"LearningNotes-Go/models"
	"LearningNotes-Go/pkg/e"
	"LearningNotes-Go/pkg/setting"
	"LearningNotes-Go/pkg/util"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 获取类型
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/GetModelValueTypes [Get]
func GetModelValueTypes(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetModelTypes(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetModelTypeTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
