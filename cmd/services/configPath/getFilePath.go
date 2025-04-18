package configpath

import (
	"fmt"
	"os"
	"path/filepath"
)

func getConfigFilePath() (string, error) {
	configDir := filepath.Join(".", "bot", "app", "config")

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", fmt.Errorf("error creating config directory: %v", err)
	}

	return filepath.Join(configDir, "config.json"), nil
}
