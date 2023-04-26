package builtins_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/BrandonAlexanderLopez/4600proj2/builtins"
)

func TestEcho(t *testing.T) {
	var (
		testString = "hello world"
		wantString = fmt.Sprintf("%s\n", testString)
	)
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "no args",
			args: nil,
		},
		{
			name: "one arg",
			args: []string{testString},
		},
		{
			name: "multiple args",
			args: []string{"this", "is", "a", "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			if err := builtins.Echo(buf, tt.args...); err != nil {
				t.Fatalf("Echo() error: %v", err)
			}
			if gotString := buf.String(); gotString != wantString {
				t.Errorf("Echo() = %q, want %q", gotString, wantString)
			}
		})
	}
}
