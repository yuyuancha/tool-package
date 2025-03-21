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

func (option *Option) getFileName() string {
	if option.FileName == "" {
		return defaultFileName
	}
	if !strings.HasSuffix(option.FileName, ".log") {
		return option.FileName + ".log"
	}
	return option.FileName
}

func (option *Option) getMaxSize() int {
	if option.MaxSize == 0 {
		return defaultMaxSize
	}
	return option.MaxSize
}

func (option *Option) getMaxBackups() int {
	if option.MaxBackups == 0 {
		return defaultMaxBackups
	}
	return option.MaxBackups
}

func (option *Option) getMaxAge() int {
	if option.MaxAge == 0 {
		return defaultMaxAge
	}
	return option.MaxAge
}
