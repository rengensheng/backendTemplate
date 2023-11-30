package logger

import (
	"github.com/hyahm/golog"
)

func InitLogger(logPath string) {
	if logPath == "" {
		logPath = "./"
	}
	golog.InitLogger(logPath, 0, true)
}
