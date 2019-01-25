package clifmt

import (
	"fmt"
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

// Sprintf returns a terminal-colorized output following the coloring format
func Sprintf(format string, a ...interface{}) (string, error) {
	// 1. Pre-process format with 'fmt'
	input := fmt.Sprintf(format, a...)
	output := ""
	cursor := int(0)

	// 2. extract color format matches
	for _, match := range extractor.FindAllStringSubmatchIndex(input, -1) {
		// (1) add gap between input start OR previous match
		output += input[cursor:match[0]]
		cursor = match[1]

		// (2) extract features
		var (
			err         = error(nil)
			text        = ""
			sForeground = ""
			sBackground = ""
			foreground  = terminalColor(0)
			background  = terminalColor(0)
		)

		if match[3]-match[2] > 0 {
			text = input[match[2]:match[3]]
		}
		if match[5]-match[4] > 0 {
			sForeground = input[match[4]:match[5]]
			foreground, err = parseColor(sForeground)
			if err != nil {
				return "", err
			}
		}
		if match[7]-match[6] > 0 {
			sBackground = input[match[6]:match[7]]
			background, err = parseColor(sBackground)
			if err != nil {
				return "", err
			}
		}

		// (3) replace text with colorized text
		if len(sForeground) > 0 {
			text = colorize(text, true, foreground)
		}
		if len(sBackground) > 0 {
			text = colorize(text, false, background)
		}
		output += text
	}

	// 3. Add end of input
	if cursor < len(input)-1 {
		output += input[cursor:]
	}

	// 3. print final output
	return output, nil
}

func Printf(format string, a ...interface{}) error {
	s, err := Sprintf(format, a...)
	if err != nil {
		return err
	}

	fmt.Print(s)
	return nil
}

func colorize(text string, foregound bool, color terminalColor) string {
	if foregound {
		return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", color.Red(), color.Green(), color.Blue(), text)
	}

	return fmt.Sprintf("\033[48;2;%d;%d;%dm%s\033[0m", color.Red(), color.Green(), color.Blue(), text)
}
