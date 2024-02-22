package middleware

import (
	"github.com/gin-gonic/gin"
	"net"
	"ops_client/configs"
	"ops_client/internal/consts"
	"ops_client/pkg/api"
	"ops_client/pkg/util"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("ClientAuthSign")
		if authHeader == "" {
			c.JSON(401, api.Err("访问未携带签名", nil))
			c.Abort()
			return
		}
		serverIp := c.ClientIP()

		var (
			localIp string
			err     error
		)

		if localIp, err = util.GetLocalIp(configs.Conf.CustomCmd.LocalIpCmd, configs.Conf.CustomCmd.LocalIpApi); err != nil {
			c.JSON(500, api.Err("获取本机IP失败", err))
			c.Abort()
			return
		}
		sign, err := util.EncryptClientAuthSign(localIp, configs.Conf.Auth.Key)
		if err != nil {
			c.JSON(500, api.Err("获取本机IP失败", err))
			c.Abort()
			return
		}
		if authHeader != sign {
			c.JSON(403, api.Err("签名错误", nil))
			c.Abort()
			return
		}
		_, ipNet, _ := net.ParseCIDR(configs.Conf.ServerSide.AllowCidr)
		_, ipNetA, _ := net.ParseCIDR(consts.CIDRTypeA)
		_, ipNetB, _ := net.ParseCIDR(consts.CIDRTypeB)
		_, ipNetC, _ := net.ParseCIDR(consts.CIDRTypeC)
		ip := net.ParseIP(serverIp)
		if serverIp != "127.0.0.1" && serverIp != configs.Conf.ServerSide.Ip && !ipNet.Contains(ip) && !ipNetA.Contains(ip) && !ipNetB.Contains(ip) && !ipNetC.Contains(ip) {
			c.JSON(403, api.Err("您的IP不在许可IP中", nil))
			c.Abort()
			return
		}
		c.Next()
	}
}
