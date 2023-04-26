package builtins

import (
	"os/exec"
)

func Ksh(args ...string) error {
	cmd := exec.Command("ksh", "-c", concatArgs(args...))
	cmd.Stderr = cmd.Stderr
	cmd.Stdout = cmd.Stdout
	cmd.Stdin = cmd.Stdin
	return cmd.Run()
}
