package utils

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func PrintError(error string) {
	color.RGB(247, 118, 242).Print("    󰃤 Something went wrong")
	color.RGB(247, 118, 242).Printf("%s", error)
}

func PrintResponse(response string) {
	color.RGB(224, 175, 104).Print("\n     ")

	slice := strings.Split(response, "\n")
	for _, line := range slice {
		fmt.Print("     ", line, "\n")
	}
}
