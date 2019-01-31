package hyperlink

import (
	"fmt"
	"regexp"
)

type export string

var Export = export("hyperlink")

func (syn export) Regex() *regexp.Regexp {
	return regexp.MustCompile(`(?m)([^\033])\[([^(?:\]()]+)\]\(([^\)]+)\)`)
}

func (syn export) Transform(args ...string) (string, error) {
	// no arg, empty -> ignore
	if len(args) < 2 || len(args[0]) < 1 {
		return "", nil
	}

	linkstart := fmt.Sprintf("\x1b]8;;%s\x1b\\", args[2])
	linkend := fmt.Sprintf("\x1b]8;;\x1b\\")

	return fmt.Sprintf("%s%s%s%s", args[0], linkstart, args[1], linkend), nil
}
