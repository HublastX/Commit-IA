package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	schemas "github.com/HublastX/Commit-IA/schema"
)

type DataStore struct {
	DataDir string
}

type CommitRequest struct {
	Timestamp   time.Time `json:"timestamp"`
	Provider    string    `json:"provider"`
	Model       string    `json:"model"`
	CodeChanges string    `json:"code_changes"`
	Description string    `json:"description"`
	Tag         string    `json:"tag"`
	Language    string    `json:"language"`
	Response    string    `json:"response"`
}

type ProviderUsage struct {
	Provider    string    `json:"provider"`
	Model       string    `json:"model"`
	UsageCount  int       `json:"usage_count"`
	LastUsed    time.Time `json:"last_used"`
	TotalTokens int       `json:"total_tokens,omitempty"`
	TotalCost   float64   `json:"total_cost,omitempty"`
}

func NewDataStore() *DataStore {
	return &DataStore{
		DataDir: "../../data",
	}
}

func (ds *DataStore) ensureDataDir() error {
	return os.MkdirAll(ds.DataDir, 0755)
}

func (ds *DataStore) SaveCommitRequest(provider, model, codeChanges, description, tag, language, response string) error {
	if err := ds.ensureDataDir(); err != nil {
		return fmt.Errorf("failed to create data directory: %v", err)
	}

	commit := CommitRequest{
		Timestamp:   time.Now(),
		Provider:    provider,
		Model:       model,
		CodeChanges: codeChanges,
		Description: description,
		Tag:         tag,
		Language:    language,
		Response:    response,
	}

	// Save individual request
	filename := fmt.Sprintf("commit_%d.json", time.Now().Unix())
	filepath := filepath.Join(ds.DataDir, filename)

	data, err := json.MarshalIndent(commit, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal commit data: %v", err)
	}

	if err := os.WriteFile(filepath, data, 0644); err != nil {
		return fmt.Errorf("failed to write commit file: %v", err)
	}

	// Update usage stats
	return ds.updateProviderUsage(provider, model)
}

func (ds *DataStore) updateProviderUsage(provider, model string) error {
	usageFile := filepath.Join(ds.DataDir, "provider_usage.json")

	var usageStats []ProviderUsage

	// Try to read existing usage stats
	if data, err := os.ReadFile(usageFile); err == nil {
		json.Unmarshal(data, &usageStats)
	}

	// Find existing entry or create new one
	found := false
	for i := range usageStats {
		if usageStats[i].Provider == provider && usageStats[i].Model == model {
			usageStats[i].UsageCount++
			usageStats[i].LastUsed = time.Now()
			found = true
			break
		}
	}

	if !found {
		usageStats = append(usageStats, ProviderUsage{
			Provider:   provider,
			Model:      model,
			UsageCount: 1,
			LastUsed:   time.Now(),
		})
	}

	// Save updated stats
	data, err := json.MarshalIndent(usageStats, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal usage stats: %v", err)
	}

	return os.WriteFile(usageFile, data, 0644)
}

func (ds *DataStore) GetProviderUsage() ([]ProviderUsage, error) {
	usageFile := filepath.Join(ds.DataDir, "provider_usage.json")

	data, err := os.ReadFile(usageFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []ProviderUsage{}, nil
		}
		return nil, fmt.Errorf("failed to read usage file: %v", err)
	}

	var usageStats []ProviderUsage
	if err := json.Unmarshal(data, &usageStats); err != nil {
		return nil, fmt.Errorf("failed to unmarshal usage stats: %v", err)
	}

	return usageStats, nil
}

func (ds *DataStore) SaveConfig(config *schemas.LLMConfig) error {
	if err := ds.ensureDataDir(); err != nil {
		return fmt.Errorf("failed to create data directory: %v", err)
	}

	configFile := filepath.Join(ds.DataDir, "current_config.json")

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}

	return os.WriteFile(configFile, data, 0644)
}

func (ds *DataStore) LoadConfig() (*schemas.LLMConfig, error) {
	configFile := filepath.Join(ds.DataDir, "current_config.json")

	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config schemas.LLMConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return &config, nil
}
