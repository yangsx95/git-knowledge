package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
)

// InitLogger 初始化日志
func InitLogger(level, dir string) {
	// 读取配置并转换为zap的日志级别
	zapLevel := convZapLevel(level)
	// 读取配置并将日志文件夹位置转换为绝对路径
	logDir, _ := filepath.Abs(dir)
	// 如果文件夹不存在则创建
	createDirIfNotExist(logDir)
	// 生成日志输出路径
	outputPath, errOutputPath := logDes(logDir)

	// 创建日志配置，并构建日志记录器对象
	config := zap.Config{
		Level:            zapLevel,
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      outputPath,
		ErrorOutputPaths: errOutputPath,
	}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	logger, err := config.Build(
		zap.AddCallerSkip(1),
	)

	if err != nil {
		log.Fatalln("日志初始化失败", err)
	}

	// 替换全局日志
	zap.ReplaceGlobals(logger)
}

func logDes(logDir string) ([]string, []string) {
	outputPath, errOutputPath := make([]string, 0), make([]string, 0)
	outputPath = append(outputPath, "stdout")
	errOutputPath = append(errOutputPath, "stderr")
	absInfoLogPath, _ := filepath.Abs(filepath.Join(logDir, "info.log"))
	absErrLogPath, _ := filepath.Abs(filepath.Join(logDir, "err.log"))

	outputPath = append(outputPath, absInfoLogPath)
	errOutputPath = append(errOutputPath, absErrLogPath)
	return outputPath, errOutputPath
}

func createDirIfNotExist(logDir string) {
	_, err := os.Stat(logDir)
	if err != nil && os.IsNotExist(err) {
		err := os.MkdirAll(logDir, 0766)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if err != nil && !os.IsNotExist(err) {
		log.Fatalln(err)
	}
}

func convZapLevel(logLevel string) zap.AtomicLevel {
	level := zap.NewAtomicLevel()
	err := level.UnmarshalText([]byte(logLevel))
	if err != nil {
		log.Fatalf("日志级别配置有误 %s", logLevel)
	}
	return level
}

func GetLogger() *zap.Logger {
	return zap.L()
}

func Debug(template string, args ...interface{}) {
	zap.S().Debug(template, args)
}

func Info(template string, args ...interface{}) {
	zap.S().Infof(template, args)
}

func Warn(template string, args ...interface{}) {
	zap.S().Warnf(template, args)
}

func Error(template string, args ...interface{}) {
	zap.S().Errorf(template, args)
}

func Fatal(template string, args ...interface{}) {
	zap.S().Fatalf(template, args)
}
