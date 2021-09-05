package jmerge

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/akubi0w1/jmerge/helper"
)

// MergeJSONByFile
func MergeJSONByFile(baseFilePath, overlayFilePath string, mode MergeMode, isFormat bool) ([]byte, error) {
	// read base file
	base, err := helper.ReadFile(baseFilePath)
	if err != nil {
		return nil, err
	}
	baseMap := make(map[string]interface{})
	if err = json.Unmarshal([]byte(base), &baseMap); err != nil {
		return nil, err
	}

	// read overlay file
	overlay, err := helper.ReadFile(overlayFilePath)
	if err != nil {
		return nil, err
	}
	overlayMap := make(map[string]interface{})
	if err = json.Unmarshal([]byte(overlay), &overlayMap); err != nil {
		return nil, err
	}

	resultMap := mergeMap(baseMap, overlayMap, mode)
	result, err := json.Marshal(resultMap)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	if isFormat {
		err = json.Indent(&out, result, "", "  ")
		if err != nil {
			return nil, err
		}
	} else {
		if _, err = out.Write(result); err != nil {
			return nil, err
		}
	}
	return out.Bytes(), nil
}

// MergeMode is behavior of merge.
// add-mode add and merge values not in base.
// ignore-mode do not add values that are not in base.
type MergeMode string

const (
	// MergeModeAdd add values not in base.
	MergeModeAdd MergeMode = "add"

	// MergeMode do not add values that are not in base.
	MergeModeIgnore MergeMode = "ignore"
)

// Validate validates merge mode
func (m MergeMode) Validate() error {
	if m == MergeModeAdd ||
		m == MergeModeIgnore {
		return nil
	}
	return fmt.Errorf("invalid merge mode=%s: merge mode is add or ignore", m)
}

// mergeMap overwrite with "overlay" based on "base"
func mergeMap(base, overlay map[string]interface{}, mode MergeMode) map[string]interface{} {
	result := base
	if mode == MergeModeAdd {
		for k, v := range overlay {
			result[k] = v
		}
	} else {
		for k, v := range overlay {
			if result[k] != nil {
				result[k] = v
			}
		}
	}
	return result
}
