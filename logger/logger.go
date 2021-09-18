package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
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

func Error(template string, args ...interface{}) {
	zap.S().Errorf(template, args)
}
