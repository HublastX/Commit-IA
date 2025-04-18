package configpath

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	schemas "github.com/HublastX/Commit-IA/schema"
)

func SaveConfig(config *schemas.LLMConfig) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("error serializing config: %v", err)
	}

	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("error creating config directory: %v", err)
	}

	if err := ioutil.WriteFile(configPath, data, 0600); err != nil {
		return fmt.Errorf("error writing config file: %v", err)
	}

	fmt.Println("Configuration saved successfully to:", configPath)
	return nil
}
