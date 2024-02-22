package service

import (
	"fmt"
	"ops_client/configs"
	"ops_client/pkg/api"
	"ops_client/pkg/util"
	"os/exec"
	"strings"
	"sync"
)

type GeneralService struct {
}

var (
	insGeneral = &GeneralService{}
)

func General() *GeneralService {
	return insGeneral
}

// 获取单服信息
func (s *GeneralService) GetServerHtmlInfo(cmdTem api.GetServerInfoCmdTem, path string, ch chan *api.GetServerInfoRes, wg *sync.WaitGroup) {
	flagCmdStr := fmt.Sprintf(cmdTem.FlagCmdTem, path)
	nameCmdStr := fmt.Sprintf(cmdTem.NameCmdTem, path)
	flagCmd := exec.Command("bash", "-c", flagCmdStr)
	nameCmd := exec.Command("bash", "-c", nameCmdStr)

	res := api.GetServerInfoRes{
		Path: api.CmdResultRes{
			Response: path,
			Status:   0,
		},
		Flag: api.CmdResultRes{
			Status: 0,
		},
		Name: api.CmdResultRes{
			Status: 0,
		},
	}

	var output []byte
	var err error
	// 取flag的值
	output, err = flagCmd.CombinedOutput()
	res.Flag.Response = string(output)
	// 因为可能name能取出，因此错误也不return
	if err != nil || flagCmd.ProcessState.ExitCode() != 0 {
		res.Flag.Status = flagCmd.ProcessState.ExitCode()
	}
	// 取name的值
	output, err = nameCmd.CombinedOutput()
	res.Name.Response = string(output)
	if err != nil || nameCmd.ProcessState.ExitCode() != 0 {
		res.Name.Status = nameCmd.ProcessState.ExitCode()
	}
	ch <- &res
	wg.Done()
}

func (s *GeneralService) verdictGameType(paths []string) (cmdTem api.GetServerInfoCmdTem) {
	var cBool bool
	for _, value := range configs.Conf.CGame.Values {
		cBool = util.StringSliceContain(paths, value)
		if cBool {
			break
		}
	}
	if !cBool {
		cmdTem = api.GetServerInfoCmdTem{
			FlagCmdTem: `awk -F\' '/flag/{printf $4}' %s/html/hhsy/single/config/base.php`,
			NameCmdTem: `awk -F\' '/name/{printf $4}' %s/html/hhsy/single/config/base.php`,
		}
	} else {
		cmdTem = api.GetServerInfoCmdTem{
			FlagCmdTem: `awk -F\' '/flag/{prinf $4}' %s/html/hhsy/config/single.php | head -n 1`,
			NameCmdTem: `awk -F\' '/name/{prinf $4}' %s/html/hhsy/config/single.php | head -n 1`,
		}
	}
	return cmdTem
}

// 获取单机上所有服
func (s *GeneralService) GetServerInfo() (*[]api.GetServerInfoRes, error) {
	var err error
	pathCmd := exec.Command("bash", "-c", "find /data -maxdepth 1 -mindepth 1 -type d -regex .*?_.*?_[a-Z][0-9]+[a-Z]")
	output, err := pathCmd.CombinedOutput()
	if err != nil || pathCmd.ProcessState.ExitCode() != 0 {
		return nil, fmt.Errorf("获取所有服错误: %v", err)
	}
	splitPath := strings.Split(string(output), "\n")
	paths := splitPath[:len(splitPath)-1]

	// 判断是否C++游戏
	cmdTem := s.verdictGameType(paths)
	// 命令采集
	channel := make(chan *api.GetServerInfoRes, len(paths))
	wg := sync.WaitGroup{}
	for _, path := range paths {
		wg.Add(1)
		go s.GetServerHtmlInfo(cmdTem, path, channel, &wg)
	}
	wg.Wait()
	close(channel)
	var result []api.GetServerInfoRes
	for res := range channel {
		result = append(result, *res)
	}
	return &result, err
}

// 单机执行命令
func (s *GeneralService) ExecCommand(command string) (result api.ShellRes, err error) {
	cmd := exec.Command("bash", "-c", command) // 连贯可以用"bash", "-c", "command"
	output, err := cmd.CombinedOutput()
	if err != nil || cmd.ProcessState.ExitCode() != 0 {
		result.Status = cmd.ProcessState.ExitCode()
		result.Response = fmt.Sprintf("错误返回: %s\n%s", string(output), err.Error())
		return result, fmt.Errorf("执行错误: %v", err)
	}
	result.Response = string(output)
	result.Status = cmd.ProcessState.ExitCode()
	return result, err
}
