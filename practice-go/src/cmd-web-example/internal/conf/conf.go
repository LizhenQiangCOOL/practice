package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

const (
	EnvPROD  = "prod"
	EnvBETA  = "beta"
	EnvDEV   = "dev"
	EnvLOCAL = "local"
	EnvTEST  = "test"

	WebDir = "html/"
)

var (
	ENV              string
	WorkDir          = "."
	ConfigFile       = ""
	ServerHost       = "0.0.0.0"
	ServerPort       = 80
	SSLHost          = "0.0.0.0"
	SSLPort          = 443
	SSLCert          string
	SSLKey           string
	ErrorLogLevel    = "warn"
	ErrorLogPath     = "logs/error.log"
	AccessLogPath    = "logs/access.log"
	SSLDefaultStatus = 1 //enable ssl by default
	ImportSizeLimit  = 10 * 1024 * 1024
	AllowList        []string
)

type SSL struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Cert string `mapstructure:"cert"`
	Key  string `mapstructure:"key"`
}

type Listen struct {
	Host string
	Port int
}

type ErrorLog struct {
	Level    string
	FilePath string `mapstructure:"file_path"`
}

type AccessLog struct {
	FilePath string `mapstructure:"file_path"`
}

type Log struct {
	ErrorLog  ErrorLog  `mapstructure:"error_log"`
	AccessLog AccessLog `mapstructure:"access_log"`
}

type Conf struct {
	Listen    Listen
	SSL       SSL
	Log       Log
	AllowList []string `mapstructure:"allow_list"`
	MaxCpu    int      `mapstructure:"max_cpu"`
}

type Config struct {
	Conf Conf
}

func InitConf() {
	//go test
	if workDir := os.Getenv("MY_WEB_WORKDIR"); workDir != "" {
		WorkDir = workDir
	}

	setupConfig()
	setupEnv()
}

func setupConfig() {
	// setup config file path
	// 没有设置configfile时，配置文件在WordDIR/conf查找
	if ConfigFile == "" {
		viper.SetConfigName("conf")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(WorkDir + "/conf")
	} else {
		viper.SetConfigFile(ConfigFile)
	}

	// load config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Sprintf("fail to find configuration: %s", ConfigFile))
		} else {
			panic(fmt.Sprintf("fail to read configuration: %s, err: %s", ConfigFile, err.Error()))
		}
	}

	// unmarshal config
	config := Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Sprintf("fail to unmarshal configuration: %s, err: %s", ConfigFile, err.Error()))
	}

	// listen
	if config.Conf.Listen.Port != 0 {
		ServerPort = config.Conf.Listen.Port
	}
	if config.Conf.Listen.Host != "" {
		ServerHost = config.Conf.Listen.Host
	}

	// SSL
	if config.Conf.SSL.Port != 0 {
		SSLPort = config.Conf.SSL.Port
	}
	if config.Conf.SSL.Cert != "" {
		SSLCert = config.Conf.SSL.Cert
	}
	if config.Conf.SSL.Key != "" {
		SSLKey = config.Conf.SSL.Key
	}

	// error log
	if config.Conf.Log.ErrorLog.Level != "" {
		ErrorLogLevel = config.Conf.Log.ErrorLog.Level
	}
	if config.Conf.Log.ErrorLog.FilePath != "" {
		ErrorLogPath = config.Conf.Log.ErrorLog.FilePath
	}

	// access log
	if config.Conf.Log.AccessLog.FilePath != "" {
		AccessLogPath = config.Conf.Log.AccessLog.FilePath
	}

	if !filepath.IsAbs(ErrorLogPath) {
		if strings.HasPrefix(ErrorLogPath, "winfile") {
			return
		}
		ErrorLogPath, err = filepath.Abs(filepath.Join(WorkDir, ErrorLogPath))
		if err != nil {
			panic(err)
		}
		if runtime.GOOS == "windows" {
			ErrorLogPath = `winfile:///` + ErrorLogPath
		}
	}
	if !filepath.IsAbs(AccessLogPath) {
		if strings.HasPrefix(AccessLogPath, "winfile") {
			return
		}
		AccessLogPath, err = filepath.Abs(filepath.Join(WorkDir, AccessLogPath))
		if err != nil {
			panic(err)
		}
		if runtime.GOOS == "windows" {
			AccessLogPath = `winfile:///` + AccessLogPath
		}
	}

	AllowList = config.Conf.AllowList

	// set degree of parallelism
	initParallelism(config.Conf.MaxCpu)
}

func setupEnv() {
	ENV = EnvPROD
	if env := os.Getenv("ENV"); env != "" {
		ENV = env
	}
}

// initialize parallelism settings
func initParallelism(choiceCores int) {
	if choiceCores < 1 {
		return
	}
	maxSupportedCores := runtime.NumCPU()

	if choiceCores > maxSupportedCores {
		choiceCores = maxSupportedCores
	}
	runtime.GOMAXPROCS(choiceCores)
}
