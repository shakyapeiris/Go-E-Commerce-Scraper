package logging

import (
	"log"
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

func includeCID(args *[]interface{}) {
	*args = append(*args, slog.String("cid", GetCID()))
}

func Info(msg string, args ...interface{}) {
	includeCID(&args)
	logging.Info(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	includeCID(&args)
	logging.Debug(msg, args...)
}

func Error(msg string, args ...interface{}) {
	includeCID(&args)
	logging.Error(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	logging.Warn(msg, args...)
}

func Fatalln(v ...any) {
	log.Fatalln(v)
}
