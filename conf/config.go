package conf

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// WORKDIR 当前程序的工作路径
var WORKDIR string

// config 从ini中读取的配置结构体
var config Config

// GetConfig 获取配置
func GetConfig() *Config {
	return &config
}

// Config 配置信息结构体
type Config struct {
	Log    LogConfig    `ini:"log"`
	Github GithubConfig `ini:"github"`
	Mongo  MongoConfig  `ini: "mongo"`
}

// LogConfig 日志配置section
type LogConfig struct {
	Level string `ini:"level"`
	Dir   string `ini:"dir"`
}

// GithubConfig 配置信息
type GithubConfig struct {
	ClientId string `ini:"client_id"`
}

type MongoConfig struct {
	Url string `ini:"url"`
}

// InitConfig 指定配置文件初始化配置
func InitConfig(path string) {
	initPath()
	initIniConfig(path)
}

func initPath() {
	WORKDIR, _ = os.Getwd()
}

// initIniConfig 初始化ini配置
func initIniConfig(path string) {
	// 将路径处理为绝对路径
	absP, _ := filepath.Abs(path)
	// 加载配置文件
	confIni, err := ini.Load(absP)
	if err != nil {
		log.Fatalf("读取配置文件出现错误, 路径: %s, 错误信息： %s\n", absP, err)
	}
	// 映射配置文件到配置结构体
	config = Config{ // 这里提供默认值
		Log: LogConfig{
			Level: "DEBUG",
			Dir:   "./log",
		},
	}
	err = confIni.MapTo(&config)
	if err != nil {
		log.Fatalln("读取配置文件出现错误")
	}
	log.Printf("加载配置文件%s成功\n", absP)
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
