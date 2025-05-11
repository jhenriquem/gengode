package file

import (
	"bufio"
	"fmt"
	"os"
)
// Function responsible for creating and writing the response in the markdown file
func CreateFile(text []string) error {
	nameFile := "gengode_temp_response.md"

	file, err := os.Create(nameFile)
	if err != nil {
		return fmt.Errorf("Error creating response file : %s", err.Error())
	}

	writer := bufio.NewWriter(file)
	for _, line := range text {
		linetoWrite := line + "\n"
		writer.WriteString(linetoWrite)
	}
	writer.Flush()
	return nil
}
