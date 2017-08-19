package coreutils

import (
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
)

// ExecCommand executes a command with args and returning the stringified output
func ExecCommand(command string, args []string, redirect bool) string {
	if ExecutableExists(command) { // If the executable exists
		currentUser, _ := user.Current()

		var output []byte
		runner := exec.Command(command, args...)

		uidInt, _ := strconv.Atoi(currentUser.Uid)
		gidInt, _ := strconv.Atoi(currentUser.Gid)

		runner.SysProcAttr = &syscall.SysProcAttr{}
		runner.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uidInt), Gid: uint32(gidInt) }

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
