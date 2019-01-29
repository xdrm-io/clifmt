package clifmt

import (
	"regexp"
)

var esc = regexp.MustCompile(`(?m)\[(?:\d+;)*\d+m`)

// displaySize returns the real size escaping special characters
func displaySize(s string) int {

	// 1. get actual size
	size := len(s)

	// 2. get all terminal coloring matches
	matches := esc.FindAllString(s, -1)
	for _, m := range matches {
		size -= len(m)
	}

	return size
}
