package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type OpenRouterClient struct {
	APIKey    string
	BaseURL   string
	Model     string
	Fallbacks []string
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func NewOpenRouterClient(apiKey, baseURL, model, fallbackRaw string) *OpenRouterClient {
	var fallbacks []string
	if fallbackRaw != "" {
		fallbacks = strings.Split(fallbackRaw, ",")
	}

	return &OpenRouterClient{
		APIKey:    apiKey,
		BaseURL:   baseURL,
		Model:     model,
		Fallbacks: fallbacks,
	}
}

func (c *OpenRouterClient) Chat(messages []Message) (string, error) {
	models := append([]string{c.Model}, c.Fallbacks...)
	for _, model := range models {
		resp, err := c.sendChatRequest(messages, model)
		if err == nil {
			return resp, nil
		}
		fmt.Printf("Error with model %s: %v, trying next...\n", model, err)
	}

	return "", fmt.Errorf("all models failed")
}

func (c *OpenRouterClient) sendChatRequest(messages []Message, model string) (string, error) {
	payload := ChatRequest{
		Model:    model,
		Messages: messages,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.BaseURL, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("non-200 response from OpenRouter: %d - %s", resp.StatusCode, string(body))
	}

	var res ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	if len(res.Choices) == 0 {
		return "", fmt.Errorf("empty choices from model")
	}

	return res.Choices[0].Message.Content, nil
}
