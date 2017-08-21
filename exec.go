package coreutils

import (
	"os"
	"os/exec"
)

// ExecCommand executes a command with args and returning the stringified output
func ExecCommand(command string, args []string, redirect bool) string {
	if ExecutableExists(command) { // If the executable exists
		var output []byte
		runner := exec.Command(command, args...)

		if redirect { // If we should redirect output to var
			output, _ = runner.CombinedOutput() // Combine the output of stderr and stdout
		} else {
			runner.Stdout = os.Stdout
			runner.Stderr = os.Stderr
			runner.Start()
		}

		return string(output[:])
	} else { // If the executable doesn't exist
		return command + " is not an executable."
	}
}

// ExecutableExists checks if an executable exists
func ExecutableExists(executableName string) bool {
	_, existsErr := exec.LookPath(executableName)
	return (existsErr == nil)
}
