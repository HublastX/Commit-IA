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
		"Local: Faster response times but requires provider, model, and API key configuration + Docker",
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

	if err := configpath.SaveConfig(config); err != nil {
		return nil, err
	}

	fmt.Println("Configuration saved successfully!")
	return config, nil
}
