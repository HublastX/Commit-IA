package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	services "github.com/HublastX/Commit-IA/services/configPath"
	"github.com/HublastX/Commit-IA/tools"
)

func UpdateConfig() error {
	existingConfig, _ := services.LoadConfig()

	var useLocal bool
	prompt := &survey.Confirm{
		Message: "Do you want to configure a local LLM model? (No = keep current config or use web service)",
		Default: false,
	}

	if err := survey.AskOne(prompt, &useLocal); err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	if !useLocal {
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

	fmt.Println("Configuration updated successfully!")
	return nil
}
