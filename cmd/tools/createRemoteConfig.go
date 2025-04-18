package tools

import schemas "github.com/HublastX/Commit-IA/schema"

func CreateRemoteConfig() *schemas.LLMConfig {
	return &schemas.LLMConfig{
		UseRemote: true,
	}
}
