package utils

import (
	"fmt"
	"os"
)

func ExitProgramn() {
	fmt.Println("\n   👋 Bye, see you later! \n")
	defer os.Remove("gengode_temp_response.md")
}
