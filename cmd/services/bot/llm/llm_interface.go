package llm

import (
	"fmt"
	"strings"

	commitprompts "github.com/HublastX/Commit-IA/services/bot/commitPrompts"
)

type LLMClient interface {
	GenerateCommitMessage(prompt string) (string, error)
	GetProvider() string
	GetModel() string
	SetAPIKey(apiKey string)
}

type CommitAnalyzer struct {
	client LLMClient
}

func NewCommitAnalyzer(provider, model, apiKey string) (*CommitAnalyzer, error) {
	var client LLMClient

	switch provider {
	case "openai":
		client = NewOpenAIClient(model)
	case "google":
		client = NewGoogleClient(model)
	default:
		return nil, fmt.Errorf("unsupported provider: %s", provider)
	}

	if apiKey != "" {
		client.SetAPIKey(apiKey)
	}

	return &CommitAnalyzer{
		client: client,
	}, nil
}

func (ca *CommitAnalyzer) AnalyzeCommit(codeChanges, description, tag, language string, commitType int, customFormatText string) (string, error) {
	var promptTemplate string
	var prompt string
	var err error

	if commitType == 4 && customFormatText != "" {
		promptTemplate, err = commitprompts.GetCustomPrompt(customFormatText)
		if err != nil {
			return "", fmt.Errorf("failed to get custom prompt: %v", err)
		}
		prompt = fmt.Sprintf(promptTemplate, codeChanges, language, description, tag, customFormatText)
	} else {
		promptTemplate, err = commitprompts.GetPrompt(commitType)
		if err != nil {
			return "", fmt.Errorf("failed to get prompt: %v", err)
		}

		if commitType == 3 {
			prompt = fmt.Sprintf(promptTemplate, codeChanges, language, description)
		} else {
			prompt = fmt.Sprintf(promptTemplate, codeChanges, language, description, tag)
		}
	}

	response, err := ca.client.GenerateCommitMessage(prompt)
	if err != nil {
		return "", fmt.Errorf("failed to generate commit message: %v", err)
	}

	return strings.TrimSpace(response), nil
}

func (ca *CommitAnalyzer) GetProvider() string {
	return ca.client.GetProvider()
}

func (ca *CommitAnalyzer) GetModel() string {
	return ca.client.GetModel()
}
