package jmerge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeMap(t *testing.T) {
	type in struct {
		base    map[string]interface{}
		overlay map[string]interface{}
		mode    MergeMode
	}
	tests := []struct {
		name string
		in   in
		out  map[string]interface{}
	}{
		{
			name: "merge mode add",
			in: in{
				base: map[string]interface{}{
					"hoge": "hoge",
					"fuga": 1,
				},
				overlay: map[string]interface{}{
					"hoge": "hoge001",
					"piyo": 1.1,
				},
				mode: MergeModeAdd,
			},
			out: map[string]interface{}{
				"hoge": "hoge001",
				"fuga": 1,
				"piyo": 1.1,
			},
		},
		{
			name: "merge mode not add",
			in: in{
				base: map[string]interface{}{
					"hoge": "hoge",
					"fuga": 1,
				},
				overlay: map[string]interface{}{
					"hoge": "hoge001",
					"piyo": 1.1,
				},
				mode: MergeModeIgnore,
			},
			out: map[string]interface{}{
				"hoge": "hoge001",
				"fuga": 1,
			},
		},
		{
			name: "merge nested json",
			in: in{
				base: map[string]interface{}{
					"hoge": "hoge",
					"fuga": map[string]interface{}{
						"fizz": "fizz",
						"buzz": "buzz",
					},
				},
				overlay: map[string]interface{}{
					"hoge": "hoge001",
					"fuga": map[string]interface{}{
						"fizz": "fizz001",
					},
				},
			},
			out: map[string]interface{}{
				"hoge": "hoge001",
				"fuga": map[string]interface{}{
					"fizz": "fizz001",
					"buzz": "buzz",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := mergeMap(tt.in.base, tt.in.overlay, tt.in.mode)
			assert.Equal(t, tt.out, out)
		})
	}
}
