package internal

type Config struct {
	Url           string `mapstructure:"url"`
	ClashFilePath string `mapstructure:"clash_file_path"`
}

type ClashConfig struct {
	Port                 int                  `yaml:"port"`
	SocksPort            int                  `yaml:"socks-port"`
	RedirPort            int                  `yaml:"redir-port"`
	AllowLan             bool                 `yaml:"allow-lan"`
	Mode                 string               `yaml:"mode"`
	LogLevel             string               `yaml:"log-level"`
	ExternalController   string               `yaml:"external-controller"`
	CfwConnBreakStrategy CfwConnBreakStrategy `yaml:"cfw-conn-break-strategy"`
	CfwLatencyTimeout    string               `yaml:"cfw-latency-timeout"`
	Proxies              []interface{}        `yaml:"proxies"`
}
type CfwConnBreakStrategy struct {
	Proxy   string `yaml:"proxy"`
	Profile bool   `yaml:"profile"`
	mode    bool   `yaml:"mode"`
}
