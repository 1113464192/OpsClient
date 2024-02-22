package configs

type Config struct {
	ServerSide ServerSide `json:"server_side"`
	Auth       Auth       `json:"auth"`
	CustomCmd  CustomCmd  `json:"custom_cmd"`
}

type ServerSide struct {
	Ip        string
	AllowCidr string
	Domain    string
	IsSSL     string
	Port      string
}

type Auth struct {
	Key string
}

type CustomCmd struct {
	LocalIpCmd string
	LocalIpApi string
	GameSumCmd string
}

var Conf = new(Config)
