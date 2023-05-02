package builtins_test

import (
	"errors"
	"os"
	"testing"

	"github.com/BrandonAlexanderLopez/4600proj2/builtins"
)

func TestUnsetEnv(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		toSet   map[string]string
		wantOut map[string]string
		wantErr error
	}{
		{
			name:    "success unset one",
			args:    []string{"foo"},
			toSet:   map[string]string{"foo": "bar", "bar": "baz"},
			wantOut: map[string]string{"bar": "baz"},
		},
		{
			name:    "success unset multiple",
			args:    []string{"foo", "bar"},
			toSet:   map[string]string{"foo": "bar", "bar": "baz"},
			wantOut: map[string]string{},
		},
		{
			name:    "error no args",
			args:    []string{},
			toSet:   map[string]string{},
			wantErr: builtins.ErrInvalidArgCount,
		},
		{
			name:    "success not set",
			args:    []string{"foo"},
			toSet:   map[string]string{},
			wantOut: map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.toSet {
				t.Setenv(k, v)
			}

			if err := builtins.UnsetEnv(tt.args...); tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("UnsetEnv() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if err != nil {
				t.Fatalf("UnsetEnv() unexpected error: %v", err)
			}

			for k := range tt.toSet {
				env := os.Getenv(k)

				if tt.wantOut[k] != env {
					t.Errorf("UnsetEnv() = %v, want %v", env, tt.wantOut[k])
				}
			}
		})
	}
}
