package syntax

import (
	"fmt"
	"regexp"
	"strings"
)

// Underline implements transform.Transformer for the underline syntax
type Underline struct{}

// Regex implements transform.Transformer
func (Underline) Regex() *regexp.Regexp {
	return regexp.MustCompile(`(?m)_([^_]+)_`)
}

// Transform implements transform.Transformer
func (Underline) Transform(args ...string) (string, error) {
	// ignore no arg or empty
	if len(args) < 1 || len(args[0]) < 1 {
		return "", nil
	}
	return fmt.Sprintf("\x1b[4m%s\x1b[24m", strings.Replace(args[0], "\x1b[0m", "\x1b[0m\x1b[4m", -1)), nil
}
