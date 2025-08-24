package commitprompts

import (
	"fmt"

	"github.com/HublastX/Commit-IA/services/bot/commitPrompts/emoji"
	"github.com/HublastX/Commit-IA/services/bot/commitPrompts/prompts"
)

func GetPrompt(promptType int, useEmoji bool) (string, error) {
	var basePrompt string

	switch promptType {
	case 1:
		basePrompt = prompts.Type1
	case 2:
		basePrompt = prompts.Type2
	case 3:
		basePrompt = prompts.Type3
	default:
		basePrompt = prompts.Type1
	}

	if useEmoji {
		emojiAddition, err := emoji.GetEmojiPromptAddition()
		if err != nil {
			return "", fmt.Errorf("error loading emoji data: %v", err)
		}
		basePrompt += emojiAddition
	}

	return basePrompt, nil
}

func GetCustomPrompt(customFormatText string, useEmoji bool) (string, error) {
	if customFormatText == "" {
		return "", fmt.Errorf("custom format text is empty")
	}

	basePrompt := prompts.Custom

	if useEmoji {
		emojiAddition, err := emoji.GetEmojiPromptAddition()
		if err != nil {
			return "", fmt.Errorf("error loading emoji data: %v", err)
		}
		basePrompt += emojiAddition
	}

	return basePrompt, nil
}
