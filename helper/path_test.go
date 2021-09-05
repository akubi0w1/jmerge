package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanJoinPath(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		out  string
	}{
		{
			name: "current",
			in:   []string{"."},
			out:  ".",
		},
		{
			name: "only filename",
			in:   []string{"jmerge.yaml"},
			out:  "jmerge.yaml",
		},
		{
			name: "ref path",
			in:   []string{"./example", "conf", "jmerge.yaml"},
			out:  "example/conf/jmerge.yaml",
		},
		{
			name: "abs path",
			in:   []string{"/home", "example", "conf", "jmerge.yaml"},
			out:  "/home/example/conf/jmerge.yaml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := CleanJoinPath(tt.in...)
			assert.Equal(t, tt.out, out)
		})
	}
}
