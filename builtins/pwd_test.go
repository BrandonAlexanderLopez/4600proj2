package builtins_test

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/BrandonAlexanderLopez/4600proj2/builtins"
)

func TestPresentWorkingDirectory(t *testing.T) {
	cwd, _ := os.Getwd()
	lwd := os.Getenv("PWD")
	rawPath, _ := filepath.EvalSymlinks(cwd)
	tests := []struct {
		name    string
		args    []string
		wantOut string
		wantErr error
	}{
		{
			name:    "success no args",
			args:    []string{},
			wantOut: lwd + "\n",
		},
		{
			name:    "success with -L arg",
			args:    []string{"-L"},
			wantOut: lwd + "\n",
		},
		{
			name:    "success with -P arg",
			args:    []string{"-P"},
			wantOut: rawPath + "\n",
		},
		{
			name:    "error too many args",
			args:    []string{"-L", "-P"},
			wantErr: builtins.ErrInvalidArgCount,
		},
		{
			name:    "error invalid arg",
			args:    []string{"-invalid"},
			wantErr: builtins.ErrInvalidArg,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer

			if err := builtins.PresentWorkingDirectory(&out, tt.args...); tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("PresentWorkingDirectory() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if err != nil {
				t.Fatalf("PresentWorkingDirectory() unexpected error: %v", err)
			}
			if got := out.String(); got != tt.wantOut {
				t.Errorf("PresentWorkingDirectory() = %v, want %v", got, tt.wantOut)
			}
		})
	}
}
