package builtins_test

import (
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/jh125486/CSCE4600/Project2/builtins"
)

func TestBash(t *testing.T) {
	tests := []struct {
		name    string
		command string
		wantErr error
	}{
		{
			name:    "success echo",
			command: "echo hello world",
		},
		{
			name:    "success ls",
			command: "ls",
		},
		{
			name:    "failure",
			command: "non-existing-command",
			wantErr: &exec.Error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := builtins.Bash(os.Stdout, tt.command); !errors.Is(err, tt.wantErr) {
				t.Errorf("Bash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
