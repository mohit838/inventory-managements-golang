package logging

import (
	"log/slog"
	"os"

	"github.com/natefinch/lumberjack"
)

func Init() {
	_ = os.MkdirAll("logs", 0755)

	logWriter := &lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     28,
		Compress:   true,
	}

	logger := slog.New(slog.NewJSONHandler(logWriter, nil))
	slog.SetDefault(logger)
}
