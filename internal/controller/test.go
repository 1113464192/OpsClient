package controller

import (
	"github.com/gin-gonic/gin"
)

// Test1
// @Tags 测试相关
// @title 测试Gin能否正常访问
// @description 返回"Hello world"
// @Summary 测试Gin能否正常访问
// @Produce  application/json
// @Param ClientAuthSign header string true "格式为：运维给的签名"
// @Success 200 {} string "{"message": "Hello world"}"
// @Router /api/v1/ping1 [get]
func Test1(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}
