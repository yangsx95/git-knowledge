package logger

import "testing"

func TestDebug(t *testing.T) {
	Debug("Hello %s", "World")
}

func TestInfo(t *testing.T) {
	Info("Hello %s", "World")
}

func TestError(t *testing.T) {
	Error("Hello %s", "World")
}
