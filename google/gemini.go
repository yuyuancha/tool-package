package google

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"github.com/yuyuancha/tool-package/config"
	"github.com/yuyuancha/tool-package/constants"
	"google.golang.org/api/option"
)

// Gemini Google Gemini 服務
type Gemini struct {
	Token string
}

// NewGemini 初始化 Google Gemini 服務
func NewGemini() *Gemini {
	config.Initial()
	config.ReadGoogleServiceConfig()

	return &Gemini{
		Token: config.GoogleService.GeminiApiKey,
	}
}

// RequestTextByText 透過文字請求文字回應
func (gemini *Gemini) RequestTextByText(question string) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(gemini.Token))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel(constants.GeminiMode)
	response, err := model.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		return "", err
	}

	answer := gemini.parseAIResponse(response)

	return answer, nil
}

// 解析 AI 回應
func (gemini *Gemini) parseAIResponse(response *genai.GenerateContentResponse) (results string) {
	for _, candidate := range response.Candidates {
		for _, part := range candidate.Content.Parts {
			results = results + fmt.Sprintf("%v", part)
		}
	}
	return
}
