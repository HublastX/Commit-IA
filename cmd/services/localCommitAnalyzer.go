package services

import (
	"fmt"

	schemas "github.com/HublastX/Commit-IA/schema"
	"github.com/HublastX/Commit-IA/services/bot/llm"
)

func ProcessLocalCommitAnalysis(config *schemas.LLMConfig, codeChanges, description, tag, language string) (*schemas.ResponsePayload, error) {
	analyzer, err := llm.NewCommitAnalyzer(config.Provider, config.Model, config.APIKey)
	if err != nil {
		return nil, fmt.Errorf("error creating commit analyzer: %v", err)
	}

	commitMessage, err := analyzer.AnalyzeCommit(codeChanges, description, tag, language, config.CommitType)
	if err != nil {
		return nil, fmt.Errorf("error analyzing commit: %v", err)
	}

	return &schemas.ResponsePayload{
		Response: commitMessage,
	}, nil
}
