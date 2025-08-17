package commitprompts

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type GitEmoji struct {
	Emoji       string `json:"emoji"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type GitEmojiData struct {
	Gitmojis []GitEmoji `json:"gitmojis"`
}

func LoadGitEmojis() ([]GitEmoji, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("error getting current file path")
	}

	emojiDir := filepath.Dir(currentFile)
	emojiPath := filepath.Join(emojiDir, "git_emoji.json")

	content, err := os.ReadFile(emojiPath)
	if err != nil {
		return nil, fmt.Errorf("error reading emoji file: %v", err)
	}

	var emojiData GitEmojiData
	if err := json.Unmarshal(content, &emojiData); err != nil {
		return nil, fmt.Errorf("error parsing emoji JSON: %v", err)
	}

	return emojiData.Gitmojis, nil
}

func GetEmojiPromptAddition() (string, error) {
	emojis, err := LoadGitEmojis()
	if err != nil {
		return "", err
	}

	emojiText := "\n\nEMOJIS DISPONÍVEIS (use o emoji apropriado no início da mensagem):\n"
	for _, emoji := range emojis {
		emojiText += fmt.Sprintf("- %s %s: %s\n", emoji.Emoji, emoji.Code, emoji.Description)
	}
	emojiText += "\nEscolha o emoji mais apropriado baseado no tipo de mudança e adicione no início da mensagem de commit.\n"

	return emojiText, nil
}
