package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type CustomLogger struct {
	entry *logrus.Entry
}

func (cl CustomLogger) Info(args ...interface{}) {
	fmt.Println(args...)
}

func (cl CustomLogger) Warn(args ...interface{}) {
	cl.entry.Warn(args)
}

func (cl CustomLogger) Debug(args ...interface{}) {
	cl.entry.Debug(args)
}

func (cl CustomLogger) Error(args ...interface{}) {
	cl.entry.Error(args)
}

func (cl CustomLogger) Panic(args ...interface{}) {
	cl.entry.Error(args)
	os.Exit(1)
}

func New(labels map[string]interface{}) *CustomLogger {
	var log = logrus.New()
	if labels == nil {
		return &CustomLogger{entry: log.WithFields(map[string]interface{}{})}
	}
	return &CustomLogger{entry: log.WithFields(labels)}
}
