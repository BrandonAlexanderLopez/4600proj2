package builtins

import (
	"os/exec"
)

func SH(args ...string) error {
	cmd := exec.Command("sh", "-c", args[0])
	cmd.Stderr = cmd.Stderr
	cmd.Stdout = cmd.Stdout
	return cmd.Run()
}
