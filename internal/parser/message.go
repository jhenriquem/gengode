package parser

import "fmt"

func VerifyMessage(message string) bool {
	if message == "\\exit" {
		fmt.Println("\n   👋 Bye, see you later! \n")
		return false
	} else if message == "\\help" {
		fmt.Println("╭───────────────────────────────────────────────────────╮")
		fmt.Println("│                                                       │")
		fmt.Println("│  🚀 I'm a tool built with  go                        │")
		fmt.Println("│                                                       │")
		fmt.Println("│  A simple CLI tool to generate code using AI          │")
		fmt.Println("│  💡  Describe your idea and I write the code!         │")
		fmt.Println("│                                                       │")
		fmt.Println("│  📌 Commands                                          │")
		fmt.Println("│    🔚  Type \\exit anytime to leave the program.      │")   // <-- Nova linha sobre sair
		fmt.Println("│    ℹ️   Type \\help to see the documentation.          │") // <-- Nova linha sobre sair
		fmt.Println("│                                                       │")
		fmt.Println("│  💻 Created by João Henrique                          │")
		fmt.Println("│      Github Profile : https://github.com/jhenriquem  │")
		fmt.Println("│     🌐Project Repository :                            │")
		fmt.Println("│                                                       │")
		fmt.Println("╰───────────────────────────────────────────────────────╯")

	}
	return true
}
