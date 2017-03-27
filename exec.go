package coreutils

import (
	"os"
	"os/exec"
)

// ExecCommand executes a utility with args and returning the stringified output
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

// ExecutableExists checks if an executable exists
func ExecutableExists(executableName string) bool {
	_, existsErr := exec.LookPath(executableName)
	return (existsErr == nil)
}
