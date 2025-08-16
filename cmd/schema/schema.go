package schemas

type LLMConfig struct {
	Provider         string `json:"provider"`
	Model            string `json:"model"`
	APIKey           string `json:"api_key"`
	UseRemote        bool   `json:"use_remote"`
	CommitType       int    `json:"commit_type"`
	CustomFormatText string `json:"custom_format_text,omitempty"`
}

type ProviderInfo struct {
	Name   string   `json:"name"`
	Models []string `json:"models"`
	EnvVar string   `json:"env_var"`
}

type RequestPayload struct {
	Message string `json:"message"`
}

type ResponsePayload struct {
	Response string `json:"response"`
}

type CommitAnalyzerRequest struct {
	CodeChanges string `json:"code_changes"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
	Language    string `json:"language"`
}
