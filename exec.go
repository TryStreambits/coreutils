package codeutilsShared

import (
	"os/exec"
	"strings"
)

// ExecCommand
// Function for executing a utility with args and returning the stringified output
func ExecCommand(utility string, args []string) string {
	runner := exec.Command(utility, args...)
	output, _ := runner.CombinedOutput() // Combine the output of stderr and stdout
	return string(output)
}

// ExecutableExists
// Function for checking if an executable exists
func ExecutableExists(executableName string) bool {
	executableExists := true // Default to true

	var emptyFlags []string
	executableCommandMessage := ExecCommand(executableName, emptyFlags) // Generate an empty call to the executable

	if strings.Contains(executableCommandMessage, "executable file not found") { // If executable does not exist
		executableExists = false
	}

	return executableExists
}
