package crontab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"ops_client/configs"
	"ops_client/internal/consts"
	"ops_client/internal/service"
	"ops_client/pkg/logger"
	"ops_client/pkg/util"
)

func CronCalculationGameSum() {
	cmd := configs.Conf.CustomCmd.GameSumCmd
	result, err := service.General().ExecCommand(cmd)
	if err != nil {
		logger.Log().Error("Cron", "CronCalculationGameSum执行错误", err)
		return
	}
	// 向服务端发送游戏服数量，假设服务端IP为192.168.18.88，路径为/api/v1/client/gameSum，无需验证信息，请求方式为post
	// Create a map to hold the POST data
	data := map[string]string{
		"game_sum": result.Response,
	}

	// Convert the map to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Log().Error("Cron", "CronCalculationGameSum Json编码报错", err)
		return
	}

	// Send the POST request

	sslStr := "https"
	if configs.Conf.ServerSide.IsSSL == "false" {
		sslStr = "http"
	}

	ipValue := configs.Conf.ServerSide.Domain
	if configs.Conf.ServerSide.Domain == "" && configs.Conf.ServerSide.Ip != "" {
		ipValue = configs.Conf.ServerSide.Ip
	}

	httpReqStr := fmt.Sprintf("%s://%s:%s/%s", sslStr, ipValue, configs.Conf.ServerSide.Port, consts.ServerGameSumApiPath)
	// Create a new request
	req, err := http.NewRequest("POST", httpReqStr, bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Log().Error("Cron", "New request error", err)
		return
	}

	// Add the ClientAuth header
	localIp, err := util.GetLocalIp(configs.Conf.CustomCmd.LocalIpCmd, configs.Conf.CustomCmd.LocalIpApi)
	if err != nil {
		logger.Log().Error("Cron", "获取本机IP失败", err)
		return
	}

	sign, err := util.EncryptClientAuthSign(localIp, configs.Conf.Auth.Key)
	if err != nil {
		logger.Log().Error("Cron", "加密生产失败", err)
		return
	}
	req.Header.Set("ClientAuth", sign)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Log().Error("Cron", "POST request error", err)
		return
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusOK {
		logger.Log().Error("Cron", "Server responded with non-OK status", resp.StatusCode)
	}

}
