package main

import (
	"fmt"

	"github.com/akubi0w1/jmerge"
)

func main() {
	basePath := "./base.json"
	overlayPath := "./overlay.json"

	out, err := jmerge.MergeJSONByFile(basePath, overlayPath, jmerge.MergeModeIgnore, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
