package builtins

import (
	"os/exec"
)

func concatArgs(args ...string) string {
	var result string
	for _, arg := range args {
		result += arg + " "
	}
	return result
}

func Tcsh(args ...string) error {
	cmd := exec.Command("tcsh", "-c", concatArgs(args...))
	cmd.Stderr = cmd.Stderr
	cmd.Stdout = cmd.Stdout
	cmd.Stdin = cmd.Stdin
	return cmd.Run()
}
