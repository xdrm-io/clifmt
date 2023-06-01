package syntax

import (
	"fmt"
	"regexp"
)

// Hyperlink implements transform.Transformer for the hyperlink syntax
type Hyperlink struct{}

// Regex implements transform.Transformer
func (Hyperlink) Regex() *regexp.Regexp {
	return regexp.MustCompile(`(?m)(^|[^\x1b])\[([^(?:\]()]+)\]\(([^\)]+)\)`)
}

// Transform implements transform.Transformer
func (Hyperlink) Transform(args ...string) (string, error) {
	// ignore no arg or empty
	if len(args) < 2 || len(args[1]) < 1 {
		return "", nil
	}

	linkstart := fmt.Sprintf("\x1b]8;;%s\x1b\\", args[2])
	linkend := fmt.Sprintf("\x1b]8;;\x1b\\")

	return fmt.Sprintf("%s%s%s%s", args[0], linkstart, args[1], linkend), nil
}
