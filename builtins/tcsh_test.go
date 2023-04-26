package builtins_test

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/BrandonAlexanderLopez/4600proj2/builtins"
)

func TestTcsh(t *testing.T) {
	tests := []struct {
		name    string
		command []string
		wantErr error
	}{
		{
			name:    "success echo",
			command: []string{"echo", "hello world"},
		},
		{
			name:    "success ls",
			command: []string{"ls"},
		},
		{
			name:    "failure",
			command: []string{"non-existing-command"},
			wantErr: &exec.Error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := builtins.Tcsh(tt.command...); !errors.Is(err, tt.wantErr) {
				t.Errorf("Tcsh() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
