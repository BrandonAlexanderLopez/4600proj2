package builtins

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	ErrInvalidArg = errors.New("invalid argument")
)

func PresentWorkingDirectory(w io.Writer, args ...string) error {
	// get the current working directory
	if len(args) > 1 {
		return fmt.Errorf("%w: pwd accepts up to one argument", ErrInvalidArgCount)
	}

	if len(args) == 1 {
		// -L behavior is implicit
		if args[0] == "-P" {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}

			path, err := filepath.EvalSymlinks(wd)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintln(w, path)
			return err
		} else if args[0] != "-L" {
			return fmt.Errorf("%w: pwd accepts -L or -P as an argument", ErrInvalidArg)
		}
	}

	// handle L flag or default behavior
	pwd := os.Getenv("PWD")
	_, err := fmt.Fprintln(w, pwd)
	return err

}
