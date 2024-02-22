package service

import (
	"fmt"
	"ops_client/pkg/api"
	"os/exec"
	"sync"
)

type UpdateService struct {
}

var (
	insUpdate = &UpdateService{}
)

func Update() *UpdateService {
	return insUpdate
}

// 执行更新
func (s *UpdateService) UpdateServer(param api.UpdateServerReq) (result []api.ShellRes, err error) {
	// SVN仓库更新
	cmd := exec.Command("bash", "/data/common_cron/record_update.sh", param.PackageDir)
	var output []byte
	if output, err = cmd.CombinedOutput(); err != nil || cmd.ProcessState.ExitCode() != 0 {
		return nil, fmt.Errorf("更新SVN仓库报错: %s\n%v", string(output), err)
	}

	wg := sync.WaitGroup{}
	channel := make(chan *api.ShellRes, len(param.ServerDir))
	for i := 0; i < len(param.ServerDir); i++ {
		wg.Add(1)
		cmdString := fmt.Sprintf("bash /data/%s/server/update_server.sh %s", param.ServerDir[i], param.PackageDir)
		// goroutine执行多目录更新
		go Shell().ExecShell(cmdString, param.ServerDir[i], nil, channel, &wg)
	}
	wg.Wait()
	close(channel)
	for res := range channel {
		result = append(result, *res)
	}
	return result, err
}

func (s *UpdateService) UpdateExec(param api.UpdateExecReq) (result []api.ShellRes, err error) {
	wg := sync.WaitGroup{}
	channel := make(chan *api.ShellRes, len(param.ServerDir))
	for i := 0; i < len(param.ServerDir); i++ {
		wg.Add(1)
		cmdString := fmt.Sprintf("bash /data/%s/server/update_exec.sh %d", param.ServerDir[i], param.Type)
		// goroutine执行多目录更新
		go Shell().ExecShell(cmdString, param.ServerDir[i], nil, channel, &wg)
	}
	wg.Wait()
	close(channel)
	for res := range channel {
		result = append(result, *res)
	}
	return result, err
}
