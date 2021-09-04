package helper

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

// MergeMap overwrite with "overlay" based on "base"
func MergeMap(base, overlay map[string]interface{}, mode MergeMode) map[string]interface{} {
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
