package configpath

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func getConfigFilePath() (string, error) {

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("error getting current file path")
	}

	baseDir := filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(currentFile))))

	configDir := filepath.Join(baseDir, "bot", "app", "config")

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", fmt.Errorf("error creating config directory: %v", err)
	}

	return filepath.Join(configDir, "config.json"), nil
}
