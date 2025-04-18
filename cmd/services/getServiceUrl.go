package services

import (
	"github.com/HublastX/Commit-IA/global"
	schemas "github.com/HublastX/Commit-IA/schema"
)

func GetServiceURL(config *schemas.LLMConfig) string {
	if config.UseRemote {
		return global.DefaultRemoteURL
	}
	return global.DefaultLocalURL
}
