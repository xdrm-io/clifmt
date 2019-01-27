package hyperlink

import (
	"fmt"
	"regexp"
)

type export string

var Export = export("hyperlink")

func (syn export) Regex() *regexp.Regexp {
	return regexp.MustCompile(`(?m)\[([^\[]+)\]\(([^\)]+)\)`)
}

func (syn export) Transform(args ...string) (string, error) {
	// no arg, empty -> ignore
	if len(args) < 2 || len(args[0]) < 1 {
		return "", nil
	}

	return fmt.Sprintf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", args[1], args[0]), nil
}
