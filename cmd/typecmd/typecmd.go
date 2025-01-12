package typecmd

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func Typecmd(text string) {
	time.Sleep(2 * time.Second)

	command := fmt.Sprintf("git commit -m \"%s\"", text)

	robotgo.TypeStr(command)
}
