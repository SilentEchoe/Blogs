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

// @Summary 获取型号名
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func GetModuleName(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	code := e.SUCCESS
	data["lists"] = models.GetModuleName(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetModuleNameTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// @Summary 新增型号名
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddModuleName(c *gin.Context) {
	moduleName := c.Query("moduleName")
	parentId := com.StrTo(c.DefaultQuery("parentId", "0")).MustInt()
	isDelete := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	code := e.INVALID_PARAMS
	// 判断是否存在
	models.AddModuleName(moduleName, parentId, isDelete)
	code = e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}
