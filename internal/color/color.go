package color

import (
	"fmt"
	"io"
	"strconv"
)

// T represents a color
type T uint32

// FromName returns the integer value of a color name
// from the built-in color map ; it is case insensitive
func FromName(t Theme, s string) (T, error) {
	value, ok := t[s]
	if !ok {
		return 0, fmt.Errorf("unknown color name '%s'", s)
	}
	return value, nil
}

// FromHex returns the integer value associated with
// an hexadecimal string (full-sized or short version)
// the format is 'abc' or 'abcdef'
func FromHex(s string) (T, error) {
	if len(s) != 3 && len(s) != 6 {
		return 0, fmt.Errorf("expect a size of 3 or 6 (without the '#' prefix)")
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

	return T(n), nil
}

// Parse tries to parse a color string (can be a name or an hexa value)
func Parse(t Theme, s string) (T, error) {

	// (0) ...
	if len(s) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	// (1) hexa
	if s[0] == '#' {
		return FromHex(s[1:])
	}

	// (2) name
	return FromName(t, s)

}

// Red component of the color
func (c T) Red() uint8 {
	return uint8((c >> 16) & 0xff)
}

// Green component of the color
func (c T) Green() uint8 {
	return uint8((c >> 8) & 0xff)
}

// Blue component of the color
func (c T) Blue() uint8 {
	return uint8(c & 0xff)
}