package service

import (
	"fmt"
	"ops_client/pkg/api"
	"os"
	"os/exec"
	"sync"
)

type ShellService struct {
}

var (
	insShell = &ShellService{}
)

func Shell() *ShellService {
	return insShell
}

// 执行shell
func (s *ShellService) ExecShell(cmdStr string, serverDir string, env []string, ch chan *api.ShellRes, wg *sync.WaitGroup) {
	var result api.ShellRes
	result.ServerDir = serverDir
	cmd := exec.Command("bash", "-c", cmdStr) // 连贯可以用"bash", "-c", "command"
	cmd.Env = append(os.Environ(), env...)
	output, err := cmd.CombinedOutput()
	if err != nil || cmd.ProcessState.ExitCode() != 0 {
		result.Status = cmd.ProcessState.ExitCode()
		result.Response = fmt.Sprintf("错误返回: %s\n%s", string(output), err.Error())
		ch <- &result
		wg.Done()
		return
	}
	result.Response = string(output)
	result.Status = cmd.ProcessState.ExitCode()
	ch <- &result
	wg.Done()
}
