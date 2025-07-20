package openrouter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *OpenRouterClient) ChatStream(ctx *gin.Context, messages []Message) error {
	models := append([]string{c.Model}, c.Fallbacks...)
	for _, model := range models {
		err := c.sendChatStreamRequest(ctx, messages, model)
		if err == nil {
			return nil
		}
		fmt.Printf("Error with model %s: %v, trying next...\n", model, err)
	}

	return fmt.Errorf("all models failed")
}

func (c *OpenRouterClient) sendChatStreamRequest(ctx *gin.Context, messages []Message, model string) error {
	payload := ChatRequest{
		Model:    model,
		Messages: messages,
		Stream:   true,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.BaseURL, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("non-200 response from OpenRouter: %d - %s", resp.StatusCode, string(body))
	}

	reader := io.Reader(resp.Body)
	bufReader := NewEventReader(reader)

	ctx.Stream(func(w io.Writer) bool {
		for {
			event, err := bufReader.ReadEvent()
			if err != nil {
				if err == io.EOF {
					return false
				}
				fmt.Println("Read error:", err)
				return false
			}

			if event == "[DONE]" {
				ctx.SSEvent("done", true)
				return false
			}

			var parsed ChatStreamResponse
			if err := json.Unmarshal([]byte(event), &parsed); err != nil {
				fmt.Println("Unmarshal error:", err)
				continue
			}

			if len(parsed.Choices) > 0 {
				content := parsed.Choices[0].Delta.Content
				if content != "" {
					ctx.SSEvent("message", content)
				}
			}
			return true
		}
	})
	return nil
}
