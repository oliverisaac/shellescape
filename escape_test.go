package shellescape

import (
	"reflect"
	"testing"
)

func TestEscapeArgs(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		wantEscaped []string
	}{
		{
			name:        "Test basic string",
			args:        []string{`hello world`},
			wantEscaped: []string{`hello\ world`},
		},
		{
			name:        "Test multiple args",
			args:        []string{`vim`, `hello world`},
			wantEscaped: []string{`vim`, `hello\ world`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEscaped := EscapeArgs(tt.args); !reflect.DeepEqual(gotEscaped, tt.wantEscaped) {
				t.Errorf("EscapeArgs(), got: %v, want: %v", gotEscaped, tt.wantEscaped)
			}
		})
	}
}
