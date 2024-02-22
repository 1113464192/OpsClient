package service

import (
	"fmt"
	"ops_client/pkg/api"
	"os/exec"
)

type GeneralService struct {
}

var (
	insGeneral = &GeneralService{}
)

func General() *GeneralService {
	return insGeneral
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
