package logger

import (
	"git-knowledge/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func init() {
	logConfig := conf.GetConfig().Log
	// 日志级别
	level := zap.NewAtomicLevel()
	err := level.UnmarshalText([]byte(logConfig.Level))
	if err != nil {
		log.Fatalf("日志级别配置有误 %s", logConfig.Level)
	}

	// 日志输出路径
	outputPath, errOutputPath := make([]string, 2), make([]string, 2)
	outputPath = append(outputPath, "stdout")
	errOutputPath = append(errOutputPath, "stdout")
	if logConfig.Dir != "" {
		outputPath = append(outputPath, logConfig.Dir)
		errOutputPath = append(outputPath, logConfig.Dir)
	}

	config := zap.Config{
		Level:            level,
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      outputPath,
		ErrorOutputPaths: errOutputPath,
	}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	logger, _ := config.Build(
		zap.AddCallerSkip(1),
	)
	zap.ReplaceGlobals(logger)
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
