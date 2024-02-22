package controller

import (
	"ops_client/internal/service"
	"ops_client/pkg/api"
	"ops_client/pkg/logger"

	"github.com/gin-gonic/gin"
)

// UpdateServer
// @Tags 更新相关
// @title 进行文件更新
// @description 进行文件更新任务
// @Summary 进行文件更新
// @Produce  application/json
// @Param ClientAuthSign header string true "格式为：运维给的签名"
// @Param data body api.UpdateServerReq true "传入任务ID及目录切片"
// @Success 200 {object} api.Response "{"data":{},"meta":{msg":"Success"}}"
// @Failure 401 {object} api.Response "{"data":{}, "meta":{"msg":"错误信息", "error":"错误格式输出(如存在)"}}"
// @Failure 403 {object} api.Response "{"data":{}, "meta":{"msg":"错误信息", "error":"错误格式输出(如存在)"}}"
// @Failure 500 {object} api.Response "{"data":{}, "meta":{"msg":"错误信息", "error":"错误格式输出(如存在)"}}"
// @Router /api/v1/update/server [post]
func UpdateServer(c *gin.Context) {
	var param api.UpdateServerReq
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(500, api.ErrorResponse(err))
		return
	}
	result, err := service.Update().UpdateServer(param)
	if err != nil {
		logger.Log().Error("Update", "更新文件失败", err)
		c.JSON(500, api.Err("更新文件失败", err))
		return
	}
	c.JSON(200, api.Response{
		Data: map[string]any{
			"response": result,
			"taskId":   param.Tid,
		},
		Meta: api.Meta{
			Msg: "Success",
		},
	})
}

// UpdateExec
// @Tags 更新相关
// @title 执行热更/闪断
// @description 执行最终的热更/闪断
// @Summary 执行热更/闪断
// @Produce  application/json
// @Param ClientAuthSign header string true "格式为：运维给的签名"
// @Param data body api.UpdateExecReq true "type参数：闪断:1 热更:2"
// @Success 200 {object} api.Response "{"data":{},"meta":{msg":"Success"}}"
// @Failure 401 {object} api.Response "{"data":{}, "meta":{"msg":"错误信息", "error":"错误格式输出(如存在)"}}"
// @Failure 403 {object} api.Response "{"data":{}, "meta":{"msg":"错误信息", "error":"错误格式输出(如存在)"}}"
// @Failure 500 {object} api.Response "{"data":{}, "meta":{"msg":"错误信息", "error":"错误格式输出(如存在)"}}"
// @Router /api/v1/update/exec [post]
func UpdateExec(c *gin.Context) {
	var param api.UpdateExecReq
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(500, api.ErrorResponse(err))
		return
	}
	result, err := service.Update().UpdateExec(param)
	if err != nil {
		logger.Log().Error("Update", "更新执行失败", err)
		c.JSON(500, api.Err("更新执行失败", err))
		return
	}
	c.JSON(200, api.Response{
		Data: map[string]any{
			"response": result,
			"taskId":   param.Tid,
		},
		Meta: api.Meta{
			Msg: "Success",
		},
	})
}
