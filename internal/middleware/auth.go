package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"net"
	"ops_client/configs"
	"ops_client/internal/consts"
	"ops_client/pkg/api"
	"ops_client/pkg/util"
	"strings"

	"github.com/gin-gonic/gin"
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

		if localIp, err = util.GetLocalIp(); err != nil {
			c.JSON(500, api.Err("获取本机IP失败", err))
			c.Abort()
			return
		}
		sign, err := Md5EncryptSign(localIp)
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
		_, ipNet, _ := net.ParseCIDR(configs.Conf.ServerIp.AllowCidr)
		_, ipNetA, _ := net.ParseCIDR(consts.CIDRTypeA)
		_, ipNetB, _ := net.ParseCIDR(consts.CIDRTypeB)
		_, ipNetC, _ := net.ParseCIDR(consts.CIDRTypeC)
		ip := net.ParseIP(serverIp)
		if serverIp != "127.0.0.1" && serverIp != configs.Conf.ServerIp.Value && !ipNet.Contains(ip) && !ipNetA.Contains(ip) && !ipNetB.Contains(ip) && !ipNetC.Contains(ip) {
			c.JSON(403, api.Err("您的IP不在许可IP中", nil))
			c.Abort()
			return
		}
		c.Next()
	}
}

func Md5EncryptSign(ip string) (sign string, err error) {
	builder := strings.Builder{}
	builder.WriteString(configs.Conf.Auth.Key)
	builder.WriteString(ip)
	md5Hash := md5.Sum([]byte(builder.String()))
	sign = hex.EncodeToString(md5Hash[:])
	return sign, err
}
