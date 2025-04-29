package ui

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
)

func Prompt() string {
	reader := bufio.NewReader(os.Stdin)

	var message string

	color.RGB(122, 126, 247).Print("\n   󱖰  ")
	message, _ = reader.ReadString('\n')
	message = strings.TrimSpace(message)

	return message
}
