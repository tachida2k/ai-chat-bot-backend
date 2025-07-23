package openrouter

import (
	"strings"
)

type OpenRouterClient struct {
	APIKey              string
	BaseURL             string
	Model               string
	Fallbacks           []string
	ConfidenceThreshold string
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Stream      bool      `json:"stream,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
}

type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func NewClient(apiKey, baseURL, model, fallbackRaw, confidenceThreshold string) *OpenRouterClient {
	var fallbacks []string
	if fallbackRaw != "" {
		fallbacks = strings.Split(fallbackRaw, ",")
	}

	return &OpenRouterClient{
		APIKey:              apiKey,
		BaseURL:             baseURL,
		Model:               model,
		Fallbacks:           fallbacks,
		ConfidenceThreshold: confidenceThreshold,
	}
}
