package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/HublastX/Commit-IA/global"
	schemas "github.com/HublastX/Commit-IA/schema"
)

func configureLocalLLM() (*schemas.LLMConfig, error) {
	config := &schemas.LLMConfig{
		UseRemote: false,
	}

	var providerNames []string
	for _, p := range global.Providers {
		providerNames = append(providerNames, p.Name)
	}

	var providerName string
	prompt := &survey.Select{
		Message: "Choose LLM provider:",
		Options: providerNames,
	}
	if err := survey.AskOne(prompt, &providerName); err != nil {
		return nil, fmt.Errorf("error selecting provider: %v", err)
	}

	provider := FindProviderByName(providerName)
	if provider == nil {
		return nil, fmt.Errorf("invalid provider: %s", providerName)
	}

	config.Provider = providerName

	var modelName string
	modelPrompt := &survey.Select{
		Message: "Choose LLM model:",
		Options: provider.Models,
	}
	if err := survey.AskOne(modelPrompt, &modelName); err != nil {
		return nil, fmt.Errorf("error selecting model: %v", err)
	}

	config.Model = modelName

	var apiKey string
	apiKeyPrompt := &survey.Password{
		Message: fmt.Sprintf("Enter your API key for %s (%s):", providerName, provider.EnvVar),
	}
	if err := survey.AskOne(apiKeyPrompt, &apiKey); err != nil {
		return nil, fmt.Errorf("error reading API key: %v", err)
	}

	config.APIKey = apiKey
	return config, nil
}
