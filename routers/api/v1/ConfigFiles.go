package v1

import (
	"LearningNotes-Go/models"
	"LearningNotes-Go/pkg/e"
	"strconv"
	"strings"

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
func GetConfigFiles(c *gin.Context) {

	data := make(map[string]interface{})
	//maps := make(map[string]interface{})
	code := e.ERROR
	// 获取modelId
	id := c.Query("modelId")
	// 转换成int类型
	modelId, err := strconv.Atoi(id)
	code = e.SUCCESS
	var modelsID []int
	if err == nil {
		// 查询配置文件Id
		var configFilesid = models.GetConfigsById(modelId)
		configFiles := strings.Split(configFilesid, ",")

		for _, v := range configFiles {
			id, err := strconv.Atoi(v)
			if err == nil {
				modelsID = append(modelsID, id)
			}

		}
		configs := models.GetConfigFileById(modelsID)
		data["configFiles"] = configs
		code = e.SUCCESS

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
