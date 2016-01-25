package codeutilsShared

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// InputDialog
// Fetch input after message
func InputMessage(message string) string {
	fmt.Print(message + ": ")                    // Print without new line
	stdinReader := bufio.NewReader(os.Stdin)     // Create a new buffer IO reader that reads stdinReader
	input, _ := stdinReader.ReadString('\n')     // Read anything before a new line
	input = strings.Replace(input, "\n", "", -1) // Remove any new lines

	return input // Return the input
}

// OutputStatus
// Outputs a "check" or not check based on true / false status, along with the message
func OutputStatus(status bool, message string) {
	statusChar := "✕ "

	if status {
		statusChar = "✔ "
	}

	fmt.Println(statusChar + message)
}
