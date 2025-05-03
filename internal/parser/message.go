package parser

import (
	"fmt"
	"gengode/internal/utils"
)

func VerifyMessage(message string) bool {
	if message == "\\exit" {
		utils.ExitProgramn()
		return false
	} else if message == "\\help" {
		fmt.Println("\n╭───────────────────────────────────────────────────────╮")
		fmt.Println("│                                                       │")
		fmt.Println("│  🚀 I'm a tool built with  go                        │")
		fmt.Println("│                                                       │")
		fmt.Println("│  A simple CLI tool to generate code using AI          │")
		fmt.Println("│  💡  Describe your idea and I write the code!         │")
		fmt.Println("│                                                       │")
		fmt.Println("│  📌 Commands                                          │")
		fmt.Println("│    🔚  Type \\exit anytime to leave the program.       │")
		fmt.Println("│    ℹ️   Type \\help to see the documentation.          │")
		fmt.Println("│                                                       │")
		fmt.Println("│  💻 Created by João Henrique                          │")
		fmt.Println("│      Github Profile : https://github.com/jhenriquem  │")
		fmt.Println("│     🌐Project Repository :                            │")
		fmt.Println("│                                                       │")
		fmt.Println("╰───────────────────────────────────────────────────────╯")

	}
	return true
}
