package api

type ShellRes struct {
	ServerDir string `json:"server_dir,omitempty"`
	Status    int    `json:"status"`
	Response  string `json:"response"`
}
