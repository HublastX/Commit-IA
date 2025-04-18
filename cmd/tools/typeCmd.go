package tools

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

func Typecmd(text string) error {

	if text == "" {
		return fmt.Errorf("mensagem de commit não pode estar vazia")
	}

	const delaySeconds = 0
	time.Sleep(delaySeconds * time.Second)

	escapedText := strings.ReplaceAll(text, "\"", "\\\"")
	command := fmt.Sprintf("git commit -m \"%s\"", escapedText)

	robotgo.TypeStr(command)

	return nil
}
