package markdown

import (
	"fmt"
	"regexp"
)

var hyperlinkRe = regexp.MustCompile(`(?m)\[([^\[]+)\]\(([^\)]+)\)`)

// linkify returns the terminal-formatted hyperlink for @url with the text : @label
func linkify(url, label string) string {
	return fmt.Sprintf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", url, label)
}

// hyperlinkTransform the @input text using markdown-like syntax :
// - "normal [link label](link url) normal"
func hyperlinkTransform(input string) (string, error) {
	output := ""
	cursor := int(0)

	// 1. Replace for each match
	for _, match := range hyperlinkRe.FindAllStringSubmatchIndex(input, -1) {

		// (1) add gap between input start OR previous match
		output += input[cursor:match[0]]
		cursor = match[1]

		// (2) extract features
		var label, url string

		if match[3]-match[2] > 0 {
			label = input[match[2]:match[3]]
		}
		if match[5]-match[4] > 0 {
			url = input[match[4]:match[5]]
		}

		// (3) replace with hyperlink
		output += linkify(url, label)
	}

	// 2. Add end of input
	if cursor < len(input) {
		output += input[cursor:]
	}

	// 3. print final output
	return output, nil

}
