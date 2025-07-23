package openrouter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type IntentData struct {
	Description string
	Examples    []string
	Keywords    []string
}
type ClassifyIntentResult struct {
	Intent     string  `json:"intent"`
	Confidence float64 `json:"confidence"`
}

var IntentList = map[string]IntentData{
	"swap_token": {
		Description: "Executes a token swap from Coinhall Routes.",
		Examples: []string{
			"I want to trade SEI for USDT.",
			"Swap 50 SEI to BNB.",
			"Convert my SEI into ETH.",
			"Trade 10 USDT into SEI.",
		},
		Keywords: []string{
			"swap", "exchange", "convert", "trade",
			"swap tokens", "exchange tokens", "convert tokens", "trade tokens",
			"swap SEI", "swap to USDT", "trade for", "convert my",
			"exchange my", "where can I swap", "how to swap",
		},
	},
	"stake": {
		Description: "Stake SEI tokens to validator.",
		Examples: []string{
			"I want to stake my SEI tokens.",
			"How do I delegate SEI?",
			"Stake 100 SEI with a validator.",
		},
		Keywords: []string{
			"stake", "staking", "delegate", "validator", "earn rewards",
			"staking rewards", "staking pool",
		},
	},
	"unstake": {
		Description: "Unstake SEI tokens.",
		Examples: []string{
			"I want to unstake my SEI tokens.",
			"How do I undelegate my SEI?",
			"Unstake 50 SEI from validator.",
		},
		Keywords: []string{
			"unstake", "unstaking", "undelegate", "withdraw stake",
			"remove stake", "stop staking", "withdraw staked", "exit staking",
		},
	},
	"send_token": {
		Description: "Send tokens to an EVM address.",
		Examples: []string{
			"Send 5 SEI to my friend.",
			"Transfer SEI to this wallet.",
			"Move 10 USDT to another account.",
		},
		Keywords: []string{
			"send", "transfer", "move", "send SEI", "transfer USDT",
			"send funds", "move tokens", "send crypto", "send to address",
		},
	},
	"check_balance": {
		Description: "Check my wallet balance.",
		Examples: []string{
			"Check my wallet balance.",
			"What is my current SEI balance?",
			"Show me my token holdings.",
		},
		Keywords: []string{
			"my balance", "wallet balance", "portfolio", "how much", "check my funds",
			"my tokens", "holdings", "how much sei", "current balance",
		},
	},
	"tx_search": {
		Description: "Search for a transaction.",
		Examples: []string{
			"Find this transaction hash on Sei.",
			"Check this transaction ID: 0x1234abcd.",
		},
		Keywords: []string{
			"tx", "transaction", "txid", "tx hash", "check tx",
		},
	},
	"analyze_wallet": {
		Description: "Analyze wallet portfolio.",
		Examples: []string{
			"Analyze this SEI wallet.",
			"What tokens does this address hold?",
			"Give me the portfolio breakdown.",
		},
		Keywords: []string{
			"analyze", "wallet overview", "token distribution", "asset summary",
			"wallet breakdown", "portfolio analysis",
		},
	},
	"forbidden_topics": {
		Description: "Topics unrelated to Sei or out of scope.",
		Examples: []string{
			"Write a Python script.",
			"What are the latest updates in Bitcoin?",
			"Can you explain AI and machine learning?",
		},
		Keywords: []string{
			"python", "AI", "machine learning", "bitcoin", "stock", "solana", "ethereum",
			"script", "trading bots", "smart contract", "chatgpt", "llama", "openai",
		},
	},
	"default": {
		Description: "Greetings, general questions about Sei.",
		Examples: []string{
			"Who is Sonia?",
			"Hey there!",
			"Tell me about Sei.",
			"I am new to Sei—where should I start?",
		},
		Keywords: []string{
			"hi", "hello", "thanks", "good morning", "good evening", "sei", "what is sei",
			"how does sei work", "tell me about sei", "getting started",
		},
	},
}

func (c *OpenRouterClient) ClassifyIntent(prompt string) (*ClassifyIntentResult, error) {
	thresholdStr := c.ConfidenceThreshold
	threshold := 0.5 // default
	if thresholdStr != "" {
		fmt.Sscanf(thresholdStr, "%f", &threshold)
	}

	var exampleList []string
	for name, data := range IntentList {
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
		return nil, fmt.Errorf("non-200 response from LLM: %d - %s", resp.StatusCode, string(body))
	}

	var res ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	if len(res.Choices) == 0 {
		return nil, fmt.Errorf("empty choices from model")
	}

	content := res.Choices[0].Message.Content

	var parsed ClassifyIntentResult
	if err := json.Unmarshal([]byte(strings.TrimSpace(content)), &parsed); err != nil {
		return &ClassifyIntentResult{Intent: "default", Confidence: 0}, nil
	}

	if parsed.Intent == "" || parsed.Confidence < threshold {
		return &ClassifyIntentResult{Intent: "default", Confidence: parsed.Confidence}, nil
	}

	return &parsed, nil
}

func getIntentKeys() []string {
	keys := make([]string, 0, len(IntentList))
	for k := range IntentList {
		keys = append(keys, k)
	}
	return keys
}
