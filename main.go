package main

import (
	"gengode/internal/file"
	"gengode/internal/gengode"
	"gengode/internal/parser"
	"gengode/internal/render"
	"gengode/internal/ui"
	"gengode/internal/utils"
	"log"
	"strings"

	"github.com/joho/godotenv"
)

// Initial function
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var RunningApp bool = true

func main() {
	// Introduction
	ui.Introduction()

	for RunningApp {
		// Prompt
		message := ui.Prompt()

		// Check the message
		if goOn := parser.VerifyMessage(message); !goOn {
			break
		}

		// Show loading
		stop := make(chan bool)
		go ui.Loading(stop)

		// Resquest
		result, err := gengode.RequestForIA(message)

		// after loading
		stop <- true

		if err != nil {
			utils.PrintError(err.Error())
		} else if len(result.Choices) > 0 {
			text := strings.Split(result.Choices[0].Message.Content, "\n")

			err = file.CreateFile(text)
			if err != nil {
				utils.PrintError(err.Error())
			}

			render.MarkdownFile()

		}
	}
}
