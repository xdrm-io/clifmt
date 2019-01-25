package markdown

import (
	"fmt"
	"regexp"
	"strings"
)

var underlineRe = regexp.MustCompile(`(?m)_([^_]+)_`)

// underline returns the terminal-formatted underline text @t
func underline(t string) string {
	return fmt.Sprintf("\x1b[4m%s\x1b[24m", strings.Replace(t, "\x1b[0m", "\x1b[0m\x1b[4m", -1))
}

// underlineTransform the @input text using markdown-like syntax :
// - "normal _underline_ normal"
func underlineTransform(input string) (string, error) {
	output := ""
	cursor := int(0)

	// 1. Replace for each match
	for _, match := range underlineRe.FindAllStringSubmatchIndex(input, -1) {

		// (1) add gap between input start OR previous match
		output += input[cursor:match[0]]
		cursor = match[1]

		// (2) extract features
		text := ""

		if match[3]-match[2] > 0 {
			text = input[match[2]:match[3]]
		}

		// (3) replace text with bold text
		output += underline(text)
	}

	// 2. Add end of input
	if cursor < len(input)-1 {
		output += input[cursor:]
	}

	// 3. print final output
	return output, nil

}
