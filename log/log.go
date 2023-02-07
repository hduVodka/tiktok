package log

import (
	"context"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func Init() {
	writer, err := rotatelogs.New("logs/%Y%m%d.log")
	if err != nil {
		Logger.Fatalln(err)
	}
	Logger.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writer,
			logrus.ErrorLevel: writer,
		}, &logrus.JSONFormatter{},
	))
}

func Info(args ...interface{}) {
	Logger.Info(args)
}

func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args)
}

func Infoln(args ...interface{}) {
	Logger.Infoln(args)
}

func Error(args ...interface{}) {
	Logger.Error(args)
}

func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args)
}

func Errorln(args ...interface{}) {
	Logger.Errorln(args)
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args)
}

func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args)
}

func Fatalln(args ...interface{}) {
	Logger.Fatalln(args)
}

func WithFiled(key string, value interface{}) *logrus.Entry {
	return Logger.WithField(key, value)
}

func WithContext(ctx context.Context) *logrus.Entry {
	return Logger.WithContext(ctx)
}
