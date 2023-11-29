package logger

import (
	"github.com/hyahm/golog"
	"path/filepath"
)

func InitLogger(logPath, logName string) {
	if logPath == "" {
		logPath = "./"
	}
	if logName == "" {
		logName = "logger"
	}
	golog.InitLogger(filepath.Join(logPath, logName), 0, true)
}
