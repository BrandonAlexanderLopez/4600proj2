package builtins

import (
	"fmt"
	"os"
)

func UnsetEnv(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("%w: expected at least one argument", ErrInvalidArgCount)
	}

	for _, arg := range args {
		err := os.Unsetenv(arg)
		if err != nil {
			return err
		}
	}

	return nil
}
