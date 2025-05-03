package render

import (
	"fmt"
	"os"

	markdown "github.com/MichaelMure/go-term-markdown"
)

func MarkdownFile() {
	path := "gengode_temp_response.md"
	source, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	result := markdown.Render(string(source), 80, 1)

	fmt.Print("\n     💡")
	fmt.Print(string(result))
}
