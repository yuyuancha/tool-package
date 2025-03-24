package logger

import (
	"errors"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

const (
	logDirPath        = "./log"
	dateTimeLayout    = "2006-01-02 15:04:05"
	defaultFileName   = "debug.log"
	defaultMaxSize    = 1
	defaultMaxBackups = 3
	defaultMaxAge     = 1
)

var baseLog = logrus.New()

// InitialLogger 初始化日誌設定
func InitialLogger(option *Option) error {
	err := createLogDir()
	if err != nil {
		return err
	}

	option = checkOption(option)

	baseLog.SetFormatter(&baseFormatter{})
	baseLog.SetReportCaller(option.IsReportCaller)
	baseLog.SetLevel(option.getSetLevel())

	writers := io.MultiWriter(
		os.Stdout, &lumberjack.Logger{
			Filename:   logDirPath + "/" + option.getFileName(),
			MaxSize:    option.getMaxSize(),
			MaxBackups: option.getMaxBackups(),
			MaxAge:     option.getMaxAge(),
			Compress:   option.Compress,
		})

	baseLog.SetOutput(writers)

	return nil
}

// LogInfo info 日誌
func LogInfo(args ...interface{}) {
	baseLog.Info(args)
}

// LogWarn warn 日誌
func LogWarn(args ...interface{}) {
	baseLog.Warn(args)
}

// LogError error 日誌
func LogError(args ...interface{}) {
	baseLog.Error(args)
}

// LogFatal fatal 日誌
func LogFatal(args ...interface{}) {
	baseLog.Fatal(args)
}

// LogPanic panic 日誌
func LogPanic(args ...interface{}) {
	baseLog.Panic(args)
}

// 檢查選項
func checkOption(option *Option) *Option {
	if option != nil {
		return option
	}

	return &Option{
		IsReportCaller: true,
		SetLevel:       logrus.DebugLevel,
		FileName:       defaultFileName,
		MaxSize:        defaultMaxSize,
		MaxBackups:     defaultMaxBackups,
		MaxAge:         defaultMaxAge,
		Compress:       false,
	}
}

// 建立 log 資料夾
func createLogDir() error {
	_, err := os.Stat(logDirPath)
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		err = os.Mkdir(logDirPath, os.ModePerm)
		if err != nil {
			return errors.New("Failed to create log directory, error: " + err.Error())
		}
	}

	return errors.New("Failed to check log directory, error: " + err.Error())
}
