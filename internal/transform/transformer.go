package transform

import (
	"regexp"
)

// Transformer defines a string transformer
type Transformer interface {
	// Regex returns the regex matching text to replace
	Regex() *regexp.Regexp

	// Transform is called to replace a match by its transformation
	// ; it takes as arguments the matched string chunks from the Regex()
	Transform(...string) (string, error)
}
