package builtins

import (
	"os"
	"os/exec"
)

func Csh(args ...string) error {
	cmd := exec.Command("csh", "-c", args[0])
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
