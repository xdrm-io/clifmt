package color

import (
	"fmt"
	"git.xdrm.io/go/clifmt/internal/color"
	"regexp"
)

// extractor helps extract features from the coloring format defined as follows :
//
// - [Color] -> [a-z]       # named color
// - [Color] -> #[0-9a-f]{3} # hexa color (shortcode)
// - [Color] -> #[0-9a-f]{6} # hexa color (full-sized)
// - [Text] -> ANY
// - [Format] -> ${Text}(Color:Color) # foreground, background colors
// - [Format] -> ${Text}(Color)       # foreground color only
// - [Format] -> ${Text}(:Color)      # background color only
var extractor = regexp.MustCompile(`(?m)\${([^$]+)}\(((?:[a-z]+|#(?:[0-9a-f]{3}|[0-9a-f]{6})))?(?:\:((?:[a-z]+|#(?:[0-9a-f]{3}|[0-9a-f]{6}))))?\)`)

// colorize returns the terminal-formatted @text colorized with the @fg and @bg colors
func colorize(t string, fg *color.T, bg *color.T) string {
	// no coloring
	if fg == nil && bg == nil {
		return t
	}

	// only foreground
	if bg == nil {
		return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", fg.Red(), fg.Green(), fg.Blue(), t)
	}
	// only background
	if fg == nil {
		return fmt.Sprintf("\x1b[48;2;%d;%d;%dm%s\x1b[0m", bg.Red(), bg.Green(), bg.Blue(), t)
	}

	// both colors
	return fmt.Sprintf("\x1b[38;2;%d;%d;%d;48;2;%d;%d;%dm%s\x1b[0m", fg.Red(), fg.Green(), fg.Blue(), bg.Red(), bg.Green(), bg.Blue(), t)
}

// Transform the @input text colorized according to the @extractor format
func Transform(input string, theme color.Theme) (string, error) {
	output := ""
	cursor := int(0)

	// 1. Replace for each match
	for _, match := range extractor.FindAllStringSubmatchIndex(input, -1) {

		// (1) add gap between input start OR previous match
		output += input[cursor:match[0]]
		cursor = match[1]

		// (2) extract features
		var (
			text          = ""
			sFg           = ""
			sBg           = ""
			fg   *color.T = nil
			bg   *color.T = nil
		)

		if match[3]-match[2] > 0 {
			text = input[match[2]:match[3]]
		}
		if match[5]-match[4] > 0 {
			sFg = input[match[4]:match[5]]
			fgv, err := color.Parse(theme, sFg)
			if err != nil {
				return "", err
			}
			fg = &fgv
		}
		if match[7]-match[6] > 0 {
			sBg = input[match[6]:match[7]]
			bgv, err := color.Parse(theme, sBg)
			if err != nil {
				return "", err
			}
			bg = &bgv
		}

		// (3) replace text with colorized text
		output += colorize(text, fg, bg)
	}

	// 2. Add end of input
	if cursor < len(input) {
		output += input[cursor:]
	}

	// 3. print final output
	return output, nil
}
