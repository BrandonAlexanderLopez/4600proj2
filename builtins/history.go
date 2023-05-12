package builtins

import (
	"fmt"
	"io"
	"os"
)

var (
	temp string
)

func History(w io.Writer, options ...string) error {

	history, err := os.ReadFile("History.txt")

	temp := string(history)
	fmt.Fprintf(w, temp)

	return err

}
