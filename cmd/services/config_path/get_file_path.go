package configpath

import (
	"fmt"
	"os"
	"path/filepath"
)

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting user home directory: %v", err)
	}

	configDir := filepath.Join(homeDir, ".commitia")

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", fmt.Errorf("error creating config directory: %v", err)
	}

	return filepath.Join(configDir, "settings.json"), nil
}
