package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	services "github.com/HublastX/Commit-IA/services/configPath"
	"github.com/HublastX/Commit-IA/tools"
)

func UpdateConfig() error {
	existingConfig, _ := services.LoadConfig()

	fmt.Println("=== CommitIA Configuration ===")

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
		return fmt.Errorf("error reading response: %v", err)
	}

	useRemote := selection == options[0]

	if useRemote {
		if existingConfig != nil {
			if !existingConfig.UseRemote {
				existingConfig.UseRemote = true
				if err := services.SaveConfig(existingConfig); err != nil {
					return err
				}
				fmt.Println("Switched to remote web service.")
			} else {
				fmt.Println("Keeping existing remote configuration.")
			}
			return nil
		}

		config := tools.CreateRemoteConfig()
		if err := services.SaveConfig(config); err != nil {
			return err
		}

		fmt.Println("Default web service configuration saved successfully!")
		return nil
	}

	config, err := configureLocalLLM()
	if err != nil {
		return err
	}

	if err := services.SaveConfig(config); err != nil {
		return err
	}

	fmt.Println("Local configuration updated successfully!")
	return nil
}
