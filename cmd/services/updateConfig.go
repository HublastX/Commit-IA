package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	services "github.com/HublastX/Commit-IA/services/configPath"
)

func UpdateConfig() error {
	existingConfig, err := services.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading existing configuration: %v", err)
	}

	if existingConfig == nil {
		return fmt.Errorf("no existing configuration found. Please run initial setup first")
	}

	fmt.Println("=== CommitIA Configuration Update ===")
	fmt.Printf("Current settings:\n")
	fmt.Printf("- Service: %s\n", map[bool]string{true: "Remote", false: "Local"}[existingConfig.UseRemote])
	if !existingConfig.UseRemote {
		fmt.Printf("- Provider: %s\n", existingConfig.Provider)
		fmt.Printf("- Model: %s\n", existingConfig.Model)
	}
	fmt.Printf("- Commit Type: %d\n", existingConfig.CommitType)
	if existingConfig.CustomFormatText != "" {
		fmt.Printf("- Custom Format: %s\n", existingConfig.CustomFormatText)
	}
	fmt.Println()

	updateOptions := []string{
		"Service Type (Local/Remote)",
		"Commit Format Type",
		"Provider & Model (Local only)",
		"API Key (Local only)",
		"Complete Reconfiguration",
	}

	var updateSelection string
	updatePrompt := &survey.Select{
		Message: "What would you like to update?",
		Options: updateOptions,
	}

	if err := survey.AskOne(updatePrompt, &updateSelection); err != nil {
		return fmt.Errorf("error reading update selection: %v", err)
	}

	switch updateSelection {
	case updateOptions[0]: // Service Type
		return updateServiceType(existingConfig)
	case updateOptions[1]: // Commit Format Type
		return updateCommitType(existingConfig)
	case updateOptions[2]: // Provider & Model
		if existingConfig.UseRemote {
			fmt.Println("Provider & Model are only applicable for local configuration.")
			return nil
		}
		return updateProviderAndModel(existingConfig)
	case updateOptions[3]: // API Key
		if existingConfig.UseRemote {
			fmt.Println("API Key is only applicable for local configuration.")
			return nil
		}
		return updateAPIKey(existingConfig)
	case updateOptions[4]: // Complete Reconfiguration
		return updateCompleteConfig(existingConfig)
	default:
		return fmt.Errorf("invalid selection")
	}
}
