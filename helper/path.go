package helper

import "path/filepath"

// CleanJoinPath get a clean path with elems joined
func CleanJoinPath(elems ...string) string {
	return filepath.Clean(filepath.Join(elems...))
}
