package api

type CmdResultRes struct {
	Status   int    `json:"status"`
	Response string `json:"response"`
}

type GetServerInfoRes struct {
	Path CmdResultRes `json:"path"`
	Flag CmdResultRes `json:"flag"`
	Name CmdResultRes `json:"name"`
}

type GetServerInfoCmdTem struct {
	FlagCmdTem string `json:"flag_cmd_tem"`
	NameCmdTem string `json:"name_cmd_tem"`
}
