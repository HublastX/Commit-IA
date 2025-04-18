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

	var useLocal bool
	prompt := &survey.Confirm{
		Message: "Do you want to configure a local LLM model? (No = use default web service)",
		Default: false,
	}

	if err := survey.AskOne(prompt, &useLocal); err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var config *schemas.LLMConfig
	var err error

	if useLocal {
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
