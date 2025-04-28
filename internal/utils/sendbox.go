package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SendBox() string {
	reader := bufio.NewReader(os.Stdin)

	var message string

	fmt.Println("╭────────────────────────────────────────────────")

	fmt.Print("│   󱖰  ")
	message, _ = reader.ReadString('\n')
	message = strings.TrimSpace(message)

	fmt.Println("╰────────────────────────────────────────────────")

	return message
}
