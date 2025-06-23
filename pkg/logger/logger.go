package logger

import (
	"TrainTracking/internal/config"
	"os"

	"github.com/sirupsen/logrus"
)

// Log is the global logger instance
var Log = logrus.New()

func InitLogger() {
	if config.GetApp().Env == "production" {
		Log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Log.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
	}

	Log.SetOutput(os.Stdout)
}

func Debug(msg string, args ...interface{}) {
	Log.Logf(logrus.DebugLevel, msg, args...)
}

func Info(msg string, args ...interface{}) {
	Log.Logf(logrus.InfoLevel, msg, args...)
}

func Warn(msg string, args ...interface{}) {
	Log.Logf(logrus.WarnLevel, msg, args...)
}

func Error(msg string, args ...interface{}) {
	Log.Logf(logrus.ErrorLevel, msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	Log.Logf(logrus.FatalLevel, msg, args...)
}

func Panic(msg string, args ...interface{}) {
	Log.Logf(logrus.PanicLevel, msg, args...)
}
