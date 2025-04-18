package global

import schemas "github.com/HublastX/Commit-IA/schema"

var Providers = []schemas.ProviderInfo{
	{
		Name:   "google",
		Models: []string{"gemini-pro", "gemini-1.5-pro", "gemini-2.0-flash"},
		EnvVar: "GOOGLE_API_KEY",
	},
	{
		Name:   "openai",
		Models: []string{"gpt-3.5-turbo", "gpt-4", "gpt-4o"},
		EnvVar: "OPENAI_API_KEY",
	},
}
