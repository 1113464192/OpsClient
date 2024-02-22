package controller

import (
	"ops_client/internal/service"
	"ops_client/pkg/api"
	"ops_client/pkg/logger"

	"github.com/gin-gonic/gin"
)

// GetServerPath
// @Tags 通用相关
// @title 获取机器上所有单服信息
// @description 获取机器上所有单服信息
// @Summary 获取机器上所有单服信息
// @Produce  application/json
// @Param ClientAuthSign header string true "格式为：运维给的签名"
// @Success 200 {object} api.Response "{"data":{},"meta":{msg":"Success"}}"
// @Failure 401 {object} api.Response "{"data":{}, "meta":{"msg":"错误信息", "error":"错误格式输出(如存在)"}}"
// @Failure 403 {object} api.Response "{"data":{}, "meta":{"msg":"错误信息", "error":"错误格式输出(如存在)"}}"
// @Failure 500 {object} api.Response "{"data":{}, "meta":{"msg":"错误信息", "error":"错误格式输出(如存在)"}}"
// @Router /api/v1/general/server-info [get]
func GetServerInfo(c *gin.Context) {
	result, err := service.General().GetServerInfo()
	if err != nil {
		logger.Log().Error("General", "获取机器所有服信息", err)
		c.JSON(500, api.Err("获取机器所有服信息失败", err))
		return
	}
	c.JSON(200, api.Response{
		Data: result,
		Meta: api.Meta{
			Msg: "Success",
		},
	})
}

// ExecCommand
// @Tags 通用相关
// @title 在单服执行命令
// @description 获取结果
// @Summary 在单服执行命令
// @Produce  application/json
// @Param ClientAuthSign header string true "格式为：运维给的签名"
// @Param data formData api.StringReq true "传入命令"
// @Success 200 {object} api.Response "{"data":{},"meta":{msg":"Success"}}"
// @Failure 401 {object} api.Response "{"data":{}, "meta":{"msg":"错误信息", "error":"错误格式输出(如存在)"}}"
// @Failure 403 {object} api.Response "{"data":{}, "meta":{"msg":"错误信息", "error":"错误格式输出(如存在)"}}"
// @Failure 500 {object} api.Response "{"data":{}, "meta":{"msg":"错误信息", "error":"错误格式输出(如存在)"}}"
// @Router /api/v1/general/exec-command [post]
func ExecCommand(c *gin.Context) {
	var param api.StringReq
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(500, api.ErrorResponse(err))
		return
	}

	result, err := service.General().ExecCommand(param.String)
	if err != nil {
		logger.Log().Error("General", "单机执行命令", err)
		c.JSON(500, api.Err("单机执行命令失败", err))
		return
	}
	c.JSON(200, api.Response{
		Data: result,
		Meta: api.Meta{
			Msg: "Success",
		},
	})
}
