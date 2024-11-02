package logging

import (
	"log/slog"
	"os"
)

var logging *slog.Logger

func init() {
	logging = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "msg" {
				a.Key = "message"
			}
			return a
		},
	}))
}

func Info(msg string, args ...interface{}) {
	logging.Info(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	logging.Debug(msg, args...)
}

func Error(msg string, args ...interface{}) {
	logging.Error(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	logging.Warn(msg, args...)
}
