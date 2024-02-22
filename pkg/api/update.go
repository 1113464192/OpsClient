package api

type UpdateServerReq struct {
	Tid        uint     `form:"tid" json:"tid" `
	ServerDir  []string `form:"server_dir" json:"server_dir"`
	PackageDir string   `form:"package_dir" json:"package_dir"`
}

type UpdateExecReq struct {
	Tid       uint     `form:"tid" json:"tid" `
	ServerDir []string `form:"server_dir" json:"server_dir"`
	Type      int      `form:"type" json:"type"` // 闪断:1 热更:2
}
