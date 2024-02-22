package middleware

import (
	"bytes"
	"fmt"
	"io"
	"ops_client/pkg/logger"

	"github.com/gin-gonic/gin"
)

func RecordLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			var body []byte
			var err error
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				logger.Log().Error("recordLog", "记录行为失败", err)
			} else {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
			recordLog := fmt.Sprintf("IP: %s  Method: %s  Path: %s  Body: %s", c.ClientIP(), c.Request.Method, c.Request.URL.Path, string(body))
			logger.Log().Info("recordLog", "记录行为", recordLog)
		} else {
			c.Next()
		}
	}
}
