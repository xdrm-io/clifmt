package syntax

import (
	"fmt"
	"regexp"
	"strings"
)

// Italic implements transform.Transformer for the italic syntax
type Italic struct{}

// Regex implements transform.Transformer
func (Italic) Regex() *regexp.Regexp {
	return regexp.MustCompile(`(?m)\*([^\*]+)\*`)
}

// Transform implements transform.Transformer
func (Italic) Transform(args ...string) (string, error) {
	// ignore no arg or empty
	if len(args) < 1 || len(args[0]) < 1 {
		return "", nil
	}

	return fmt.Sprintf("\x1b[3m%s\x1b[23m", strings.Replace(args[0], "\x1b[0m", "\x1b[0m\x1b[3m", -1)), nil
}
