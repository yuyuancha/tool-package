package logger

import (
	"github.com/sirupsen/logrus"
	"strings"
)

// Option 選項
type Option struct {
	IsReportCaller bool         // 是否紀錄檔案位置及行數
	SetLevel       logrus.Level // 設定紀錄等級
	FileName       string       // 檔案名稱
	MaxSize        int          // 最大檔案大小
	MaxBackups     int          // 最大檔案數
	MaxAge         int          // 最大檔案保存天數
	Compress       bool         // 是否壓縮
}

// 取得紀錄等級
func (option *Option) getSetLevel() logrus.Level {
	if option.SetLevel == 0 {
		return logrus.DebugLevel
	}
	return option.SetLevel
}

// 取得檔案名稱
func (option *Option) getFileName() string {
	if option.FileName == "" {
		return defaultFileName
	}
	if !strings.HasSuffix(option.FileName, ".log") {
		return option.FileName + ".log"
	}
	return option.FileName
}

// 取得最大檔案大小
func (option *Option) getMaxSize() int {
	if option.MaxSize == 0 {
		return defaultMaxSize
	}
	return option.MaxSize
}

// 取得最大檔案數
func (option *Option) getMaxBackups() int {
	if option.MaxBackups == 0 {
		return defaultMaxBackups
	}
	return option.MaxBackups
}

// 取得最大檔案保存天數
func (option *Option) getMaxAge() int {
	if option.MaxAge == 0 {
		return defaultMaxAge
	}
	return option.MaxAge
}
