package clifmt

import (
	"fmt"
	"io"
	"strconv"
)

type terminalColor uint32

var colorMap = make(map[string]terminalColor)

func init() {
	colorMap["red"] = 0xff0000
	colorMap["green"] = 0x00ff00
	colorMap["blue"] = 0x0000ff
}

// fromName returns the integer value of a color name
// from the built-in color map ; it is case insensitive
func fromName(s string) (terminalColor, error) {
	value, ok := colorMap[s]
	if !ok {
		return 0, fmt.Errorf("unknown color name '%'", s)
	}
	return value, nil
}

// fromHex returns the integer value associated with
// an hexadecimal string (full-sized or short version)
// the format is 'abc' or 'abcdef'
func fromHex(s string) (terminalColor, error) {
	if len(s) != 3 && len(s) != 6 {
		return 0, fmt.Errorf("expect a size of 3 or 6 (remove the '#' prefix)")
	}

	// short version
	input := s
	if len(s) == 3 {
		input = fmt.Sprintf("%c%c%c%c%c%c", s[0], s[0], s[1], s[1], s[2], s[2])
	}

	n, err := strconv.ParseUint(input, 16, 32)
	if err != nil {
		return 0, err
	}

	return terminalColor(n), nil
}

// parseColor tries to parse a color string (can be a name or an hexa value)
func parseColor(s string) (terminalColor, error) {

	// (0) ...
	if len(s) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	// (1) hexa
	if s[0] == '#' {
		return fromHex(s[1:])
	}

	// (2) name
	return fromName(s)

}

// Red component of the color
func (c terminalColor) Red() uint8 {
	return uint8((c >> 16) & 0xff)
}

// Green component of the color
func (c terminalColor) Green() uint8 {
	return uint8((c >> 8) & 0xff)
}

// Blue component of the color
func (c terminalColor) Blue() uint8 {
	return uint8(c & 0xff)
}
