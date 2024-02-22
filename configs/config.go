package configs

type Config struct {
	ServerIp  ServerIp  `json:"server_ip"`
	Auth      Auth      `json:"auth"`
	CustomCmd CustomCmd `json:"custom_cmd"`
}

type ServerIp struct {
	Value     string
	AllowCidr string
}

type Auth struct {
	Key string
}

type CustomCmd struct {
	GetLocalIpCmd string
	GetLocalIpApi string
}

var Conf = new(Config)
