package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func EncryptClientAuthSign(ip string, authKey string) (sign string, err error) {
	builder := strings.Builder{}
	builder.WriteString(authKey)
	builder.WriteString(ip)
	md5Hash := md5.Sum([]byte(builder.String()))
	sign = hex.EncodeToString(md5Hash[:])
	return sign, err
}
