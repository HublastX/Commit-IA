package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type OpenAIClient struct {
	APIKey string
	Model  string
}

type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIRequest struct {
	Model       string          `json:"model"`
	Messages    []OpenAIMessage `json:"messages"`
	Temperature float64         `json:"temperature"`
	MaxTokens   int             `json:"max_tokens"`
}

type OpenAIChoice struct {
	Message OpenAIMessage `json:"message"`
}

type OpenAIResponse struct {
	Choices []OpenAIChoice `json:"choices"`
	Error   *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error,omitempty"`
}

func NewOpenAIClient(model string) *OpenAIClient {
	return &OpenAIClient{
		APIKey: os.Getenv("OPENAI_API_KEY"),
		Model:  model,
	}
}

func (c *OpenAIClient) SetAPIKey(apiKey string) {
	c.APIKey = apiKey
}

func (c *OpenAIClient) GenerateCommitMessage(prompt string) (string, error) {
	if c.APIKey == "" {
		return "", fmt.Errorf("OpenAI API key not set")
	}

	requestBody := OpenAIRequest{
		Model: c.Model,
		Messages: []OpenAIMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Temperature: 0.3,
		MaxTokens:   150,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	var response OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	if response.Error != nil {
		return "", fmt.Errorf("OpenAI API error: %s", response.Error.Message)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no response choices returned")
	}

	return response.Choices[0].Message.Content, nil
}

func (c *OpenAIClient) GetProvider() string {
	return "openai"
}

func (c *OpenAIClient) GetModel() string {
	return c.Model
}
