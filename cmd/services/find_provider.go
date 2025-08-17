package services

import (
	"github.com/HublastX/Commit-IA/global"
	schemas "github.com/HublastX/Commit-IA/schema"
)

func FindProviderByName(name string) *schemas.ProviderInfo {
	for _, p := range global.Providers {
		if p.Name == name {
			return &p
		}
	}
	return nil
}
