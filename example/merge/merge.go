package main

import (
	"fmt"

	"github.com/akubi0w1/jmerge"
	"github.com/akubi0w1/jmerge/helper"
)

func main() {
	// prepare base
	base, err := helper.ReadFile("./base.json")
	if err != nil {
		panic(err)
	}

	// prepare overlay
	overlay, err := helper.ReadFile("./overlay.json")
	if err != nil {
		panic(err)
	}

	// merge
	out, err := jmerge.MergeJSON(base, overlay, jmerge.MergeModeIgnore, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
