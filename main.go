package main

import (
	"fmt"
	"gengode/internal/gengode"
	"gengode/internal/parser"
	"gengode/internal/utils"
	"log"

	"github.com/joho/godotenv"
)

// Initial function, variable assignment
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var RunningApp bool = true

func main() {
	// Introduction
	utils.Introduction()

	for RunningApp {
		// SendBox
		message := utils.SendBox()

		// Check the message
		if goOn := parser.VerifyMessage(message); !goOn {
			break
		}

		//
		result, err := gengode.RequestForIA(message)

		if err != nil {
			fmt.Printf("    󰃤 Something went wrong : %s", err.Error())
		} else if len(result.Choices) > 0 {
			fmt.Println("\n\n    ", result.Choices[0].Message.Content)
		}
	}
}
