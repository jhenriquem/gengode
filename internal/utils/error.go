package utils

import (
	"github.com/fatih/color"
)

func PrintError(error string) {
	color.RGB(247, 118, 242).Print("    󰃤 Something went wrong")
	color.RGB(247, 118, 242).Printf("%s", error)
}
