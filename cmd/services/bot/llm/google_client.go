package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GoogleClient struct {
	APIKey string
	Model  string
}

type GoogleContent struct {
	Parts []GooglePart `json:"parts"`
}

type GooglePart struct {
	Text string `json:"text"`
}

type GoogleRequest struct {
	Contents         []GoogleContent `json:"contents"`
	GenerationConfig struct {
		Temperature     float64 `json:"temperature"`
		MaxOutputTokens int     `json:"maxOutputTokens"`
	} `json:"generationConfig"`
}

type GoogleCandidate struct {
	Content GoogleContent `json:"content"`
}

type GoogleResponse struct {
	Candidates []GoogleCandidate `json:"candidates"`
	Error      *struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"error,omitempty"`
}

func NewGoogleClient(model string) *GoogleClient {
	return &GoogleClient{
		APIKey: os.Getenv("GOOGLE_API_KEY"),
		Model:  model,
	}
}

func (c *GoogleClient) SetAPIKey(apiKey string) {
	c.APIKey = apiKey
}

func (c *GoogleClient) GenerateCommitMessage(prompt string) (string, error) {
	if c.APIKey == "" {
		return "", fmt.Errorf("Google API key not set")
	}

	requestBody := GoogleRequest{
		Contents: []GoogleContent{
			{
				Parts: []GooglePart{
					{Text: prompt},
				},
			},
		},
	}
	requestBody.GenerationConfig.Temperature = 0.3
	requestBody.GenerationConfig.MaxOutputTokens = 150

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", c.Model, c.APIKey)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	var response GoogleResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	if response.Error != nil {
		return "", fmt.Errorf("Google API error: %s", response.Error.Message)
	}

	if len(response.Candidates) == 0 || len(response.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response content returned")
	}

	return response.Candidates[0].Content.Parts[0].Text, nil
}

func (c *GoogleClient) GetProvider() string {
	return "google"
}

func (c *GoogleClient) GetModel() string {
	return c.Model
}
