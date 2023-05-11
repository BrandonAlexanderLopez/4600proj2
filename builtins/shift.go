// made by Nathan Mann
package builtins

import (
	//"errors"
	"fmt"
	//"io"
	//"os"
	//"path/filepath"
)

//var (
//	ErrInvalidArg = errors.New("invalid argument")
//)

func LeftShift(n int, args ...string) error {
	// shift the command line arguments to the left
	//var s []string
	
	for n > 0 {
		for len(args) > 0 {
			//x := args[0]
			args = args[1:]
			//return args
		}
	}

	fmt.Println(args)
	return nil//args
}