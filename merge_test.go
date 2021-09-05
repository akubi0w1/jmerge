package jmerge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeJSON_success(t *testing.T) {
	type in struct {
		base     []byte
		overlay  []byte
		mode     MergeMode
		isFormat bool
	}
	tests := []struct {
		name string
		in   in
		out  []byte
	}{
		{
			name: "merge mode add",
			in: in{
				base: []byte(`{
					"hoge": "hoge",
					"fuga": 1
				}`),
				overlay: []byte(`{
					"hoge": "hoge001",
					"piyo": 1.1
				}`),
				mode:     MergeModeAdd,
				isFormat: false,
			},
			out: []byte(`{"fuga":1,"hoge":"hoge001","piyo":1.1}`),
		},
		{
			name: "merge mode ignore",
			in: in{
				base: []byte(`{
					"hoge": "hoge",
					"fuga": 1
				}`),
				overlay: []byte(`{
					"hoge": "hoge001",
					"piyo": 1.1
				}`),
				mode:     MergeModeIgnore,
				isFormat: false,
			},
			out: []byte(`{"fuga":1,"hoge":"hoge001"}`),
		},
		{
			name: "merge nested json (mode is ignore)",
			in: in{
				base: []byte(`{
					"hoge": "hoge",
					"fuga": {
						"fizz": "fizz",
						"buzz": "buzz"
					}
				}`),
				overlay: []byte(`{
					"hoge": "hoge001",
					"fuga": {
						"fizz": "fizz001"
					}
				}`),
				mode:     MergeModeIgnore,
				isFormat: false,
			},
			out: []byte(`{"fuga":{"buzz":"buzz","fizz":"fizz001"},"hoge":"hoge001"}`),
		},
		{
			name: "merge nested json (mode is add)",
			in: in{
				base: []byte(`{
					"hoge": "hoge",
					"fuga": {
						"fizz": "fizz",
						"buzz": "buzz"
					}
				}`),
				overlay: []byte(`{
					"hoge": "hoge001",
					"fuga": {
						"fizz": "fizz001",
						"piyo": 1
					}
				}`),
				mode:     MergeModeAdd,
				isFormat: false,
			},
			out: []byte(`{"fuga":{"buzz":"buzz","fizz":"fizz001","piyo":1},"hoge":"hoge001"}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, _ := MergeJSON(tt.in.base, tt.in.overlay, tt.in.mode, tt.in.isFormat)
			assert.EqualValues(t, tt.out, out)
		})
	}
}

func TestMergeJSON_failure(t *testing.T) {
	type in struct {
		base     []byte
		overlay  []byte
		mode     MergeMode
		isFormat bool
	}
	tests := []struct {
		name string
		in   in
	}{
		{
			name: "failed to base unmarshal",
			in: in{
				base: []byte(``),
				overlay: []byte(`{
					"hoge": "hoge001",
					"piyo": 1.1
				}`),
				mode:     MergeModeAdd,
				isFormat: false,
			},
		},
		{
			name: "failed to overlay unmarshal",
			in: in{
				base: []byte(`{
					"hoge": "hoge",
					"fuga": 1
				}`),
				overlay:  []byte(``),
				mode:     MergeModeAdd,
				isFormat: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := MergeJSON(tt.in.base, tt.in.overlay, tt.in.mode, tt.in.isFormat)
			assert.NotNil(t, err)
		})
	}
}

func TestMergeMode_Validate(t *testing.T) {
	tests := []struct {
		name    string
		in      MergeMode
		isError bool
	}{
		{
			name:    "mode is add",
			in:      MergeModeAdd,
			isError: false,
		},
		{
			name:    "mode is ignore",
			in:      MergeModeIgnore,
			isError: false,
		},
		{
			name:    "mode is invalid",
			in:      "invalid",
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.in.Validate()
			if tt.isError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

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
