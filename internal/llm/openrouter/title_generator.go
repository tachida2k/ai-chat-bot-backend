package openrouter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *OpenRouterClient) GenerateChatTitle(userMessage string) (string, error) {
	systemPrompt := fmt.Sprintf(`
You are an AI assistant that generates concise and contextually relevant sentences based on the user's message.
Your response should be short, clear, and aligned with the given input. Your goal is only summarizing the message.
Only respond the summarize as a title shortly.

User's message: (%s)`,
		userMessage)

	payload := ChatRequest{
		Model: c.Model,
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
		},
		MaxTokens:   50,
		Temperature: 0.4,
	}

	data, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", c.BaseURL, bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("non-200 response from LLM: %d - %s", resp.StatusCode, string(body))
	}

	var res ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	if len(res.Choices) == 0 {
		return "", fmt.Errorf("empty choices from model")
	}

	content := res.Choices[0].Message.Content

	if content == "" {
		return "Untitled", nil
	}

	return content, nil
}
