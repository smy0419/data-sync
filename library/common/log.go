package common

import (
	"bufio"
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/xormplus/core"

	"os"
	"path"
	"time"
)

type WrapLog struct {
	logrus *logrus.Logger
}

var Logger = WrapLog{}

func init() {
	Logger.logrus = logrus.New()
	if Cfg.Env == ENV_DEVELOP_NET {
		Logger.logrus.SetOutput(os.Stdout)
		configLocalFilesystemLogger(Logger.logrus)
		Logger.logrus.SetLevel(logrus.DebugLevel)
	} else {
		configLocalFilesystemLogger(Logger.logrus)
		Logger.logrus.SetLevel(logrus.InfoLevel)
		// When the file is entered, close the console output
		setNull(Logger.logrus)
	}
}

func configLocalFilesystemLogger(l *logrus.Logger) {
	logPath := path.Join(Cfg.LogDir, Cfg.ProjectName)
	os.MkdirAll(logPath, os.ModePerm)
	baseLogPath := path.Join(logPath, Cfg.ProjectName+".log")

	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		// Generate a soft link to the latest log file
		rotatelogs.WithLinkName(baseLogPath),
		// Maximum file save time
		rotatelogs.WithMaxAge(time.Hour*24*7),
		// Log cutting interval
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{DisableColors: true})

	l.AddHook(lfHook)
}

func setNull(l *logrus.Logger) {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	writer := bufio.NewWriter(src)
	l.SetOutput(writer)
}

func (wrapLog WrapLog) Debug(v ...interface{}) {
	wrapLog.logrus.Debug(v...)
}
func (wrapLog WrapLog) Debugf(format string, v ...interface{}) {
	wrapLog.logrus.Debugf(format, v...)
}
func (wrapLog WrapLog) Error(v ...interface{}) {
	wrapLog.logrus.Error(v...)
}

func (wrapLog WrapLog) ErrorPanic(info string, err error) {
	wrapLog.Error(info, err)
	panic(err)
}

func (wrapLog WrapLog) Errorf(format string, v ...interface{}) {
	wrapLog.logrus.Errorf(format, v...)
}

func (wrapLog WrapLog) ErrorfPanic(err error, format string, v ...interface{}) {
	wrapLog.Errorf(format, v...)
	wrapLog.Error(err)
	panic(err)
}

func (wrapLog WrapLog) Info(v ...interface{}) {
	wrapLog.logrus.Info(v...)
}

func (wrapLog WrapLog) Infof(format string, v ...interface{}) {
	wrapLog.logrus.Infof(format, v...)
}

func (wrapLog WrapLog) Printf(format string, v ...interface{}) {
	wrapLog.Errorf(format, v...)
}

func (wrapLog WrapLog) Warn(v ...interface{}) {
	wrapLog.logrus.Warn(v...)
}
func (wrapLog WrapLog) Warnf(format string, v ...interface{}) {
	wrapLog.logrus.Warnf(format, v...)
}
func (wrapLog WrapLog) Level() core.LogLevel {
	return core.LOG_INFO
}
func (wrapLog WrapLog) SetLevel(l core.LogLevel) {
}

func (wrapLog WrapLog) ShowSQL(show ...bool) {
	wrapLog.Infof("show sql: %t", show)
}

func (wrapLog WrapLog) IsShowSQL() bool {
	return Cfg.ShowSql
}
