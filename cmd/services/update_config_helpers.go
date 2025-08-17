package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	schemas "github.com/HublastX/Commit-IA/schema"
	services "github.com/HublastX/Commit-IA/services/config_path"
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

func updateEmojiUsage(config *schemas.LLMConfig) error {
	useEmojiPrompt := &survey.Confirm{
		Message: "Deseja usar emojis do Git nas mensagens de commit?",
		Default: config.UseGitEmoji,
	}

	var useEmoji bool
	if err := survey.AskOne(useEmojiPrompt, &useEmoji); err != nil {
		return fmt.Errorf("error reading emoji preference: %v", err)
	}

	config.UseGitEmoji = useEmoji

	if err := services.SaveConfig(config); err != nil {
		return err
	}

	fmt.Printf("Git emoji usage updated to: %t\n", useEmoji)
	return nil
}

func updateProviderAndModel(config *schemas.LLMConfig) error {
	updateOptions := []string{
		"Update both Provider and Model",
		"Update only Model",
	}

	var updateSelection string
	updatePrompt := &survey.Select{
		Message: "What would you like to update?",
		Options: updateOptions,
	}

	if err := survey.AskOne(updatePrompt, &updateSelection); err != nil {
		return fmt.Errorf("error reading update selection: %v", err)
	}

	if updateSelection == updateOptions[0] {
		// Update both provider and model
		newConfig, err := configureLocalLLM()
		if err != nil {
			return err
		}
		config.Provider = newConfig.Provider
		config.Model = newConfig.Model
		config.APIKey = newConfig.APIKey
	} else {
		// Update only model
		err := updateModelOnly(config)
		if err != nil {
			return err
		}
	}

	if err := services.SaveConfig(config); err != nil {
		return err
	}

	fmt.Printf("Provider and model updated to: %s (%s)\n", config.Provider, config.Model)
	return nil
}

func updateModelOnly(config *schemas.LLMConfig) error {
	provider := FindProviderByName(config.Provider)
	if provider == nil {
		return fmt.Errorf("current provider %s not found", config.Provider)
	}

	modelOptions := append(provider.Models, "Custom: digite o nome do modelo manualmente")

	var modelSelection string
	modelPrompt := &survey.Select{
		Message: "Choose LLM model:",
		Options: modelOptions,
		Default: config.Model,
	}
	if err := survey.AskOne(modelPrompt, &modelSelection); err != nil {
		return fmt.Errorf("error selecting model: %v", err)
	}

	var modelName string
	if modelSelection == "Custom: digite o nome do modelo manualmente" {
		customModelPrompt := &survey.Input{
			Message: "Digite o nome do modelo (ex: gpt-4o-mini, claude-3-sonnet, etc.):",
		}
		if err := survey.AskOne(customModelPrompt, &modelName); err != nil {
			return fmt.Errorf("error reading custom model name: %v", err)
		}
	} else {
		modelName = modelSelection
	}

	config.Model = modelName
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

	useEmojiPrompt := &survey.Confirm{
		Message: "Deseja usar emojis do Git nas mensagens de commit?",
		Default: false,
	}

	var useEmoji bool
	if err := survey.AskOne(useEmojiPrompt, &useEmoji); err != nil {
		return fmt.Errorf("error reading emoji preference: %v", err)
	}

	config.UseRemote = useRemote
	config.CommitType = commitType
	config.CustomFormatText = customFormatText
	config.UseGitEmoji = useEmoji

	if err := services.SaveConfig(config); err != nil {
		return err
	}

	fmt.Println("Complete configuration updated successfully!")
	return nil
}
