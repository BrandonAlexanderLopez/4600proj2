package builtins_test

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/BrandonAlexanderLopez/4600proj2/builtins"
)

func TestKsh(t *testing.T) {
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
			if err := builtins.Ksh(tt.command); !errors.Is(err, tt.wantErr) {
				t.Errorf("Ksh() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
