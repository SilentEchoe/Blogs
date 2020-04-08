package v1

import (
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

}
