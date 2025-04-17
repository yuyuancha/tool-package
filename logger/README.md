## Log 日誌記錄

### 使用方式

```shell
go get -u github.com/yuyuancha/tool-package/logger
```

初始化日誌記錄器。

```
err := logger.InitialLogger(&logger.Option{
    IsReportCaller: true,
	SetLevel:       logrus.DebugLevel,
	FileName:       "test",
	MaxSize:        1,
	MaxBackups:     3,
	MaxAge:         1,
	Compress:       true,
})
if err != nil {
    panic(err)
}
```
