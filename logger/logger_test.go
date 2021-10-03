package logger

import "testing"

func TestDebug(t *testing.T) {
	InitLogger("debug", "./logs")
	Debug("Hello %s", "World")
}

func TestInfo(t *testing.T) {
	InitLogger("debug", "./logs")
	Info("Hello %s", "World")
}

func TestError(t *testing.T) {
	InitLogger("debug", "./logs")
	Error("Hello %s", "World")
}
