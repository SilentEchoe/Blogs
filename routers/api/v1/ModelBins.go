package v1

import (
	"LearningNotes-Go/models"
	"LearningNotes-Go/pkg/e"
	"LearningNotes-Go/pkg/setting"
	"LearningNotes-Go/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary 获取bin文件
// @Produce  json
// @Param name query string true "Name"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/GetModelBins [Get]
func GetModelBins(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	code := e.ERROR_EXIST_TAG
	// 先根据modelId 和compatibility 查到model_type_id
	modelId, err := strconv.Atoi(c.PostForm("modelId"))
	compatibilityType := c.PostForm("compatibilityType")
	if err != nil {
		MadalenaTypeId := models.GetModelTypeId(modelId, compatibilityType)
		maps["MadalenaTypeId"] = MadalenaTypeId
		maps["attrKey"] = c.PostForm("attrKey")
		maps["attrValue"] = c.PostForm("attrValue")
		maps["Version"] = c.PostForm("version")
		models.GetBinTemplate(maps)

		data["lists"] = models.GetBin(util.GetPage(c), setting.PageSize, maps)
		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
