package cli

import (
	"fmt"

	"github.com/HublastX/Commit-IA/services"
	"github.com/HublastX/Commit-IA/tools"

	"github.com/spf13/cobra"
)

func ExecuteCLI(outDiff string, url string) *cobra.Command {
	return &cobra.Command{
		Use:   "commitia",
		Short: "A CLI tool for handling commits",
		Run: func(cmd *cobra.Command, args []string) {

			description, err := cmd.Flags().GetString("description")

			if err != nil {
				fmt.Printf("Error receiving commit description: %v\n", err)
				return
			}

			language, err := cmd.Flags().GetString("language")

			if err != nil {
				fmt.Printf("Error receiving commit language: %v\n", err)
				return
			}

			tagCommit, err := cmd.Flags().GetString("tag")

			if err != nil {
				fmt.Printf("Error receiving commit tag: %v\n", err)
				return
			}

			commitMessage := services.CreateCommitMessage(outDiff, language, description, tagCommit)

			response, err := services.SendCommitAnalysisRequest(url, commitMessage, description, language, tagCommit)
			if err != nil {
				fmt.Printf("Error sending message: %v\n", err)
				return
			}

			err = tools.Typecmd(response.Response)

			if err != nil {
				fmt.Printf("Error typing command: %v\n", err)
				return
			}
		},
	}
}
