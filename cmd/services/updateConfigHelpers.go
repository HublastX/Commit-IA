package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	schemas "github.com/HublastX/Commit-IA/schema"
	services "github.com/HublastX/Commit-IA/services/configPath"
)

func updateServiceType(config *schemas.LLMConfig) error {
	options := []string{
		"Web: Simple and fast to use, no extra configuration needed",
		"Local: Faster response times but requires provider, model, and API key configuration",
	}

	var selection string
	prompt := &survey.Select{
		Message: "How do you want to use CommitIA?",
		Options: options,
		Default: options[0],
	}

	if err := survey.AskOne(prompt, &selection); err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	useRemote := selection == options[0]

	if useRemote && !config.UseRemote {
		config.UseRemote = true
		config.Provider = ""
		config.Model = ""
		config.APIKey = ""
	} else if !useRemote && config.UseRemote {
		config.UseRemote = false
		newConfig, err := configureLocalLLM()
		if err != nil {
			return err
		}
		config.Provider = newConfig.Provider
		config.Model = newConfig.Model
		config.APIKey = newConfig.APIKey
	}

	if err := services.SaveConfig(config); err != nil {
		return err
	}

	fmt.Printf("Service type updated to: %s\n", map[bool]string{true: "Remote", false: "Local"}[useRemote])
	return nil
}

func updateCommitType(config *schemas.LLMConfig) error {
	commitTypeOptions := []string{
		"Tipo 1: feat(api): implementei nova rota",
		"Tipo 2: feat: implementei nova rota na api",
		"Tipo 3: implementei nova rota na api",
		"Custom: usar meu próprio formato de prompt",
	}

	var commitTypeSelection string
	commitTypePrompt := &survey.Select{
		Message: "Escolha o formato de commit que prefere:",
		Options: commitTypeOptions,
		Default: commitTypeOptions[config.CommitType-1],
	}

	if err := survey.AskOne(commitTypePrompt, &commitTypeSelection); err != nil {
		return fmt.Errorf("error reading commit type response: %v", err)
	}

	var commitType int
	var customFormatText string

	switch commitTypeSelection {
	case commitTypeOptions[0]:
		commitType = 1
	case commitTypeOptions[1]:
		commitType = 2
	case commitTypeOptions[2]:
		commitType = 3
	case commitTypeOptions[3]:
		commitType = 4
		var customTextPrompt = &survey.Multiline{
			Message: "Digite o formato de commit que você deseja (exemplo: 'tipo(escopo): descrição' ou 'Commit: descrição breve'):",
		}
		if err := survey.AskOne(customTextPrompt, &customFormatText); err != nil {
			return fmt.Errorf("error reading custom format text: %v", err)
		}
	default:
		commitType = 1
	}

	config.CommitType = commitType
	config.CustomFormatText = customFormatText

	if err := services.SaveConfig(config); err != nil {
		return err
	}

	fmt.Printf("Commit type updated to: Type %d\n", commitType)
	if customFormatText != "" {
		fmt.Printf("Custom format: %s\n", customFormatText)
	}
	return nil
}

func updateProviderAndModel(config *schemas.LLMConfig) error {
	newConfig, err := configureLocalLLM()
	if err != nil {
		return err
	}

	config.Provider = newConfig.Provider
	config.Model = newConfig.Model
	config.APIKey = newConfig.APIKey

	if err := services.SaveConfig(config); err != nil {
		return err
	}

	fmt.Printf("Provider and model updated to: %s (%s)\n", config.Provider, config.Model)
	return nil
}

func updateAPIKey(config *schemas.LLMConfig) error {
	var apiKey string
	apiKeyPrompt := &survey.Password{
		Message: fmt.Sprintf("Enter your %s API key:", config.Provider),
	}

	if err := survey.AskOne(apiKeyPrompt, &apiKey); err != nil {
		return fmt.Errorf("error reading API key: %v", err)
	}

	config.APIKey = apiKey

	if err := services.SaveConfig(config); err != nil {
		return err
	}

	fmt.Println("API key updated successfully!")
	return nil
}

func updateCompleteConfig(config *schemas.LLMConfig) error {
	options := []string{
		"Web: Simple and fast to use, no extra configuration needed",
		"Local: Faster response times but requires provider, model, and API key configuration",
	}

	var selection string
	prompt := &survey.Select{
		Message: "How do you want to use CommitIA?",
		Options: options,
		Default: options[0],
	}

	if err := survey.AskOne(prompt, &selection); err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	useRemote := selection == options[0]

	commitTypeOptions := []string{
		"Tipo 1: feat(api): implementei nova rota",
		"Tipo 2: feat: implementei nova rota na api",
		"Tipo 3: implementei nova rota na api",
		"Custom: usar meu próprio formato de prompt",
	}

	var commitTypeSelection string
	commitTypePrompt := &survey.Select{
		Message: "Escolha o formato de commit que prefere:",
		Options: commitTypeOptions,
		Default: commitTypeOptions[0],
	}

	if err := survey.AskOne(commitTypePrompt, &commitTypeSelection); err != nil {
		return fmt.Errorf("error reading commit type response: %v", err)
	}

	var commitType int
	var customFormatText string

	switch commitTypeSelection {
	case commitTypeOptions[0]:
		commitType = 1
	case commitTypeOptions[1]:
		commitType = 2
	case commitTypeOptions[2]:
		commitType = 3
	case commitTypeOptions[3]:
		commitType = 4
		var customTextPrompt = &survey.Multiline{
			Message: "Digite o formato de commit que você deseja (exemplo: 'tipo(escopo): descrição' ou 'Commit: descrição breve'):",
		}
		if err := survey.AskOne(customTextPrompt, &customFormatText); err != nil {
			return fmt.Errorf("error reading custom format text: %v", err)
		}
	default:
		commitType = 1
	}

	if !useRemote {
		newConfig, err := configureLocalLLM()
		if err != nil {
			return err
		}
		config.Provider = newConfig.Provider
		config.Model = newConfig.Model
		config.APIKey = newConfig.APIKey
	} else {
		config.Provider = ""
		config.Model = ""
		config.APIKey = ""
	}

	config.UseRemote = useRemote
	config.CommitType = commitType
	config.CustomFormatText = customFormatText

	if err := services.SaveConfig(config); err != nil {
		return err
	}

	fmt.Println("Complete configuration updated successfully!")
	return nil
}
