package codeutilsShared

import (
	"os"
	"os/exec"
	"strings"
)

// ExecCommand
// Function for executing a utility with args and returning the stringified output
func ExecCommand(utility string, args []string, liveOutput bool) string {
	var output []byte
	runner := exec.Command(utility, args...)

	if liveOutput { // If we should immediately output the results of the command
		runner.Stdout = os.Stdout
		runner.Stderr = os.Stderr
		runner.Start()
	} else { // If we should redirect output to var
		output, _ = runner.CombinedOutput() // Combine the output of stderr and stdout
	}

	return string(output[:])
}

// ExecutableExists
// Function for checking if an executable exists
func ExecutableExists(executableName string) bool {
	executableExists := true // Default to true

	var emptyFlags []string
	executableCommandMessage := ExecCommand(executableName, emptyFlags, false) // Generate an empty call to the executable

	if strings.Contains(executableCommandMessage, "executable file not found") { // If executable does not exist
		executableExists = false
	}

	return executableExists
}
