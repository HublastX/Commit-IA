package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	schemas "github.com/HublastX/Commit-IA/schema"
	commitprompts "github.com/HublastX/Commit-IA/services/bot/commitPrompts"
	configpath "github.com/HublastX/Commit-IA/services/config_path"
)

func SendCommitAnalysisRequest(url string, codeChanges, description, tag, language string) (*schemas.ResponsePayload, error) {
	config, err := configpath.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading configuration: %v", err)
	}

	if config != nil && !config.UseRemote {
		return ProcessLocalCommitAnalysis(config, codeChanges, description, tag, language)
	}
	prompt, err := generatePromptForRemote(config, codeChanges, description, tag, language)
	if err != nil {
		return nil, fmt.Errorf("error generating prompt: %v", err)
	}

	payload := map[string]string{
		"prompt": prompt,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error serializing prompt payload: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error making POST request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API response error: %s, body: %s", resp.Status, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var apiResponse struct {
		Response string `json:"response"`
		Error    string `json:"error,omitempty"`
	}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("error deserializing response: %v", err)
	}

	if apiResponse.Error != "" {
		return nil, fmt.Errorf("API error: %s", apiResponse.Error)
	}

	return &schemas.ResponsePayload{
		Response: apiResponse.Response,
	}, nil
}

func generatePromptForRemote(config *schemas.LLMConfig, codeChanges, description, tag, language string) (string, error) {
	var promptTemplate string
	var prompt string
	var err error

	if config.CommitType == 4 && config.CustomFormatText != "" {
		promptTemplate, err = commitprompts.GetCustomPrompt(config.CustomFormatText, config.UseGitEmoji)
		if err != nil {
			return "", fmt.Errorf("failed to get custom prompt: %v", err)
		}
		prompt = fmt.Sprintf(promptTemplate, codeChanges, language, description, tag, config.CustomFormatText)
	} else {
		promptTemplate, err = commitprompts.GetPrompt(config.CommitType, config.UseGitEmoji)
		if err != nil {
			return "", fmt.Errorf("failed to get prompt: %v", err)
		}

		if config.CommitType == 3 {
			prompt = fmt.Sprintf(promptTemplate, codeChanges, language, description)
		} else {
			prompt = fmt.Sprintf(promptTemplate, codeChanges, language, description, tag)
		}
	}

	return prompt, nil
}
