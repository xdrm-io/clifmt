package bold

import (
	"fmt"
	"regexp"
	"strings"
)

type export string

var Export = export("bold")

func (syn export) Regex() *regexp.Regexp {
	return regexp.MustCompile(`(?m)\*\*((?:[^\*]+\*?)+)\*\*`)
}

func (syn export) Transform(args ...string) (string, error) {
	// no arg, empty -> ignore
	if len(args) < 1 || len(args[0]) < 1 {
		return "", nil
	}

	return fmt.Sprintf("\x1b[1m%s\x1b[22m", strings.Replace(args[0], "\x1b[0m", "\x1b[0m\x1b[1m", -1)), nil
}
