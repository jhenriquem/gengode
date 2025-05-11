package ui

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)
// Function responsible for the loading signal to wait for the response
func Loading(stopChan chan bool) {
	frames := []string{"|", "/", "-", "\\"}
	i := 0
	for {
		select {
		case <-stopChan:
			fmt.Print("\r                          \r") 
			return
		default:
			color.RGB(192, 202, 245).Printf("\r     Thinking %s", frames[i%len(frames)])
			time.Sleep(200 * time.Millisecond)
			i++
		}
	}
}
