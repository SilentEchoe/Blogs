package v1

import (
	"LearningNotes-Go/models"
	"LearningNotes-Go/pkg/e"
	"strconv"

	/*"LearningNotes-Go/pkg/setting"
	"LearningNotes-Go/pkg/util"*/
	"github.com/gin-gonic/gin"
	"net/http"
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

	id := c.PostForm("modelId")
	/*compatibilityType := c.Query("compatibilityType")
	attrKey := c.Query("attrKey")
	attrValue := c.Query("attrValue")
	version := c.Query("version")*/

	compatibilityType := c.PostForm("compatibilityType")
	attrKey := c.PostForm("attrKey")
	attrValue := c.PostForm("attrValue")
	version := c.PostForm("version")

	// string 转换int
	modelId, err := strconv.Atoi(id)
	if err == nil {
		// 先根据modelId 和compatibility 查到model_type_id
		MadalenaTypeId := models.GetModelTypeId(modelId, compatibilityType)
		// 根据model_type_id,attrKey, attrValue, version 这四种属性查到bin模板
		BinTemplate := models.GetBin(MadalenaTypeId, attrKey, attrValue, version)

		data["BinTemplate"] = BinTemplate

		Bins := models.GetBins(BinTemplate[0].ID)
		data["Bins"] = Bins
		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
