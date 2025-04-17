package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
	"github.com/yuyuancha/tool-package/constants"
)

var (
	v          *viper.Viper
	root       = constants.ConfigDefaultEnvRoot
	name       = constants.ConfigDefaultEnvName
	configType = constants.ConfigDefaultEnvType
)

// SetCustomizedConfigPath 設定客製化 config 路徑
func SetCustomizedConfigPath(r, n, ct string) {
	root, name, configType = r, n, ct
}

// Initial 初始化 config
func Initial() {
	var once sync.Once
	once.Do(func() {
		readConfig()
	})
}

// 讀取 config
func readConfig() {
	v = viper.New()
	v.SetConfigName(name)
	v.SetConfigType(configType)
	v.AddConfigPath(root)

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln("讀取 config 發生錯誤: ", err.Error())
	}
}
