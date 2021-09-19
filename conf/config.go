package conf

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strings"
)

var WORKDIR string

var config Config

func GetConfig() Config {
	return config
}

type Config struct {
	Log LogConfig `ini:"log"`
}

type LogConfig struct {
	Level string `ini:"level"`
	Dir   string `ini:"dir"`
}

func init() {
	initPath()
	initIniConfig()

}

func initPath() {
	WORKDIR, _ = os.Getwd()
}

// initIniConfig 初始化ini配置
func initIniConfig() {
	iniPath := getIniConfigPath()

	confIni, err := ini.Load(iniPath)
	if err != nil {
		log.Fatalln("读取配置文件出现错误")
	}

	config = Config{}
	err = confIni.MapTo(&config)
	if err != nil {
		log.Fatalln("读取配置文件出现错误")
	}
	log.Printf("加载配置文件%s成功\n", iniPath)

}

// getIniConfigPath 获取ini配置文件路径
func getIniConfigPath() string {
	var iniPathBuilder strings.Builder
	iniPathBuilder.WriteString(WORKDIR)
	iniPathBuilder.WriteRune(os.PathSeparator)
	iniPathBuilder.WriteString("git-knowledge.ini")
	iniPath := iniPathBuilder.String()
	return iniPath
}
