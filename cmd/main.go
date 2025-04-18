package main

import (
	"fmt"
	"os"

	services "github.com/HublastX/Commit-IA/services"
	"github.com/HublastX/Commit-IA/services/cli"
	configpath "github.com/HublastX/Commit-IA/services/configPath"
	"github.com/HublastX/Commit-IA/tools"
)

func main() {

	for _, arg := range os.Args[1:] {
		if arg == "--update" {
			if err := services.UpdateConfig(); err != nil {
				fmt.Printf("Error updating configuration: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Configuration updated. Run 'commitia' to use.")
			os.Exit(0)
		}
	}

	config, err := configpath.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		return
	}

	if config == nil {
		config, err = services.FirstTimeSetup()
		if err != nil {
			fmt.Printf("Error in initial configuration: %v\n", err)
			return
		}
		fmt.Println("Initial configuration complete. Run 'commitia' to use.")
		os.Exit(0)
	}

	if !config.UseRemote {
		provider := services.FindProviderByName(config.Provider)
		if provider != nil {
			os.Setenv(provider.EnvVar, config.APIKey)
		}
	}

	projectPath, err := tools.GetProjectPath()
	if err != nil {
		fmt.Println(err)
		return
	}

	outDiff, err := tools.RunGitDiff(projectPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	serviceURL := services.GetServiceURL(config)

	if config.UseRemote {
		fmt.Println("Using remote service:", serviceURL)
	} else {
		fmt.Println("Using local service:", serviceURL)
	}

	rootCmd := cli.ExecuteCLI(outDiff, serviceURL)

	rootCmd.Flags().StringP("description", "d", "", "Basic description of what was done in the commit")
	rootCmd.Flags().StringP("language", "l", "português", "Language in which the commit should be written")
	rootCmd.Flags().StringP("tag", "t", "", "Semantic commit tag e.g.: feat, fix, chore, etc")
	rootCmd.Flags().Bool("update", false, "Update LLM configuration")
	rootCmd.Flags().BoolP("toggle-mode", "m", false, "Toggle between local and remote mode")

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error executing CLI command: %v\n", err)
		os.Exit(1)
	}
}
