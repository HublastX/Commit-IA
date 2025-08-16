package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	schemas "github.com/HublastX/Commit-IA/schema"
	configpath "github.com/HublastX/Commit-IA/services/configPath"
	"github.com/HublastX/Commit-IA/tools"
)

func FirstTimeSetup() (*schemas.LLMConfig, error) {
	fmt.Println("Welcome to CommitIA! Let's configure the LLM service.")

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
		return nil, fmt.Errorf("error reading response: %v", err)
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
		return nil, fmt.Errorf("error reading commit type response: %v", err)
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
			return nil, fmt.Errorf("error reading custom format text: %v", err)
		}
	default:
		commitType = 1
	}

	var config *schemas.LLMConfig
	var err error

	if !useRemote {
		config, err = configureLocalLLM()
		if err != nil {
			return nil, err
		}
	} else {
		config = tools.CreateRemoteConfig()
	}

	config.CommitType = commitType
	config.CustomFormatText = customFormatText

	if err := configpath.SaveConfig(config); err != nil {
		return nil, err
	}

	fmt.Println("Configuration saved successfully!")
	return config, nil
}
