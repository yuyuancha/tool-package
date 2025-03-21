package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
)

// 基本格式
type baseFormatter struct{}

// Format 實作格式化
func (formatter *baseFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var (
		newLog     string
		timestamp  = entry.Time.Format(dateTimeLayout)
		byteBuffer = &bytes.Buffer{}
	)
	if entry.Buffer != nil {
		byteBuffer = entry.Buffer
	}

	if entry.HasCaller() {
		fileName := path.Base(entry.Caller.File)
		newLog = fmt.Sprintf("%s [%s] [%s:%d] %s\n",
			timestamp, entry.Level.String(), fileName, entry.Caller.Line, entry.Message)
	} else {
		newLog = fmt.Sprintf("%s [%s] %s\n", timestamp, entry.Level.String(), entry.Message)
	}

	byteBuffer.WriteString(newLog)
	return byteBuffer.Bytes(), nil
}
