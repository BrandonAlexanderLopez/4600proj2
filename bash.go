package builtins

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

func Bash(w io.Writer, args ...string) error {
    cmd := exec.Command("bash", "-c", strings.Join(args, " "))
    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout
    cmd.Stdin = os.Stdin
    return cmd.Run()
}

