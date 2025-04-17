package config

import (
	"errors"
)

// GoogleService google service 結構
var GoogleService googleServiceStruct

// google service 結構
type googleServiceStruct struct {
	GeminiApiKey string
}

// ReadGoogleServiceConfig 讀取 google 服務
func ReadGoogleServiceConfig() error {
	if v == nil {
		return errors.New("config 讀取失敗")
	}

	// 配置 google service 結構
	GoogleService = googleServiceStruct{
		GeminiApiKey: v.GetString("GOOGLE_GEMINI_API_KEY"),
	}

	return nil
}
