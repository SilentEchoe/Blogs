package v1

import (
	"LearningNotes-Go/models"
	"LearningNotes-Go/pkg/setting"
	"LearningNotes-Go/pkg/util"
	"github.com/gin-gonic/gin"
)

// @Summary 获取bin文件
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/moduleNames [Get]
func GetModelBins(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	maps["modelId"] = c.PostForm("modelId")
	maps["compatibilityType"] = c.PostForm("compatibilityType")
	maps["attrKey"] = c.PostForm("attrKey")
	maps["attrValue"] = c.PostForm("attrValue")

	data["lists"] = models.GetModelNames(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetModelNameTotal(maps)

}
