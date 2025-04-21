package tools

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

func isWSL() bool {
	if runtime.GOOS != "linux" {
		return false
	}
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(data)), "microsoft")
}

func Typecmd(text string) error {
	if text == "" {
		return fmt.Errorf("commit message cannot be empty")
	}

	const delaySeconds = 0
	time.Sleep(delaySeconds * time.Second)

	escapedText := strings.ReplaceAll(text, "\"", "\\\"")
	command := fmt.Sprintf("git commit -m \"%s\"", escapedText)

	if isWSL() {
		fmt.Println("WSL2 detected. Automatic typing is not supported in WSL.")
		fmt.Println("Please copy and paste the command below into your terminal:")
		fmt.Println(command)

		return nil
	}

	robotgo.TypeStr(command)
	return nil
}
