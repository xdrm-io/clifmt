package syntax

import (
	"fmt"
	"regexp"
	"strings"
)

// Bold implements transform.Transformer for the bold syntax
type Bold struct{}

// Regex implements transform.Transformer
func (Bold) Regex() *regexp.Regexp {
	return regexp.MustCompile(`(?m)\*\*((?:[^\*]+\*?)+)\*\*`)
}

// Transform implements transform.Transformer
func (Bold) Transform(args ...string) (string, error) {
	// ignore no arg or empty
	if len(args) < 1 || len(args[0]) < 1 {
		return "", nil
	}

	return fmt.Sprintf("\x1b[1m%s\x1b[22m", strings.Replace(args[0], "\x1b[0m", "\x1b[0m\x1b[1m", -1)), nil
}
