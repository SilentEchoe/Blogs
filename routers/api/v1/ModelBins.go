package v1

import (
	"LearningNotes-Go/models"
	"LearningNotes-Go/pkg/e"
	/*"LearningNotes-Go/pkg/setting"
	"LearningNotes-Go/pkg/util"*/
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
	//maps := make(map[string]interface{})
	code := e.ERROR

	id := c.Query("modelId")
	compatibilityType := c.Query("compatibilityType")
	attrKey := c.Query("attrKey")
	attrValue := c.Query("attrValue")
	version := c.Query("version")
	// string 转换int
	modelId, err := strconv.Atoi(id)
	if err == nil {
		// 先根据modelId 和compatibility 查到model_type_id
		MadalenaTypeId := models.GetModelTypeId(modelId, compatibilityType)
		// 根据model_type_id,attrKey, attrValue, version 这四种属性查到bin模板
		Bins := models.GetBin(MadalenaTypeId, attrKey, attrValue, version)

		// 如果bin模板为空 代表使用的是bin文件
		if Bins[0].BinTemplate != "" {

		}

		data["base64"] = Bins

		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
