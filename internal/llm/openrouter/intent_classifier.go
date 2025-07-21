package openrouter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/otosei-ai/otosei-ai-backend/internal/intent"
)

type IntentResult struct {
	Intent     string  `json:"intent"`
	Confidence float64 `json:"confidence"`
}

func (c *OpenRouterClient) ClassifyIntent(prompt string) (*IntentResult, error) {
	thresholdStr := c.ConfidenceThreshold
	threshold := 0.5 // default
	if thresholdStr != "" {
		fmt.Sscanf(thresholdStr, "%f", &threshold)
	}

	var exampleList []string
	for name, data := range intent.IntentList {
		block := fmt.Sprintf(
			"Intent: %s\nDescription: %s\nKeywords: %s",
			name,
			data.Description,
			strings.Join(data.Keywords, ", "),
		)
		exampleList = append(exampleList, block)
	}
	exampleListStr := strings.Join(exampleList, "\n\n")

	systemPrompt := fmt.Sprintf(`
Your task is to classify the intent of the following user input. 
You must respond **strictly** with a JSON object in the format:
{"intent": "<detected_intent>", "confidence": <confidence_value>}
Do not include any extra text, explanations, or formatting.

You are an AI that classifies user intents and extracts transaction details.
Available intents: %s

Here are the intents with examples:

%s

Extract the intent and a confidence score (0-1).
If the confidence is below %s, classify it as "default".

**IMPORTANT !**
**ONLY MAKE YOUR RESPOND AS I INSTRUCTED YOU AS JSON FORMAT WE NEED. DO NOT WRITE ANYTHING ELSE !**
`, strings.Join(getIntentKeys(), ", "), exampleListStr, thresholdStr)

	payload := ChatRequest{
		Model: c.Model,
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: prompt},
		},
		MaxTokens:   50,
		Temperature: 0.2,
	}

	data, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", c.BaseURL, bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("non-200 response from OpenRouter: %d - %s", resp.StatusCode, string(body))
	}

	var res ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	if len(res.Choices) == 0 {
		return nil, fmt.Errorf("empty choices from model")
	}

	content := res.Choices[0].Message.Content

	var parsed IntentResult
	if err := json.Unmarshal([]byte(strings.TrimSpace(content)), &parsed); err != nil {
		return &IntentResult{Intent: "default", Confidence: 0}, nil
	}

	if parsed.Intent == "" || parsed.Confidence < threshold {
		return &IntentResult{Intent: "default", Confidence: parsed.Confidence}, nil
	}

	return &parsed, nil
}

func getIntentKeys() []string {
	keys := make([]string, 0, len(intent.IntentList))
	for k := range intent.IntentList {
		keys = append(keys, k)
	}
	return keys
}
