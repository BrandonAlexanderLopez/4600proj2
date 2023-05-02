package builtins

import (
	"fmt"
	"io"
)

func Echo(w io.Writer, args ...string) error {
	_, err := fmt.Fprintf(w, "%s\n", args)
	return err
}
